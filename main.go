package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatIDString := os.Getenv("TELEGRAM_CHAT_ID")

	chatID, err := strconv.ParseInt(chatIDString, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	content := "🔔kindly reminder:\n" +
		"Jangan lupa minum air putih\n" +
		"untuk umur yang lebih panjang+\n" +
		"🙏😌\n"

	msg := tgbotapi.NewMessage(chatID, content)
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
