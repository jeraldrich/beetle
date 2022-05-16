package main

import (
	"reflect"

	"github.com/scylladb/gocqlx/v2"
)

type Consumer struct {
	action *ConsumerActions
}

type ConsumerActions interface {
	CreateMessage(message Message) (err error)
}

func NewConsumer() *Consumer {
	a := new(ConsumerActions)
	return &Consumer{a}
}

func (p *Consumer) FilterMessages(messages []Message) (cleanMessages []Message, success_count int, error_count int) {
	filter := NewFilter()

	fieldsToFilter := map[string]interface{}{
		"ID":         "uuid",
		"FromUserId": "uuid",
		"ToUserId":   "uuid",
		"CreatedAt":  "datetime",
		"UpdatedAt":  "datetime",
		"SendAt":     "datetime",
		"SentAt":     "datetime",
		"ReadAt":     "datetime",
		"SlackTs":    "datetime",
		"CanceledAt": "datetime",
		"DeletedAt":  "datetime",
		"ActedOnAt":  "datetime",
		"ViewedAt":   "datetime",
		"PausedAt":   "datetime",
	}

	for i := 0; i < len(messages); i++ {
		message := messages[i]

		// Loops through message keys and match with fields to filter.
		// Probably better way to do this.
		for filterFieldName, filterType := range fieldsToFilter {
			reflectValue := reflect.ValueOf(&message).Elem()
			fieldValue := reflectValue.FieldByName(filterFieldName)

			var filteredResult string
			var err error

			switch filterType {
			case "uuid":
				filteredResult, err = filter.Uuid(fieldValue.String())
			case "datetime":
				filteredResult, err = filter.DatetimeString(fieldValue.String())
			}
			if err != nil {
				//log.Printf("error when cleaning message id [%s]: %s", message.ID, err)
				// Set dirty fields so we can later store then in seperate table.
				error_count += 1
				message.dirtyField = true
				message.dirtyFieldError = err
				continue
			}

			// Set original field value to new filtered value
			reflectValue.FieldByName(filterFieldName).SetString(filteredResult)

			success_count += 1
			message.dirtyField = false
			messages[i] = message
		}
	}

	return messages, success_count, error_count
}

func (p *Consumer) CreateMessages(messages []Message, session gocqlx.Session) (err error) {

	for i := 0; i < len(messages); i++ {
		message := messages[i]

		q := session.Query(MessagesTable.Insert()).BindStruct(message)
		if err := q.ExecRelease(); err != nil {
			return err
		}
	}

	return err
}
