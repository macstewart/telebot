package telegram

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/exp/slog"
)

var client *tgbotapi.BotAPI
var chatID int64

func init() {
	var err error
	chatIdEnv := os.Getenv("TELEGRAM_CHATID")
	if chatID, err = strconv.ParseInt(chatIdEnv, 10, 64); err != nil {
		slog.Error("Failed to parse chatID", "error", err)
		os.Exit(1)
	}
	if client, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN")); err != nil {
		slog.Error("Failed to init", "error", err)
		os.Exit(1)
	}
}

func Notify(message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	if _, err := client.Send(msg); err != nil {
		slog.Error("Failed to send message", "error", err)
		return err
	}
	return nil
}
