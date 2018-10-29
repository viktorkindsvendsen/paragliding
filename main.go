package main

import (
	"log"
	"net/http"
	"time"
)


func routes(response http.ResponseWriter, request *http.Request) {
	log.Println("Hei ", request.Host, request.URL.Path)
	switch urlEnding := request.URL.Path; urlEnding {
	case "/" :
		print("ERROR 404")
	case "/bin" :
		log.Print("Error 404")
	default:
		log.Print("Default error")
	}


}

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.HandlerFunc(routes),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	//time.AfterFunc(1*time.Second, makeRequest)
	log.Println("Serve at ", server.Addr)
	log.Panic(server.ListenAndServe())
}




