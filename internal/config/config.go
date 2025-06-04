package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	BotToken      string        `env:"BOT_TOKEN"`
	TimeOutPoller time.Duration `env:"TIMEOUT_POLLER"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	botToken := os.Getenv("BOT_TOKEN")
	timeOutPoller := os.Getenv("TIMEOUT_POLLER")
	if botToken == "" || timeOutPoller == "" {
		log.Printf("Необходимые переменные окружения отсутствуют", "BOT_TOKEN", botToken, "TIMEOUT_POLLER", timeOutPoller)
		return nil, errors.New("необходимые переменные окружения отсутствуют")
	}
	timeout, err := time.ParseDuration(timeOutPoller)
	if err != nil {
		log.Printf("Ошибка при парсинге TIMEOUT_POLLER: %v", err)
		return nil, errors.New("неверный формат TIMEOUT_POLLER")
	}
	return &Config{
		BotToken:      botToken,
		TimeOutPoller: timeout,
	}, nil
}
