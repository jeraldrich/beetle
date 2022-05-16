package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func Produce(consumerChannel chan<- []Message, urls []string) {
	for i := 0; i < len(urls); i++ {
		url := urls[i]
		producer := NewProducer()
		messages, err := producer.GetMessagesFromUrl(url)
		if err != nil {
			panic(err)
		}
		consumerChannel <- messages
	}

	close(consumerChannel)
}

func Consume(wg *sync.WaitGroup, consumerChannel <-chan []Message, session gocqlx.Session) {
	defer wg.Done()
	consumer := NewConsumer()
	transformedMessages := []TransformedMessage{}
	messages := <-consumerChannel
	transformedMessages, success_count, error_count := consumer.FilterMessages(messages)
	log.Printf("Filtered %d messages: success[%d] error_count[%d]", len(messages), success_count, error_count)

	err := consumer.CreateMessages(transformedMessages, session)
	if err != nil {
		panic(err)
	}
}

func main() {
	host := "127.0.0.1:9042"
	cluster := gocql.NewCluster(host)

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Failed to connect to cql cluster: [%s]", host)
	}
	defer session.Close()

	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("configuration file error:", err)
	}

	// Create keyspace, drop existing messages table, and create messages table.
	err = InitializeMessagesKeyspace(session)
	if err != nil {
		panic(err) // AHHHHH!
	}
	err = DropMessagesTable(session)
	if err != nil {
		panic(err)
	}
	err = InitializeMessagesTable(session)
	if err != nil {
		panic(err)
	}

	// producerChannel := make(chan string)
	var wg sync.WaitGroup
	consumerChannel := make(chan []Message)

	for i := 0; i < len(config.JsonDataUrls); i++ {
		wg.Add(1)
		go Consume(&wg, consumerChannel, session)
	}
	Produce(consumerChannel, config.JsonDataUrls)

	wg.Wait()
}

func InitializeMessagesKeyspace(session gocqlx.Session) (err error) {
	err = session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS production_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`)
	if err != nil {
		fmt.Println("Not able to create scylla messages keyspace: ", err)
	}

	return err
}

func InitializeMessagesTable(session gocqlx.Session) (err error) {
	// attributes is set as text, but stores json just fine and can be unmarshalled.
	err = session.ExecStmt(`CREATE TABLE IF NOT EXISTS production_db.messages (
		id uuid PRIMARY KEY,
		created_at timestamp,
		updated_at timestamp,
		send_at timestamp,
		sent_at timestamp,
		read_at timestamp,
		from_user_id uuid,
		to_user_id uuid,
		body text,
		state text,
		sent_automatically boolean,
		tag text,
		type text,
		associated_type text,
		associated_id uuid,
		is_flagged boolean,
		slack_ts timestamp,
		channel_id uuid,
		canceled_at timestamp,
		deleted_at timestamp,
		attributes text,
		acted_on_at timestamp,
		sender_user_id uuid,
		correlation_id uuid,
		sub_type text,
		viewed_at timestamp,
		viewed_duration int,
		urls text,
		duration int,
		paused_at timestamp,
		delivery_type text,
		notification_count int
		)`)
	// PRIMARY KEY (id),
	// PARTITION KEY (id),
	if err != nil {
		fmt.Println("Not able to create scylla messages table: ", err)
		return err
	}

	return err
}

func DropMessagesTable(session gocqlx.Session) (err error) {
	err = session.ExecStmt(`DROP TABLE IF EXISTS production_db.messages`)
	if err != nil {
		fmt.Println("Not able to drop messages table: ", err)
	}

	return err
}
