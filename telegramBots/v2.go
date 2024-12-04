package telegramBots

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"os"
)

type Handler interface {
	check(update tgbotapi.Update) bool
	execute(update tgbotapi.Update, bot tgbotapi.BotAPI) bool
	getId() uuid.UUID
}

type BaseHandler struct {
	uuid     uuid.UUID
	callback func(update tgbotapi.Update, bot tgbotapi.BotAPI)
}

func (h BaseHandler) check(_ tgbotapi.Update) bool { return false }

func (h BaseHandler) getId() uuid.UUID {
	return h.uuid
}

type MessageHandler struct {
	BaseHandler
	filters []func(message tgbotapi.Message) bool
}

func (h MessageHandler) check(update tgbotapi.Update) bool {
	if update.Message == nil {
		return false
	}

	for _, filter := range h.filters {
		if !filter(*update.Message) {
			return false
		}
	}

	return true
}

func (h MessageHandler) execute(update tgbotapi.Update, bot tgbotapi.BotAPI) bool {
	if h.check(update) {
		h.callback(update, bot)
		return true
	}

	return false
}

type CommandHandler struct {
	BaseHandler
	cmdText string
}

func (h CommandHandler) check(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.IsCommand() && update.Message.Command() == h.cmdText
}

func (h CommandHandler) execute(update tgbotapi.Update, bot tgbotapi.BotAPI) bool {
	if h.check(update) {
		h.callback(update, bot)
		return true
	}

	return false
}

type ActiveHandlers struct {
	handlers []Handler
}

func (hl ActiveHandlers) handleAll(update tgbotapi.Update, bot tgbotapi.BotAPI) map[uuid.UUID]bool {
	result := make(map[uuid.UUID]bool)

	for _, h := range hl.handlers {
		exec := h.execute(update, bot)
		result[h.getId()] = exec
	}

	return result
}

// ---------------------------------------------------
// ---------------------------------------------------
// ------------END OF TECHNICAL EXPERIMENT------------
// ------------------TEST CODE NEXT-------------------
// ---------------------------------------------------
// ---------------------------------------------------

func HelpCmd(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I understand /sayhi and /status."

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SayHi(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hi :)"

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func StatusCmd(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I'm ok."

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func BotV2Main() {
	// bot init
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Successfully authorized on account @%s", bot.Self.UserName)

	actions := ActiveHandlers{handlers: []Handler{ // All th actions bot will react
		CommandHandler{BaseHandler{uuid.New(), HelpCmd}, "help"},
		CommandHandler{BaseHandler{uuid.New(), SayHi}, "sayhi"},
		CommandHandler{BaseHandler{uuid.New(), StatusCmd}, "status"},
	}}

	// start bot
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// get updates
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		runRes := actions.handleAll(update, *bot)
		fmt.Println("Run results: [ID|called]")
		fmt.Println(runRes)
	}
}
