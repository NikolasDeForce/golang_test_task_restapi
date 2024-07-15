package main

import (
	"fmt"
	"net/http"
	"os"
	"testtask/handlers"
	"time"
)

var port = ":8010"

var idInstance = "1234567890"
var apiTokenInstance = "A123p456I789"

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

	mux.Handle(fmt.Sprintf("/waInstance%v/getSettings/%v", idInstance, apiTokenInstance), http.HandlerFunc(handlers.GetSettingsHandler))
	mux.Handle(fmt.Sprintf("/waInstance%v/sendMessage/%v", idInstance, apiTokenInstance), http.HandlerFunc(handlers.SendMessageHandler))
	mux.Handle(fmt.Sprintf("/waInstance%v/getStateInstance/%v", idInstance, apiTokenInstance), http.HandlerFunc(handlers.GetStateInstanceHandler))
	mux.Handle(fmt.Sprintf("/waInstance%v/sendFileByUrl/%v", idInstance, apiTokenInstance), http.HandlerFunc(handlers.SendFileByUrlHandler))

	fmt.Println("Ready to serve at", port)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
