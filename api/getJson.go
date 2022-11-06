package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type course struct {
	Date string `json:"Date"`
	//Valute interface{} `json:"Valute"`
	Valute struct {
		Part struct {
			valuta
		} `json:-`
	} `json:"Valute"`
}

type valuta struct {
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
}

// func GetJson() struct{ course } {
func GetJson() course {
	resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	getObject, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//var jsonСurrency map[string]interface{}
	var jsonСurrency course
	if err := json.Unmarshal(getObject, &jsonСurrency); err != nil {
		panic(err)
	}
	//Convert the body to type string
	//sb := string(getObject)
	//log.Printf(jsonСurrency)
	//log.Println(jsonСurrency)
	//gjson.Get(jsonСurrency, "Date")
	//num := jsonСurrency["num"].(float64) // для того, чтобы получить из свойства num число
	//fmt.Println(num)
	return jsonСurrency
}

//resp, err := http.Get("https://www.cbr-xml-daily.ru/daily_json.js")
//if err != nil {
//	log.Fatalln(err)
//}
//body, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//	log.Fatalln(err)
//}
//
//sb := string(body)
//log.Printf(sb)

//get := []byte()
//jsonСurrency := map[string]interface{}
//json.m
//fmt.Println(resp)
//return jsonСurrency
