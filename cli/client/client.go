package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testtask/methods"
	"time"
)

var GetSettings = methods.GetSettings{
	Wid:                               "11001234567@c.us",
	CountryInstance:                   "",
	TypeAccount:                       "",
	WebhookUrl:                        "https://mysite.com/webhook/green-api/",
	WebhookUrlToken:                   "",
	DelaySendMessagesMilliseconds:     "5000",
	MarkIncomingMessagesReaded:        "no",
	MarkIncomingMessagesReadedOnReply: "no",
	SharedSession:                     "no",
	OutgoingWebhook:                   "yes",
	OutgoingMessageWebhook:            "yes",
	OutgoingAPIMessageWebhook:         "yes",
	IncomingWebhook:                   "yes",
	DeviceWebhook:                     "no",
	StatusInstanceWebhook:             "no",
	StateWebhook:                      "no",
	EnableMessagesHistory:             "no",
	KeepOnlineStatus:                  "no",
	PollMessageWebhook:                "no",
	IncomingBlockWebhook:              "yes",
	IncomingCallWebhook:               "yes",
}

var GetStateInstance = methods.GetStateInstance{
	StateInstance: "authorized",
}

var SendMessage = methods.SendMessage{
	ChatId:          "11001234567@c.us",
	Message:         "Hello from Russia",
	QuotedMessageId: "",
	LinkPreview:     false,
}

var SendFileByUrl = methods.SendFileByUrl{
	ChatId:   "11001234567@c.us",
	UrlFile:  "https://my.site.com/img/horse.png",
	FileName: "horse.png",
	Caption:  "Лошадка",
}

func GetSettingsHandler(host string, server string, idInstance string, apiTokenInstance string, g methods.GetSettings) int {
	getSettingsMarshall, _ := json.Marshal(g)
	u := bytes.NewReader(getSettingsMarshall)

	req, err := http.NewRequest("GET", host+server+fmt.Sprintf("/waInstance%v/getSettings/%v", idInstance, apiTokenInstance), u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil {
		return http.StatusNotFound
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode
}

func GetStateInstanceHangler(host string, server string, idInstance string, apiTokenInstance string, g methods.GetStateInstance) int {
	getStateInstanceMarshall, _ := json.Marshal(g)
	u := bytes.NewReader(getStateInstanceMarshall)

	req, err := http.NewRequest("GET", host+server+fmt.Sprintf("/waInstance%v/getStateInstance/%v", idInstance, apiTokenInstance), u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil {
		return http.StatusNotFound
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode
}

func SendMessageHandler(host string, server string, idInstance string, apiTokenInstance string, s methods.SendMessage) int {
	sendMessageMarshall, _ := json.Marshal(s)
	u := bytes.NewReader(sendMessageMarshall)

	req, err := http.NewRequest("POST", host+server+fmt.Sprintf("/waInstance%v/sendMessage/%v", idInstance, apiTokenInstance), u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}

	return resp.StatusCode
}

func SendFileByUrlHandler(host string, server string, idInstance string, apiTokenInstance string, s methods.SendFileByUrl) int {
	sendFileByUrlMarshall, _ := json.Marshal(s)
	u := bytes.NewReader(sendFileByUrlMarshall)

	req, err := http.NewRequest("POST", host+server+fmt.Sprintf("/waInstance%v/sendFileByUrl/%v", idInstance, apiTokenInstance), u)
	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}

	return resp.StatusCode
}

func main() {
	host := "http://localhost"
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server")
		return
	}
	server := os.Args[1]

	resp := GetSettingsHandler(host, server, "1234567890", "A123p456I789", GetSettings)
	fmt.Println("/getSettings return code:", resp)

	resp = GetStateInstanceHangler(host, server, "1234567890", "A123p456I789", GetStateInstance)
	fmt.Println("/getStateInstance return code:", resp)

	resp = SendMessageHandler(host, server, "1234567890", "A123p456I789", SendMessage)
	if resp != http.StatusOK {
		fmt.Println("Err code:", resp)
	} else {
		fmt.Println("Data returns", SendMessage, "StatusCode", resp)
	}

	resp = SendFileByUrlHandler(host, server, "1234567890", "A123p456I789", SendFileByUrl)
	if resp != http.StatusOK {
		fmt.Println("Err code:", resp)
	} else {
		fmt.Println("Data returns", SendFileByUrl, "StatusCode", resp)
	}
}
