package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"goServer/api"
	"goServer/internal/app/apiserver"
	"log"
	"time"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "pathto  config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
	//upadateCurse()
	//fmt.Println(db.DataPush(api.GetJson()))
	//DataPush()
	//http.HandleFunc("/", httpserver.Sayhello) // Устанавливаем роутер
	//err := http.ListenAndServe(":8080", nil)  // устанавливаем порт веб-сервера
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}

func upadateCurse() {
	for {
		fmt.Println(api.GetJson())
		time.Sleep(time.Hour)
	}
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Worls!")
// })
// http.ListenAndServe(":80", nil)
