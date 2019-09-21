package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"web/util"
)

const host = "http://www.huya.com/"

type gameStreamInfo struct {
	SHlsUrl       string `json:"sHlsUrl"`
	SStreamName   string `json:"sStreamName"`
	SHlsUrlSuffix string `json:"sHlsUrlSuffix"`
}

type Data struct {
	GameStreamInfo []gameStreamInfo `json:"gameStreamInfoList"`
}

type StreamData struct {
	Status int    `json:"status"`
	Data   []Data `json:"data"`
}

type CategoryInfo struct {
	RoomInfo []RoomInfo `json:"roomInfo"`
}

type RoomInfo struct {
	RoomId string `json:"roomId"`
	Title  string `json:"title"`
	Nick   string `json:"nick"`
}

func GetStreamData(roomId string, streamBit string) []string {
	streamJson := util.GetHuyaStreamData(host + roomId)
	var streamData StreamData
	e := json.Unmarshal([]byte(streamJson), &streamData)
	if e != nil {
		log.Println(e)
		return nil
	}
	ss := make([]string, 0)
	for _, v := range streamData.Data {
		for _, v1 := range v.GameStreamInfo {
			if v1.SHlsUrl != "" && v1.SHlsUrlSuffix != "" && v1.SStreamName != "" {
				var url string
				if streamBit != "" {
					url = fmt.Sprint(v1.SHlsUrl, "/", v1.SStreamName, "_", streamBit, ".", v1.SHlsUrlSuffix)
				} else {
					url = fmt.Sprint(v1.SHlsUrl, "/", v1.SStreamName, ".", v1.SHlsUrlSuffix)
				}
				ss = append(ss, url)
			}
		}
	}
	return ss
}

// /g/lol /g/seeTogether
func GetHuyaLiveCategory(action string) *CategoryInfo {
	lol, err := goquery.NewDocument(host + action)
	if err != nil {
		log.Println(err)
		return nil
	}
	gameNode := (*lol).Find(".game-live-item")
	roomInfos := make([]RoomInfo, 0)
	gameNode.Each(func(i int, selection *goquery.Selection) {
		roomUrl, exist := selection.Find("a").First().Attr("href")
		if exist {
			roomId := getRoomId(roomUrl)
			title := selection.Find(".title").Text()
			nick := selection.Find(".nick").Text()
			roomInfos = append(roomInfos, RoomInfo{roomId, title, nick})
		}
	})
	return &CategoryInfo{roomInfos}
}

func getRoomId(roomUrl string) string {
	index := strings.LastIndex(roomUrl, "/")
	return roomUrl[index+1:]
}
