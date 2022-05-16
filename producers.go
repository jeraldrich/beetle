package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Producer struct {
	actions *Actions
}

type Actions interface {
	GetMessagesFromUrl(url string) (Messages, err error)
}

func NewProducer() *Producer {
	a := new(Actions)
	return &Producer{a}
}

func (p *Producer) GetMessagesFromUrl(url string) ([]Message, error) {
	var messages []Message
	var httpBody io.Reader
	client, method := &http.Client{}, "GET"
	req, err := http.NewRequest(method, url, httpBody)
	// Force close after request is made or sometimes a panic EOF occurs.
	// defer resp.Body is not enough.
	req.Close = true
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("Fetching messages from url: ", url)
	resp, err := client.Do(req)
	if err != nil {
		return messages, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	// Cause decoding unknown fields to error with "json: unknown field ..."
	decoder.DisallowUnknownFields()
	// result := make([]Message, 0)
	err = decoder.Decode(&messages)
	if err != nil {
		fmt.Println("json decode error ", err)
		return messages, err
	}

	return messages, err
}
