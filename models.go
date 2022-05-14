package main

import (
	"github.com/gocql/gocql"
)

type Configuration struct {
	messagesDataUrls []string
}

type Message struct {
	ID                gocql.UUID `json:"id"`
	CreatedAt         string     `json:"created_at"`
	UpdatedAt         string     `json:"updated_at"`
	SendAt            string     `json:"send_at"`
	SentAt            string     `json:"sent_at"`
	FromUserId        gocql.UUID `json:"from_user_id"` // can be uuid or 00000000-0000-0000-0000-000000000000
	ToUserId          gocql.UUID `json:"to_user_id"`   // can be uuid or 00000000-0000-0000-0000-000000000000
	Body              string     `json:"body"`
	State             string     `json:"state"`
	ReadAt            string     `json:"read_at"`
	SentAutomatically bool       `json:"sent_automatically"`
	Tag               string     `json:"tag"`
	MessageType       string     `json:"type"`
	AssociatedType    string     `json:"associated_type"`
	AssociatedId      string     `json:"associated_id"` // can be uuid, 00000000-0000-0000-0000-000000000000 or null
	IsFlagged         bool       `json:"is_flagged"`
	SlackTs           string     `json:"slack_ts"`       // Nullstring
	ChannelId         string     `json:"channel_id"`     // can be uuid, 00000000-0000-0000-0000-000000000000 or null
	CanceledAt        string     `json:"canceled_at"`    // Nullstring
	DeletedAt         string     `json:"deleted_at"`     // Nullstring
	Attributes        string     `json:"attributes"`     // Nullstring
	ActedOnAt         string     `json:"acted_on_at"`    // Nullstring
	SenderUserId      string     `json:"sender_user_id"` // Nullstring
	CorrelationId     string     `json:"correlation_id"` // Nullstring
	SubType           string     `json:"sub_type"`       // Nullstring
	ViewedAt          string     `json:"viewed_at"`      // Nullstring
	ViewedDuration    int        `json:"viewed_duration"`
	Urls              string     `json:"urls"` // Nullstring
	Duration          int        `json:"duration"`
	PausedAt          string     `json:"paused_at"`     // Nullstring
	DeliveryType      string     `json:"delivery_type"` // Nullstring
	NotificationCount int        `json:"notification_count"`

	// Used to store validation errors when filtering through all message results.
	// Want to store this in db at some point to report bad imports.
	// TODO: Refactor to map, or array of errors.
	dirty_fields      bool
	dirty_field_error error
}
