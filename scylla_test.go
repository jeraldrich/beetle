package main

import (
	"testing"

	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/gocqlxtest"
)

func TestScylla(t *testing.T) {
	// cfg := gocql.ClusterConfig{Hosts: []string{"127.0.0.1"}, Port: 9042}
	cluster := gocqlxtest.CreateCluster()
	cluster.Hosts = []string{"127.0.0.1"}
	cluster.Port = 9042

	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		t.Fatal("create session:", err)
	}
	defer session.Close()

	session.ExecStmt(`DROP KEYSPACE test_messages`)

	readScyllaVersion(t, session)
	createKeyspace(t, session)
}

// This example shows how to use query builder to work with
func createKeyspace(t *testing.T, session gocqlx.Session) {
	// TODO: Add 		attributes map<uuid, text, map<uuid, >>
	// {\"id\": \"00000000-0000-0000-0000-000000000000\", \"notes\": \"\", \"workout\": {\"id\": \"00000000-0000-0000-0000-000000000000\", \"name\": \"\", \"type\": \"\", \"missed_at\": \"0001-01-01T00:00:00Z\", \"description\": \"\", \"is_optional\": null, \"scheduled_at\": \"0001-01-01T00:00:00Z\", \"activity_type\": \"\", \"trainer_notes\": \"\"}, \"timezone\": \"\", \"difficulty\": null, \"started_at\": \"0001-01-01T00:00:00Z\", \"is_optional\": null, \"completed_at\": \"0001-01-01T00:00:00Z\", \"sets_skipped\": null, \"sets_too_long\": null, \"actual_duration\": 0, \"weights_changed\": null, \"completion_state\": \"\", \"sets_not_started\": null, \"active_sets_total\": 0, \"active_sets_completed\": 0, \"completed_automatically\": false}",
	err := session.ExecStmt(`CREATE TABLE IF NOT EXISTS test_db.messages (
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
		sub_type string,
		viewed_at timestamp,
		viewed_duration int,
		urls set<text>,
		duration int,
		paused_at timestamp,
		delivery_type string,
		notification_count int
		)`)
	if err != nil {
		t.Fatal("create table:", err)
	}

}

func readScyllaVersion(t *testing.T, session gocqlx.Session) {
	var releaseVersion string

	err := session.Query("SELECT release_version FROM system.local", nil).Get(&releaseVersion)
	if err != nil {
		t.Fatal("Get() failed:", err)
	}

	t.Logf("Scylla version is: %s", releaseVersion)
}
