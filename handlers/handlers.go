package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testtask/methods"
)

var getSettings = methods.GetSettings{}
var getStateInstance = methods.GetStateInstance{}
var sendMessage = methods.SendMessage{}
var sendFileByUrl = methods.SendFileByUrl{}

func GetSettingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &getSettings)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(getSettings)
}

func GetStateInstanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &getStateInstance)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(getSettings)
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &sendMessage)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(sendMessage)
}

func SendFileByUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &sendFileByUrl)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(sendFileByUrl)
}
