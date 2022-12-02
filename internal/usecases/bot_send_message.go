package usecases

import (
	"bot-mongo/internal/configs"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func sendTextToTelegramChat(text string) error {
	configBOT := configs.GetBOT()
	response, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", configBOT.Token),
		url.Values{
			"chat_id": {configBOT.ChatId},
			"text":    {text},
		})
	defer response.Body.Close()

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return err
	}

	return nil
}
