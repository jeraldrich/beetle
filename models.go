package main

import (
	"database/sql"

	"github.com/gocql/gocql"
)

type Message struct {
	ID                 gocql.UUID
	created_at         string
	updated_at         string
	send_at            string
	sent_at            string
	from_user_id       gocql.UUID
	to_user_id         gocql.UUID
	body               string
	state              string
	read_at            string
	sent_automatically bool
	message_type       string
	associated_type    string
	associated_id      gocql.UUID
	is_flagged         bool
	slack_ts           sql.NullString
	channel_id         gocql.UUID
	canceled_at        sql.NullString
	deleted_at         sql.NullString
	attributes         sql.NullString
	acted_on_at        sql.NullString
	sender_user_id     sql.NullString
	correlation_id     sql.NullString
	sub_type           sql.NullString
	viewed_at          sql.NullString
	viewed_duration    int
	urls               sql.NullString
	duration           int
	paused_at          sql.NullString
	delivery_type      sql.NullString
	notification_count int

	Tags []string
	Data []byte
}
