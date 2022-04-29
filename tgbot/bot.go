package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	switch body.Message.Text {
	case "/start":
		sendMessage(body.Message.Chat.ID, "Здарова")
	case "/cool":
		sendMessage(body.Message.Chat.ID, coolify(body.Message.Text))
	default:
		if strings.Contains(body.Message.Text, "/cool") {
			sendMessage(body.Message.Chat.ID, coolify(body.Message.Text))
		} else {
			return
		}
	}

	if err := sendMessage(body.Message.Chat.ID, "plol"); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

func sendMessage(chatID int64, message string) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}

	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot<TOKEN GOES HERE>/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func coolify(message string) string {
	var cooled []string
	result := strings.Split(message, " ")
	for i := 1; i < len(result); i++ {
		cooled = append(cooled, result[i])
	}
	cooled = append(cooled, "😎")
	return strings.Join(cooled, " ")
}
