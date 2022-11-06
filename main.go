package main

import (
	"fmt"
	"goServer/api"
)

func main() {
	Curse := api.GetJson()

	fmt.Println(Curse)

	//http.HandleFunc("/", controllers.Sayhello) // Устанавливаем роутер
	//err := http.ListenAndServe(":8080", nil)   // устанавливаем порт веб-сервера
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
}

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello Worls!")
// })
// http.ListenAndServe(":80", nil)
