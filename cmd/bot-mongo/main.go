package main

import (
	"bot-mongo/internal/configs"
	"bot-mongo/internal/usecases"
)

func main() {
	configs.LoadEnvs()
	usecases.RunCronScheduler()
}
