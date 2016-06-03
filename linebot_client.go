package main

import (
	"fmt"
	logger "log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/nats-io/nats"
)

var mainLineBotClient *linebot.Client
var mainLineEventHandler *LineBotEventHandler
var log *logger.Logger
var nc *nats.Conn

func main() {
	f, _ := os.OpenFile("./linebot.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	log = new(logger.Logger)
	log.SetOutput(f)

	strID := os.Getenv("LINE_CHANNEL_ID")
	lineChannelID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Fatal("Wrong environment setting about LINE_CHANNEL_ID")
	}
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineMID := os.Getenv("LINE_MID")
	mainLineBotClient, _ = linebot.NewClient(lineChannelID, lineChannelSecret, lineMID)

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
	received, err := mainLineBotClient.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}
	for _, result := range received.Results {
		content := result.Content()
		switch content.OpType {
		case linebot.OpTypeAddedAsFriend:
			mainLineEventHandler.OnAddedAsFriendOperation(result.RawContent.Params)
			break
		case linebot.OpTypeBlocked:
			mainLineEventHandler.OnBlockedAccountOperation(result.RawContent.Params)
			break
		default:
			handleContentByType(mainLineEventHandler, content)
			break
		}
	}

}

func handleContentByType(eventHandler *LineBotEventHandler, content *linebot.ReceivedContent) {
	switch content.ContentType {
	case linebot.ContentTypeText:
		textContent, err := content.TextContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnTextMessage(content.From, textContent.Text)
		break
	case linebot.ContentTypeImage:
		imageContent, err := content.ImageContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnImageMessage(content.From, imageContent.ReceivedContent)
		break
	case linebot.ContentTypeVideo:
		videoContent, err := content.VideoContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnVideoMessage(content.From, videoContent.ReceivedContent)
		break
	case linebot.ContentTypeAudio:
		audioContent, err := content.AudioContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnAudioMessage(content.From, audioContent.Duration)
		break
	case linebot.ContentTypeLocation:
		locationContent, err := content.LocationContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnLocationMessage(content.From, locationContent.Title, locationContent.Address, locationContent.Latitude, locationContent.Longitude)
		break
	case linebot.ContentTypeSticker:
		stickerContent, err := content.StickerContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnStickerMessage(content.From, stickerContent.ID, stickerContent.PackageID, stickerContent.Version)
		break
	case linebot.ContentTypeContact:
		contactContent, err := content.ContactContent()
		if err != nil {
			log.Print(err)
			return
		}
		eventHandler.OnContactMessage(content.From, contactContent.Mid, contactContent.DisplayName)
		break
	}

}
