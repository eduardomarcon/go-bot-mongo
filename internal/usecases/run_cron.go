package usecases

import (
	"bot-mongo/internal/configs"
	"bot-mongo/internal/repositories"
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

func RunCronScheduler() {
	time.Local, _ = time.LoadLocation(configs.GetTZ().TZ)

	shedule := gocron.NewScheduler(time.Local)
	shedule.Every(5).Second().Do(func() {
		todoRepository := repositories.NewTodoRepository()
		todoRepository.Insert()

		nowString := time.Now().Format("02-01-2006 15:04:05")
		message := fmt.Sprintf("message sent at %s", nowString)
		err := sendTextToTelegramChat(message)
		if err != nil {
			fmt.Println(err)
		}

	})
	shedule.StartBlocking()
}
