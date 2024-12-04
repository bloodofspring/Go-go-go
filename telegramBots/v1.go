package telegramBots

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func BotV1Main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		newMsg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		newMsg.ReplyToMessageID = update.Message.MessageID
		if _, err := bot.Send(newMsg); err != nil {
			panic(err)
		}
	}
}
