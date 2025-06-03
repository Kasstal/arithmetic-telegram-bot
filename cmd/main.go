package main

import (
	"arithmetic-telegram-bot/internal/app/usecase"
	"arithmetic-telegram-bot/internal/config"
	"arithmetic-telegram-bot/internal/infrastructure/evaluator"
	telegramClient "arithmetic-telegram-bot/internal/infrastructure/telegram"
	telegramHandler "arithmetic-telegram-bot/internal/interfaces/delivery/telegram"
	"fmt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("ошибка загрузки конфигурации:", err)
		return
	}

	evaluatorRepo := evaluator.NewGoEvaluator()

	calculatorSvc := usecase.NewCalculatorService(evaluatorRepo)

	botClient, err := telegramClient.NewTelegramBot(cfg.BotToken, cfg.TimeOutPoller)
	if err != nil {
		fmt.Println("ошибка создания Telegram клиента:", err)
		return
	}
	handler := telegramHandler.NewTelegramHandler(calculatorSvc, botClient)

	botClient.RegisterHandler(telegramClient.OnText, handler.HandleMessage)
	botClient.RegisterHandler("/start", handler.HandleStartCommand)
	botClient.RegisterHandler("/help", handler.HandleHelpCommand)

	botClient.Start()
}
