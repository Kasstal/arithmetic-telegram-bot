package config

import (
	"errors"
	"log"
	"os"
	"time"
)

type Config struct {
	BotToken      string
	TimeOutPoller time.Duration
}

func LoadConfig() (*Config, error) {
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
