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

var eventHandler BotEventHandler

func main() {
	strID := os.Getenv("LINE_CHANNEL_ID")
	channelID, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		log.Fatal("Wrong environment setting about LINE_CHANNEL_ID")
	}
	channelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	mid := os.Getenv("LINE_MID")

	botClient, err = linebot.NewClient(channelID, channelSecret, mid)

	// EventHandler
	//	eventHandler = NewEventHandler()
	//	botClient.SetEventHandler(eventHandler)

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
	for _, result := range received.Results {
		content := result.Content()
		log.Println("content:%s\n", content)

		from := content.From
		if content.OpType == linebot.OpTypeAddedAsFriend {
			MIDs := result.RawContent.Params
			eventHandler.OnAddedAsFriendOperation(MIDs)
		}
		if content.OpType == linebot.OpTypeBlocked {
			MIDs := result.RawContent.Params
			eventHandler.OnBlockedAccountOperation(MIDs)
		}
		if content.ContentType == linebot.ContentTypeText {
			from := content.From
			textContent, err := content.TextContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnTextMessage(from, textContent.Text)
		}
		if content.ContentType == linebot.ContentTypeImage {
			eventHandler.OnImageMessage(from)
		}
		if content.ContentType == linebot.ContentTypeVideo {
			eventHandler.OnVideoMessage(from)
		}
		if content.ContentType == linebot.ContentTypeAudio {
			eventHandler.OnAudioMessage(from)
		}
		if content.ContentType == linebot.ContentTypeLocation {
			locationContent, err := content.LocationContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnLocationMessage(from, locationContent.Title, locationContent.Address, locationContent.Latitude, locationContent.Longitude)
		}
		if content.ContentType == linebot.ContentTypeSticker {
			stickerContent, err := content.StickerContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnStickerMessage(from, stickerContent.ID, stickerContent.PackageID, stickerContent.Version)
		}
		if content.ContentType == linebot.ContentTypeContact {
			contactContent, err := content.ContactContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnContactMessage(from, contactContent.Mid, contactContent.DisplayName)
		}
	}

}

// BotEventHandler ...
type BotEventHandler struct {
}

// NewEventHandler ...
func NewEventHandler() *BotEventHandler {
	return &BotEventHandler{}
}

// OnAddedAsFriendOperation ...
func (be *BotEventHandler) OnAddedAsFriendOperation(mids []string) {
	botClient.SendText(mids, "友達追加してくれてありがとうね！")
}

// OnBlockedAccountOperation ...
func (be *BotEventHandler) OnBlockedAccountOperation(mids []string) {
	botClient.SendText(mids, "あらら,,, (このメッセージは届かない)")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(from, text string) {
	log.Print("=== Received Text ===")
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(from string) {
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(from string) {
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(from string) {
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (be *BotEventHandler) OnContactMessage(from, MID, displayName string) {
	log.Print("=== Received Contact ===")
}
