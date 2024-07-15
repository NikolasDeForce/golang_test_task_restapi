package main

import (
	"fmt"
	"net/http"
	"os"
	"testtask/handlers"
	"time"
)

var port = ":8010"

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		port = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         port,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/getSettings", http.HandlerFunc(handlers.GetSettingsHandler))
	mux.Handle("/sendMessage", http.HandlerFunc(handlers.SendMessageHandler))
	mux.Handle("/getStateInstance", http.HandlerFunc(handlers.GetStateInstanceHandler))
	mux.Handle("/sendFileByUrl", http.HandlerFunc(handlers.SendFileByUrlHandler))

	fmt.Println("Ready to serve at", port)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
