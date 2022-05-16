package main

import (
	"log"
	"reflect"
	"time"

	"github.com/gocql/gocql"
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

func (p *Consumer) FilterMessages(messages []Message) (t []TransformedMessage, success_count int, error_count int) {
	filter := NewFilter()
	transformedMessages := []TransformedMessage{}
	var blankUuid [16]byte // used when checking if blank uuid returned on uuid filter.
	fieldsToFilter := map[string]interface{}{
		"ID":           "uuid",
		"FromUserId":   "uuid",
		"ToUserId":     "uuid",
		"AssociatedId": "uuid",
		"ChannelId":    "uuid",
		"SenderUserId": "uuid",
		"Correlation":  "uuid",
		"CreatedAt":    "datetime",
		"UpdatedAt":    "datetime",
		"SendAt":       "datetime",
		"SentAt":       "datetime",
		"ReadAt":       "datetime",
		"SlackTs":      "datetime",
		"CanceledAt":   "datetime",
		"DeletedAt":    "datetime",
		"ActedOnAt":    "datetime",
		"ViewedAt":     "datetime",
		"PausedAt":     "datetime",
	}

	for i := 0; i < len(messages); i++ {
		message := messages[i]
		// Set all fields from message that do not need to be transformed.
		transformedMessage := TransformedMessage{
			Body:              message.Body,
			State:             message.State,
			SentAutomatically: message.SentAutomatically,
			Tag:               message.Tag,
			Type:              message.Type,
			AssociatedType:    message.AssociatedType,
			IsFlagged:         message.IsFlagged,
			Attributes:        message.Attributes,
			SubType:           message.SubType,
			ViewedDuration:    message.ViewedDuration,
			Urls:              message.Urls,
			Duration:          message.Duration,
			DeliveryType:      message.DeliveryType,
			NotificationCount: message.NotificationCount,
		}

		// Loops through fieldsToFilter to match with message and transformedMessage fields.
		// Transform the value from message and set in transformedMessage.
		for filterFieldName, filterType := range fieldsToFilter {
			messageReflectValue := reflect.ValueOf(&message).Elem()
			messageFieldValue := messageReflectValue.FieldByName(filterFieldName)

			transformReflectValue := reflect.ValueOf(&transformedMessage).Elem()
			transformFieldValue := transformReflectValue.FieldByName(filterFieldName)

			var filteredUuid gocql.UUID
			var filteredDatetime time.Time
			var err error

			switch filterType {
			case "uuid":
				filteredUuid, err = filter.Uuid(messageFieldValue.String())
				if filteredUuid != blankUuid && err == nil {
					// log.Println("setting filtered uuid 2 for  filterFieldName", filterFieldName, filteredUuid)
					transformFieldValue.Set(reflect.ValueOf(filteredUuid))
				}
			case "datetime":
				// TODO: Handle 2022/05/16
				filteredDatetime, err = filter.DatetimeString(messageFieldValue.String())
				if !filteredDatetime.IsZero() && err == nil {
					// log.Println("setting filtered date for fieldname ", filterFieldName, filteredDatetime)
					transformFieldValue.Set(reflect.ValueOf(filteredDatetime))
				}
			}
			if err != nil {
				log.Printf("error when cleaning message field name [%s]: %s", filterFieldName, err)
				// Set dirty fields so we can later store then in seperate table.
				error_count += 1
				message.dirtyField = true
				message.dirtyFieldError = err
				continue
			}

			success_count += 1
			message.dirtyField = false
			messages[i] = message
			transformedMessages = append(transformedMessages, transformedMessage)
		}
	}

	return transformedMessages, success_count, error_count
}

func (p *Consumer) CreateMessages(transformedMessages []TransformedMessage, session gocqlx.Session) (err error) {

	for i := 0; i < len(transformedMessages); i++ {
		transformedMessage := transformedMessages[i]
		// log.Println("inserting message ", transformedMessage)

		q := session.Query(MessagesTable.Insert()).BindStruct(transformedMessage)
		if err := q.ExecRelease(); err != nil {
			return err
		}
	}

	return err
}
