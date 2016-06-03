package main

import (
	//	logger "log"
	"regexp"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotEventHandler ...
type LineBotEventHandler struct {
	botClient *linebot.Client
}

// SetLineBotClient to assign linebot client handler
func (s *LineBotEventHandler) SetLineBotClient(bc *linebot.Client) {
	s.botClient = bc
}

// OnAddedAsFriendOperation ...
func (s *LineBotEventHandler) OnAddedAsFriendOperation(mids []string) {
	s.botClient.SendText(mids, "感謝你加入....！")
}

// OnBlockedAccountOperation ...
func (s *LineBotEventHandler) OnBlockedAccountOperation(mids []string) {
	s.botClient.SendText(mids, "被封鎖了")
}

func (s *LineBotEventHandler) matchString(pattern, text string) (result bool) {
	result, _ = regexp.MatchString(pattern, text)
	return
}

// OnTextMessage ...
func (s *LineBotEventHandler) OnTextMessage(from, text string) {
	strAfterCut := strings.ToUpper(text)
	log.Printf("Received text \"%s\" from %s", text, from)

	subj := "aitc.text.service"
	msg, err := nc.Request(subj, []byte(strAfterCut), 1*time.Second)
	if err != nil {
		log.Fatalf("Error in Request: %v\n", err)
	}
	strResult := string(msg.Data)
	if len(strResult) != 0 {
		s.botClient.SendText([]string{from}, strResult)
	} else {
		s.botClient.SendImage([]string{from},
			"https://linebot.gaze.tw/exrate.png",
			"https://linebot.gaze.tw/exrate.png")
	}
}

// OnImageMessage ...
func (s *LineBotEventHandler) OnImageMessage(from string, rc *linebot.ReceivedContent) {
	s.botClient.SendText([]string{from}, "收到一張照片")
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (s *LineBotEventHandler) OnVideoMessage(from string, rc *linebot.ReceivedContent) {
	s.botClient.SendText([]string{from}, "收到一段錄影")
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (s *LineBotEventHandler) OnAudioMessage(from string, duration int) {
	s.botClient.SendText([]string{from}, "收到一段錄音")
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (s *LineBotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	s.botClient.SendText([]string{from}, "收到地點資訊")
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (s *LineBotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	s.botClient.SendText([]string{from}, "收到一張貼紙")
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (s *LineBotEventHandler) OnContactMessage(from, MID, displayName string) {
	s.botClient.SendText([]string{from}, "收到聯絡人資料")
	log.Print("=== Received Contact ===")
}
