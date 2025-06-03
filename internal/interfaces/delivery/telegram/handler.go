package telegramHandler

import (
	"arithmetic-telegram-bot/internal/app/usecase"
	"arithmetic-telegram-bot/internal/domain"
	"arithmetic-telegram-bot/internal/infrastructure/telegram"
	"fmt"
	"gopkg.in/telebot.v4"
	"log"
)

type TelegramHandler struct {
	calculatorSvc  usecase.CalculatorService
	telegramClient telegramClient.TelegramClient
}

func NewTelegramHandler(calculatorSvc usecase.CalculatorService,
	telegramClient telegramClient.TelegramClient,
) *TelegramHandler {
	return &TelegramHandler{
		calculatorSvc:  calculatorSvc,
		telegramClient: telegramClient,
	}
}
func (h *TelegramHandler) HandleMessage(c telebot.Context) error {

	expr := c.Message().Text // Получаем текст из telebot.Context
	log.Printf("[%s] Received expression: %s", c.Sender().Username, expr)

	result, calcErr := h.calculatorSvc.Calculate(expr)
	var responseText string
	if calcErr != nil {
		switch calcErr.Type {
		case domain.ErrorTypeInvalidExpression:
			responseText = fmt.Sprintf("❌ Ошибка синтаксиса: %s\nПожалуйста, проверьте выражение. Пример: `2 * (5 + 3)`", calcErr.Message)
		case domain.ErrorTypeDivisionByZero:
			responseText = fmt.Sprintf("❌ Ошибка: %s", calcErr.Message)
		case domain.ErrorTypeUnknown:
			responseText = fmt.Sprintf("❌ Произошла неизвестная ошибка: %s", calcErr.Message)
		default:
			responseText = fmt.Sprintf("❌ Произошла ошибка: %s", calcErr.Message)
		}
	} else {
		responseText = fmt.Sprintf("✅ Результат: `%.2f`", result)
	}

	err := h.telegramClient.SendMessage(c.Chat().ID, responseText)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	return nil
}

func (h *TelegramHandler) HandleStartCommand(c telebot.Context) error {
	err := h.telegramClient.SendMessage(c.Chat().ID, "Привет! Я бот-калькулятор. Просто отправь мне арифметическое выражение, и я его вычислю. Например: `2 + 2 * (3 - 1)`")
	if err != nil {
		log.Printf("Failed to send start message: %v", err)
		return err
	}
	return nil
}

func (h *TelegramHandler) HandleHelpCommand(c telebot.Context) error {
	err := h.telegramClient.SendMessage(c.Chat().ID, "Я умею вычислять простые арифметические выражения. Поддерживаю `+`, `-`, `*`, `/` и скобки. Пример: `(10 + 5) / 3`")
	if err != nil {
		log.Printf("Failed to send help message: %v", err)
		return err
	}
	return nil
}
