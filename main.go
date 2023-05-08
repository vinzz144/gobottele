package main

import (
	"log"
	"os"
	"strconv"
	"time"

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

	for {
		now := time.Now()
		targetTime1 := time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, now.Location())
		targetTime2 := time.Date(now.Year(), now.Month(), now.Day(), 13, 0, 0, 0, now.Location())
		targetTime3 := time.Date(now.Year(), now.Month(), now.Day(), 16, 0, 0, 0, now.Location())
		targetTime4 := time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, now.Location())
		targetTime5 := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())

		if now.After(targetTime1) && now.Before(targetTime1.Add(time.Minute)) ||
			now.After(targetTime2) && now.Before(targetTime2.Add(time.Minute)) ||
			now.After(targetTime3) && now.Before(targetTime3.Add(time.Minute)) ||
			now.After(targetTime4) && now.Before(targetTime4.Add(time.Minute)) ||
			now.After(targetTime5) && now.Before(targetTime5.Add(time.Minute)) {

			msg := tgbotapi.NewMessage(chatID, content)
			_, err = bot.Send(msg)
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(time.Minute)
		}

		time.Sleep(time.Second)
	}
}
