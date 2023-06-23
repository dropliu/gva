package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// 获取浮动费率
func GetFloatExchangeRate() (r float64, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://bitcoinp2p.org/?spread=1.0080", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	text := doc.Find("title").First().Text()

	if text == "" {
		err = errors.New("获取汇率失败")
		return
	}

	r, err = strconv.ParseFloat(text, 64)
	return
}
