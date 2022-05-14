package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Producer struct {
	action *Action
}

type Action interface {
	GetMessagesFromUrl(url string) (Messages, err error)
}

func NewProducer() *Producer {
	a := new(Action)
	return &Producer{a}
}

func (p *Producer) GetMessagesFromUrl(url string) (Messages, error) {
	var jsonResult Messages

	resp, err := http.Get(url)
	if err != nil {
		return jsonResult, err
	}
	defer resp.Body.Close()

	fmt.Println("Fetching messages from url: ", url)

	err = json.NewDecoder(resp.Body).Decode(&jsonResult)
	if err != nil {
		return jsonResult, err
	}

	return jsonResult, err
}
