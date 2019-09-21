package util

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetHuyaStreamData(url string) string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err)
	}
	s := (*doc).Text()
	startIndex := strings.LastIndex(s, "\"stream\"")
	endIndex := strings.LastIndex(s, "window.TT_LIVE_TIMING")
	if startIndex == -1 || endIndex == -1 {
		return ""
	}
	streamJson := s[startIndex+9 : endIndex-12]
	return streamJson
}
