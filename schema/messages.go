// Code generated by "gocqlx/cmd/schemagen"; DO NOT EDIT.

package schema

import "github.com/scylladb/gocqlx/v2/table"

// Table models.
var (
	Messages = table.New(table.Metadata{
		Name: "messages",
		Columns: []string{
			"acted_on_at",
			"associated_id",
			"associated_type",
			"body",
			"canceled_at",
			"channel_id",
			"correlation_id",
			"created_at",
			"deleted_at",
			"delivery_type",
			"duration",
			"from_user_id",
			"id",
			"is_flagged",
			"notification_count",
			"paused_at",
			"read_at",
			"send_at",
			"sender_user_id",
			"sent_at",
			"sent_automatically",
			"slack_ts",
			"state",
			"sub_type",
			"tags",
			"to_user_id",
			"updated_at",
			"urls",
			"viewed_at",
			"viewed_duration",
		},
		PartKey: []string{
			"id",
		},
		SortKey: []string{},
	})
)