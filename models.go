package main

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/table"
)

type Config struct {
	JsonDataUrls []string
}

// metadata specifies table name and columns it must be in sync with schema.
var MessagesMetadata = table.Metadata{
	Name: "production_db.messages",
	Columns: []string{
		"id", "created_at", "updated_at", "send_at", "sent_at", "from_user_id",
		"to_user_id", "body", "state", "read_at", "sent_automatically", "tag",
		"type", "associated_type", "associated_id", "is_flagged", "slack_ts",
		"channel_id", "canceled_at", "deleted_at", "attributes", "acted_on_at",
		"sender_user_id", "correlation_id", "sub_type", "viewed_at", "viewed_duration",
		"urls", "duration", "paused_at", "delivery_type", "notification_count",
	},
	PartKey: []string{"id"},
	SortKey: []string{"created_at"},
}
var MessagesTable = table.New(MessagesMetadata)

// The message to the serialized from json
type Message struct {
	ID                string `json:"id"` // gocql.UUID
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	SendAt            string `json:"send_at"`
	SentAt            string `json:"sent_at"`
	FromUserId        string `json:"from_user_id"` // gocql.UUID can be uuid or 00000000-0000-0000-0000-000000000000
	ToUserId          string `json:"to_user_id"`   // gocql.UUID can be uuid or 00000000-0000-0000-0000-000000000000
	Body              string `json:"body"`
	State             string `json:"state"`
	ReadAt            string `json:"read_at"`
	SentAutomatically bool   `json:"sent_automatically"`
	Tag               string `json:"tag"`
	Type              string `json:"type"`
	AssociatedType    string `json:"associated_type"`
	AssociatedId      string `json:"associated_id"` // can be uuid, 00000000-0000-0000-0000-000000000000 or null
	IsFlagged         bool   `json:"is_flagged"`
	SlackTs           string `json:"slack_ts"`       // Nullstring
	ChannelId         string `json:"channel_id"`     // can be uuid, 00000000-0000-0000-0000-000000000000 or null
	CanceledAt        string `json:"canceled_at"`    // Nullstring
	DeletedAt         string `json:"deleted_at"`     // Nullstring
	Attributes        string `json:"attributes"`     // Nullstring
	ActedOnAt         string `json:"acted_on_at"`    // Nullstring
	SenderUserId      string `json:"sender_user_id"` // Nullstring
	CorrelationId     string `json:"correlation_id"` // Nullstring
	SubType           string `json:"sub_type"`       // Nullstring
	ViewedAt          string `json:"viewed_at"`      // Nullstring
	ViewedDuration    int    `json:"viewed_duration"`
	Urls              string `json:"urls"` // Nullstring
	Duration          int    `json:"duration"`
	PausedAt          string `json:"paused_at"`     // Nullstring
	DeliveryType      string `json:"delivery_type"` // Nullstring
	NotificationCount int    `json:"notification_count"`

	// Used to store validation errors when filtering through all message results.
	// Want to store this in db at some point to report bad imports.
	// TODO: Refactor to map, or array of errors.
	dirtyField      bool
	dirtyFieldError error
}

// The message to be inserted into the database.
type TransformedMessage struct {
	ID                gocql.UUID
	CreatedAt         time.Time
	UpdatedAt         time.Time
	SendAt            time.Time
	SentAt            time.Time
	FromUserId        gocql.UUID
	ToUserId          gocql.UUID
	Body              string
	State             string
	ReadAt            time.Time
	SentAutomatically bool
	Tag               string
	Type              string
	AssociatedType    string
	AssociatedId      gocql.UUID
	IsFlagged         bool
	SlackTs           time.Time
	ChannelId         gocql.UUID
	CanceledAt        time.Time
	DeletedAt         time.Time
	Attributes        string
	ActedOnAt         time.Time
	SenderUserId      gocql.UUID
	CorrelationId     gocql.UUID
	SubType           string
	ViewedAt          time.Time
	ViewedDuration    int
	Urls              string
	Duration          int
	PausedAt          time.Time
	DeliveryType      string
	NotificationCount int
}
