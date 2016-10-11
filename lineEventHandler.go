package main

import (
	//	logger "log"
	"regexp"
	//	"strings"
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

func (s *LineBotEventHandler) matchString(pattern, text string) (result bool) {
	result, _ = regexp.MatchString(pattern, text)
	return
}

// OnTextMessage ...
func (s *LineBotEventHandler) OnTextMessage(from, text string) {
	log.Printf("Received text \"%s\" from %s", text, from)

	subj := "aitc.text.service"
	msg, err := nc.Request(subj, []byte(text), 1*time.Second)
	if err != nil {
		log.Fatalf("Error in Request: %v\n", err)
	}
	strResult := string(msg.Data)
	if len(strResult) != 0 {
		if _, err = s.botClient.ReplyMessage(from, linebot.NewTextMessage(strResult)).Do(); err != nil {
			log.Printf("Send message \"%s\" to \"%s\" fail, error:%d", strResult, from, err)
		}
	} else {

	}
}

/*
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
*/
