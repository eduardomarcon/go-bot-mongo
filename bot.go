package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func sendTextToTelegramChat(chatId string, text string) (string, error) {
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	response, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", telegramToken),
		url.Values{
			"chat_id": {chatId},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = io.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
