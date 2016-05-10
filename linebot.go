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
		//		log.Println("content:%s\n", content)

		if content.OpType == linebot.OpTypeAddedAsFriend {
			MIDs := result.RawContent.Params
			eventHandler.OnAddedAsFriendOperation(MIDs)
		}
		if content.OpType == linebot.OpTypeBlocked {
			MIDs := result.RawContent.Params
			eventHandler.OnBlockedAccountOperation(MIDs)
		}
		if content.ContentType == linebot.ContentTypeText {
			textContent, err := content.TextContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnTextMessage(content.From, textContent.Text)
		}
		if content.ContentType == linebot.ContentTypeImage {
			imageContent, err := content.ImageContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnImageMessage(content.From)
		}
		if content.ContentType == linebot.ContentTypeVideo {
			videoContent, err := content.VideoContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnVideoMessage(content.From)
		}
		if content.ContentType == linebot.ContentTypeAudio {
			audioContent, err := content.AudioContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnAudioMessage(content.From, audioContent.Duration)
		}
		if content.ContentType == linebot.ContentTypeLocation {
			locationContent, err := content.LocationContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnLocationMessage(content.From, locationContent.Title, locationContent.Address, locationContent.Latitude, locationContent.Longitude)
		}
		if content.ContentType == linebot.ContentTypeSticker {
			stickerContent, err := content.StickerContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnStickerMessage(content.From, stickerContent.ID, stickerContent.PackageID, stickerContent.Version)
		}
		if content.ContentType == linebot.ContentTypeContact {
			contactContent, err := content.ContactContent()
			if err != nil {
				log.Print(err)
				return
			}
			eventHandler.OnContactMessage(content.From, contactContent.Mid, contactContent.DisplayName)
		}
	}

}

// BotEventHandler ...
type BotEventHandler struct {
}

// OnAddedAsFriendOperation ...
func (be *BotEventHandler) OnAddedAsFriendOperation(mids []string) {
	botClient.SendText(mids, "感謝你加入....！")
}

// OnBlockedAccountOperation ...
func (be *BotEventHandler) OnBlockedAccountOperation(mids []string) {
	botClient.SendText(mids, "被封鎖了")
}

// OnTextMessage ...
func (be *BotEventHandler) OnTextMessage(from, text string) {
	botClient.SendText([]string{from}, text)
	log.Print("Received text \"%s\" from %s", text, from)
}

// OnImageMessage ...
func (be *BotEventHandler) OnImageMessage(from string) {
	botClient.SendText([]string{from}, "收到一張照片")
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (be *BotEventHandler) OnVideoMessage(from string) {
	botClient.SendText([]string{from}, "收到一段錄影")
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (be *BotEventHandler) OnAudioMessage(from string, duration int) {
	botClient.SendText([]string{from}, "收到一段錄音")
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (be *BotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	botClient.SendText([]string{from}, "收到地點資訊")
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (be *BotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	botClient.SendText([]string{from}, "收到一張貼紙")
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (be *BotEventHandler) OnContactMessage(from, MID, displayName string) {
	botClient.SendText([]string{from}, "收到聯絡人資料")
	log.Print("=== Received Contact ===")
}
