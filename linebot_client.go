package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

var botClient *linebot.Client

func main() {
	strID := os.Getenv("LINE_CHANNEL_ID")
	channelID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Fatal("Wrong environment setting about LINE_CHANNEL_ID")
	}
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	mid := os.Getenv("LINE_MID")

	botClient, err = linebot.NewClient(channelID, channelSecret, mid)

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("LINEBOT_PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("=== callback ===\n")
	received, err := botClient.ParseRequest(r)
	if err != nil {
		log.Print(err)
		return
	}
	eventHandler := new(LineBotEventHandler)
	eventHandler.setLineBotClient(botClient)

	for _, result := range received.Results {
		content := result.Content()
		switch content.OpType {
		case linebot.OpTypeAddedAsFriend:
			eventHandler.OnAddedAsFriendOperation(result.RawContent.Params)
			break
		case linebot.OpTypeBlocked:
			eventHandler.OnBlockedAccountOperation(result.RawContent.Params)
			break
		default:
			handleContentByType(eventHandler, content)
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
