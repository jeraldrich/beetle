package main

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
	var err error

	for i := 0; i < len(messages); i++ {
		message := messages[i]

		message.CreatedAt, err = filter.DatetimeString(message.CreatedAt)
		if err != nil {
			//log.Printf("error when cleaning message id [%s]: %s", message.ID, err)
			// Set dirty fields so we can later store then in seperate table.
			error_count += 1
			message.dirty_fields = true
			message.dirty_field_error = err
			continue
		}

		success_count += 1
		message.dirty_fields = false
		messages[i] = message
	}

	return messages, success_count, error_count
}
