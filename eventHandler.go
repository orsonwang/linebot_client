package main

import (
	"log"
//	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/wangbin/jiebago"
)

// LineBotEventHandler ...
type LineBotEventHandler struct {
	botClient *linebot.Client
	seg *jiebago.Segmenter
}

// SetLineBotClient to assign linebot client handler
func (lbe *LineBotEventHandler) SetLineBotClient(bc *linebot.Client) {
	lbe.botClient = bc
}

//InitSegmenter to initial Chinese word segmenter
func (lbe *LineBotEventHandler) InitSegmenter() {
	lbe.seg.LoadDictionary("dict.txt")
}
// OnAddedAsFriendOperation ...
func (lbe *LineBotEventHandler) OnAddedAsFriendOperation(mids []string) {
	lbe.botClient.SendText(mids, "感謝你加入....！")
}

// OnBlockedAccountOperation ...
func (lbe *LineBotEventHandler) OnBlockedAccountOperation(mids []string) {
	lbe.botClient.SendText(mids, "被封鎖了")
}

// OnTextMessage ...
func (lbe *LineBotEventHandler) OnTextMessage(from, text string) {
/*
	chanAfterCut := lbe.seg.Cut(text, false) // 進行精確斷字，斷字結果以空白間隔，後續就可以用它做語意操作
	strAfterCut := "" // 因為Jiebago的輸出是channel，所以先把它轉換成字串陣列
	for strCutMeta := range chanAfterCut {
		strAfterCut += "," + strCutMeta
        }	
*/

// 以下只是一個非智慧型示範，跟以上斷字無關
	strAfterCut := text
    strResult := text
	if strings.Contains(strAfterCut,"利率") {
		if strings.Contains(strAfterCut,"外幣") {
			strResult = "常用外幣利率表\n 美元 定存 2.3% 活存 1.8% \n 日圓 定存 0.1% 活存 0.1%"
		} else {
			strResult = "台幣活存利率表 \n 活存 0.5% 活儲 0.6% \n 定存\n 三個月 0.76% 六個月 0.78% 一年 0.80% 三年 0.80%</P> <A href=\"https://www.skbank.com.tw/RAT/RAT2_TWSaving.aspx\"> 台幣利率</A>"
		}
	}
	if strings.Contains(strAfterCut,"匯率") {
		if strings.Contains(strAfterCut,"歷史") {
			strResult = "</P> <A href=\"https://www.skbank.com.tw/RAT/RAT2_Historys.aspx\"> 歷史匯率</A>"
		} else {
			strResult = "</P> <A href=\"https://www.skbank.com.tw/RAT/RAT2_Foreigns.aspx\"> 外幣匯率</A>"
		}
	} 
        if strings.Contains(strAfterCut,"行動") {
                if strings.Contains(strAfterCut,"應用") || strings.Contains(strings.ToUpper(strAfterCut), "APP") {
                        strResult = "</P> 請參考 <A href=\"https://itunes.apple.com/tw/app/xin-guang-yin-xing/id495872725?l=zh&mt=8\"> 新光銀行行動銀行</A>"
                } else if strings.Contains(strAfterCut,"網頁") {
                        strResult = "</P> 很抱歉，我們還沒有建置行動網頁，請使用<A href=\"https://www.skbank.com.tw/\">網路銀行</A>或<A href=\"https://itunes.apple.com/tw/app/xin-guang-yin-xing/id495872725?l=zh&mt=8\"> 行動銀行</A>"
                }
        }
	
	lbe.botClient.SendText([]string{from}, strResult)
	log.Printf("Received text \"%s\" from %s", text, from)
}

// OnImageMessage ...
func (lbe *LineBotEventHandler) OnImageMessage(from string, rc *linebot.ReceivedContent) {
	lbe.botClient.SendText([]string{from}, "收到一張照片")
	log.Print("=== Received Image ===")
}

// OnVideoMessage ...
func (lbe *LineBotEventHandler) OnVideoMessage(from string, rc *linebot.ReceivedContent) {
	lbe.botClient.SendText([]string{from}, "收到一段錄影")
	log.Print("=== Received Video ===")
}

// OnAudioMessage ...
func (lbe *LineBotEventHandler) OnAudioMessage(from string, duration int) {
	lbe.botClient.SendText([]string{from}, "收到一段錄音")
	log.Print("=== Received Audio ===")
}

// OnLocationMessage ...
func (lbe *LineBotEventHandler) OnLocationMessage(from, title, address string, latitude, longitude float64) {
	lbe.botClient.SendText([]string{from}, "收到地點資訊")
	log.Print("=== Received Location ===")
}

// OnStickerMessage ...
func (lbe *LineBotEventHandler) OnStickerMessage(from string, stickerPackageID, stickerID, stickerVersion int) {
	lbe.botClient.SendText([]string{from}, "收到一張貼紙")
	log.Print("=== Received Sticker ===")
}

// OnContactMessage ...
func (lbe *LineBotEventHandler) OnContactMessage(from, MID, displayName string) {
	lbe.botClient.SendText([]string{from}, "收到聯絡人資料")
	log.Print("=== Received Contact ===")
}
