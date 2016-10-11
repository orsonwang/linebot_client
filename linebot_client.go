package main

import (
	"fmt"
	logger "log"
	"net/http"
	"os"
	//	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/nats-io/nats"
)

var mainLineBotClient *linebot.Client
var mainLineEventHandler *LineBotEventHandler
var log *logger.Logger
var nc *nats.Conn

func main() {
	f, err := os.OpenFile("./linebot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	log = new(logger.Logger)
	log.SetOutput(f)

	mainLineBotClient, err = linebot.New(os.Getenv("LINE_CHANNEL_SECRET"), os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatalf("Linebot init error: %s\n", err)
	}
	mainLineEventHandler = new(LineBotEventHandler)
	mainLineEventHandler.SetLineBotClient(mainLineBotClient)

	urls := "nats://localhost:4222"
	nc, err = nats.Connect(urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}
	defer nc.Close()

	http.HandleFunc("/callback", lineCallbackHandler)

	port := os.Getenv("LINEBOT_PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func lineCallbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("=== Line Callback ===")
	events, err := mainLineBotClient.ParseRequest(r)
	if err != nil {
		log.Printf("Parse error: %s\n", err)
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				mainLineEventHandler.OnTextMessage(event.ReplyToken, message.Text)
				break
			}
		}
	}
}
