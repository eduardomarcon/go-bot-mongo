package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var cfg *config

type config struct {
	DB  DBConfig
	BOT BOTConfig
	TZ  TZConfig
}

type DBConfig struct {
	URI string
}

type BOTConfig struct {
	Token  string
	ChatId string
}

type TZConfig struct {
	TZ string
}

func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	cfg = new(config)
	cfg.DB = DBConfig{
		URI: os.Getenv("MONGODB_URI"),
	}
	cfg.BOT = BOTConfig{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		ChatId: os.Getenv("TELEGRAM_CHAT_ID"),
	}
	cfg.TZ = TZConfig{
		TZ: os.Getenv("LOCATION_TIMEZONE"),
	}
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetBOT() BOTConfig {
	return cfg.BOT
}

func GetTZ() TZConfig {
	return cfg.TZ
}
