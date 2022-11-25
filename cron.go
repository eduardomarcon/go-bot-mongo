package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"os"
	"time"
)

func startCron() {
	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	mongoCliente := connectMongo()

	time.Local, _ = time.LoadLocation(os.Getenv("LOCATION_TIMEZONE"))
	s := gocron.NewScheduler(time.Local)
	s.Every(5).Second().Do(func() {
		insert(mongoCliente)

		nowString := time.Now().Format("02-01-2006 15:04:05")
		message := fmt.Sprintf("message sent at %s", nowString)
		if chatId == "" || message == "" {
			log.Fatal("error to send the message")
		}

		fmt.Printf("send message at %s\n", nowString)
		sendTextToTelegramChat(chatId, message)

	})
	s.StartBlocking()
}
