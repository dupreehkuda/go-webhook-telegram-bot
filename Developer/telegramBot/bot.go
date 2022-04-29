package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testy/schedule"
	"testy/yfapi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5340786204:AAGEzI8eVZcAk5xqGiuJNgfHbvcgVTTCuzY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var replyMessage string

			switch update.Message.Text {
			case "/start":
				replyMessage = "Здарова"
			case "/today", "/today@cooltesticles_bot":
				replyMessage = schedule.GetSchedule("Сегодня", 0, schedule.UpComming)
			case "/tomorrow", "/tomorrow@cooltesticles_bot":
				replyMessage = schedule.GetSchedule("Завтра", 1, schedule.UpComming)
			case "/sanyok":
				replyMessage = "@KtulhuSlayer бля Сань проснись"
			default:
				if strings.Contains(update.Message.Text, "/stonk") {
					replyMessage = yfapi.StockAnswer(update.Message.Text)
				} else {
					continue
				}
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyMessage)
			msg.Text = replyMessage

			bot.Send(msg)
		}
	}
}

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

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// // Check if the message contains the word "marco"
	// // if not, return without doing anything
	// if !strings.Contains(strings.ToLower(body.Message.Text), "marco") {
	// 	return
	// }

	switch body.Message.Text {
	case "/start":
		sendMessage(body.Message.Chat.ID, "Здарова")
	case "/today", "/today@cooltesticles_bot":
		sendMessage(body.Message.Chat.ID, schedule.GetSchedule("Сегодня", 0, schedule.UpComming))
	case "/tomorrow", "/tomorrow@cooltesticles_bot":
		sendMessage(body.Message.Chat.ID, schedule.GetSchedule("Завтра", 1, schedule.UpComming))
	case "/sanyok":
		sendMessage(body.Message.Chat.ID, "@KtulhuSlayer бля Сань проснись")
	case "/now":
		sendMessage(body.Message.Chat.ID, schedule.Current(schedule.UpComming))
	case "/cool":
		sendMessage(body.Message.Chat.ID, coolify(body.Message.Text))
	default:
		if strings.Contains(body.Message.Text, "/stonk") {
			sendMessage(body.Message.Chat.ID, yfapi.StockAnswer(body.Message.Text))
		} else if strings.Contains(body.Message.Text, "/stonk") {
			sendMessage(body.Message.Chat.ID, coolify(body.Message.Text))
		} else {
			return
		}
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := sendMessage(body.Message.Chat.ID, "plol"); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// sayPolo takes a chatID and sends "polo" to them
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
	res, err := http.Post("https://api.telegram.org/bot5340786204:AAGEzI8eVZcAk5xqGiuJNgfHbvcgVTTCuzY/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
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
