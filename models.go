package main

import (
	"database/sql"

	"github.com/gocql/gocql"
)

type Messages struct {
	Messages []Message // `json:"messages"`
}

type Message struct {
	ID                 gocql.UUID     `json:"id"`
	created_at         string         `json:"created_at"`
	updated_at         string         `json:"updated_at"`
	send_at            string         `json:"send_at"`
	sent_at            string         `json:"sent_at"`
	from_user_id       gocql.UUID     `json:"from_user_id"`
	to_user_id         gocql.UUID     `json:"to_user_id"`
	body               string         `json:"body"`
	state              string         `json:"state"`
	read_at            string         `json:"read_at"`
	sent_automatically bool           `json:"sent_automatically"`
	tag                string         `json:"tag"`
	message_type       string         `json:"type"`
	associated_type    string         `json:"associated_type"`
	associated_id      gocql.UUID     `json:"associated_id"`
	is_flagged         bool           `json:"is_flagged"`
	slack_ts           sql.NullString `json:"slack_ts"`
	channel_id         gocql.UUID     `json:"channel_id"`
	canceled_at        sql.NullString `json:"canceled_at"`
	deleted_at         sql.NullString `json:"deleted_at"`
	attributes         sql.NullString `json:"attributes"`
	acted_on_at        sql.NullString `json:"acted_on_at"`
	sender_user_id     sql.NullString `json:"sender_user_id"`
	correlation_id     sql.NullString `json:"correlation_id"`
	sub_type           sql.NullString `json:"sub_type"`
	viewed_at          sql.NullString `json:"viewed_at"`
	viewed_duration    int            `json:"viewed_duration"`
	urls               sql.NullString `json:"urls"`
	duration           int            `json:"duration"`
	paused_at          sql.NullString `json:"paused_at"`
	delivery_type      sql.NullString `json:"delivery_type"`
	notification_count int            `json:"notification_count"`
}
