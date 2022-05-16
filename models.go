package main

import (
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

// messageTable allows for simple CRUD operations based on messageMetadata.
// var messagesTable = table.New(MessagesMetadata)
//  "{\"id\": \"00000000-0000-0000-0000-000000000000\", \"notes\": \"\", \"workout\": {\"id\": \"00000000-0000-0000-0000-000000000000\", \"name\": \"\", \"type\": \"\", \"missed_at\": \"0001-01-01T00:00:00Z\", \"description\": \"\", \"is_optional\": null, \"scheduled_at\": \"0001-01-01T00:00:00Z\", \"activity_type\": \"\", \"trainer_notes\": \"\"}, \"timezone\": \"\", \"difficulty\": null, \"started_at\": \"0001-01-01T00:00:00Z\", \"is_optional\": null, \"completed_at\": \"0001-01-01T00:00:00Z\", \"sets_skipped\": null, \"sets_too_long\": null, \"actual_duration\": 0, \"weights_changed\": null, \"completion_state\": \"\", \"sets_not_started\": null, \"active_sets_total\": 0, \"active_sets_completed\": 0, \"completed_automatically\": false}",

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
