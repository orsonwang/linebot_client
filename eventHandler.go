package linebot_client

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BotEventHandler ...
type LineBotEventHandler struct {
	BotClient *linebot.Client
}

// Assign linebot client handler
func (lbe *LineBotEventHandler) setLineBotClient(bc *linebot.Client) {
	lbe.BotClient = bc
}

// OnAddedAsFriendOperation ...
func (lbe *LineBotEventHandler) OnAddedAsFriendOperation(mids []string) {
	lbe.BotClient.SendText(mids, "感謝你加入....！")
}

// OnBlockedAccountOperation ...
func (lbe *LineBotEventHandler) OnBlockedAccountOperation(mids []string) {
	lbe.BotClient.SendText(mids, "被封鎖了")
}

// OnTextMessage ...
func (lbe *LineBotEventHandler) OnTextMessage(from, text string) {
	lbe.BotClient.SendText([]string{from}, text)
	log.Printf("Received text \"%s\" from %s", text, from)
}

// OnImageMessage ...
func (lbe *LineBotEventHandler) OnImageMessage(from string, rc *linebot.ReceivedContent) {
	lbe.BotClient.SendText([]string{from}, "收到一張照片")
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (lbe *LineBotEventHandler) OnVideoMessage(from string, rc *linebot.ReceivedContent) {
	lbe.BotClient.SendText([]string{from}, "收到一段錄影")
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (lbe *LineBotEventHandler) OnAudioMessage(from string, duration int) {
	lbe.BotClient.SendText([]string{from}, "收到一段錄音")
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (lbe *LineBotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	lbe.BotClient.SendText([]string{from}, "收到地點資訊")
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (lbe *LineBotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	lbe.BotClient.SendText([]string{from}, "收到一張貼紙")
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (lbe *LineBotEventHandler) OnContactMessage(from, MID, displayName string) {
	lbe.BotClient.SendText([]string{from}, "收到聯絡人資料")
	log.Print("=== Received Contact ===")
}
