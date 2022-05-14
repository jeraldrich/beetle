package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func main() {
	host := "127.0.0.1:9042"
	cluster := gocql.NewCluster(host)

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		fmt.Printf("Failed to connect to cql cluster: [%s]", host)
	}
	defer session.Close()

	err = InitializeMessagesKeyspace(session)
	if err != nil {
		panic(err)
	}

	err = InitializeMessagesTable(session)
	if err != nil {
		panic(err)
	}

	// TODO: setup consumer, setup consumer channel pool
	producer := NewProducer()
	messages, err := producer.GetMessagesFromUrl("")
	if err != nil {
		panic(err)
	}
	log.Println(messages)

	consumer := NewConsumer()
	messages, success_count, error_count := consumer.FilterMessages(messages)
	log.Printf("Filtered %d messages: success[%d] error_count[%d]", len(messages), success_count, error_count)

	// err = consumer.CreateMessages(messages)
}

func InitializeMessagesKeyspace(session gocqlx.Session) (err error) {
	err = session.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS production_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`)
	if err != nil {
		fmt.Println("Not able to create scylla messages keyspace: ", err)
	}

	return err
}

func InitializeMessagesTable(session gocqlx.Session) (err error) {
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
		tags set<text>,
		associated_type text,
		associated_id uuid,
		is_flagged boolean,
		slack_ts timestamp,
		channel_id uuid,
		canceled_at timestamp,
		deleted_at timestamp,
		acted_on_at timestamp,
		sender_user_id uuid,
		correlation_id uuid,
		sub_type text,
		viewed_at timestamp,
		viewed_duration int,
		urls set<text>,
		duration int,
		paused_at timestamp,
		delivery_type text,
		notification_count int
		)`)
	if err != nil {
		fmt.Println("Not able to create scylla messages table: ", err)
		return err
	}

	return err
}
