package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// wakatimeから今日のコーディング時間合計を取得
// jsonデータ例
// [{"grand_total": {"digital": "2:11", "hours": 2, "minutes": 11, "text": "2 hrs 11 mins", "total_seconds": 7910.190419},
//  "range": {"date": "2020-08-24", "end": "2020-08-24T14:59:59Z", "start": "2020-08-23T15:00:00Z",
//  "text": "Mon Aug 24th 2020", "timezone": "Asia/Tokyo"}},....]
type GrandTotal struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}

type Range struct {
	Date     string    `json:"date"`
	End      time.Time `json:"end"`
	Start    time.Time `json:"start"`
	Text     string    `json:"text"`
	Timezone string    `json:"timezone"`
}

type Data []struct {
	GrandTotal GrandTotal `json:"grand_total"`
	Range Range `json:"range"`
}

type Wakatime struct {
	Data Data `json:"data"`
}

func readJSONFromUrl(url string) (Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(Wakatime)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}
	return data.Data, nil
}

func fetchCodingTimeFromJSON(wakatimeData Data) (string, error) {
	t := time.Now().Format("2006-01-02")
	for _, d := range wakatimeData {
		if (d.Range.Date == t) {
			return d.GrandTotal.Text, nil
		}
	}	
	return "", nil
}

func GetTodaysCodingTime() (string, error) {
	// NOTE: This is public URL
	url := "https://wakatime.com/share/@30b97a12-9063-408a-8939-31ffa7cc1987/f8ab89d3-0b65-423b-ad5c-02d9a87f1d8a.json"
	wakatimeData, err := readJSONFromUrl(url)
	if err != nil {
		return "", err
	}
	codingTime, err := fetchCodingTimeFromJSON(wakatimeData)
	if err != nil {
		return "", err
	}	
	return codingTime, nil
}
