package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Сourse struct {
	Date   string            `json:"Date"`
	Valute map[string]Valuta `json:"Valute"`
}

type Valuta struct {
	//Data     string
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
}

func GetJson() *Сourse {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		log.Fatalln(err)
	}
	getObject, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var jsonСurrency Сourse
	if err := json.Unmarshal(getObject, &jsonСurrency); err != nil {
		panic(err)
	}
	return &jsonСurrency
}
