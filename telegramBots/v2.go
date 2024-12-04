package telegramBots

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"os"
)

type Handler interface {
	checkType(update tgbotapi.Update) bool
	checkFilters(update tgbotapi.Update) bool
	run(update tgbotapi.Update, bot tgbotapi.BotAPI) bool
	getId() uuid.UUID
}

type BaseHandler struct {
	uuid      uuid.UUID
	queryType string
	callback  func(client tgbotapi.BotAPI, update tgbotapi.Update) bool
	filters   []func(update tgbotapi.Update) bool
}

func (h BaseHandler) getId() uuid.UUID {
	return h.uuid
}

func (h BaseHandler) checkType(update tgbotapi.Update) bool {
	switch h.queryType {
	case "message":
		return update.Message != nil
	case "callbackQuery":
		return update.CallbackQuery != nil
	case "command":
		return update.Message != nil && update.Message.IsCommand()
	default:
		fmt.Printf("WARNING! Unsupported query type: %s\n", h.queryType)
		return false
	}
}

func (h BaseHandler) checkFilters(update tgbotapi.Update) bool {
	for _, f := range h.filters {
		if !f(update) {
			return false
		}
	}

	return true
}

func (h BaseHandler) run(update tgbotapi.Update, bot tgbotapi.BotAPI) bool {
	if h.checkType(update) && h.checkFilters(update) {
		return h.callback(bot, update)
	}

	return false
}

type ActiveHandlers struct {
	handlers []Handler
}

func (hl ActiveHandlers) handleAll(update tgbotapi.Update, bot tgbotapi.BotAPI) map[uuid.UUID]bool {
	result := make(map[uuid.UUID]bool)

	for _, h := range hl.handlers {
		runResult := h.run(update, bot)
		result[h.getId()] = runResult
	}

	return result
}

func MessageHandler(callback func(client tgbotapi.BotAPI, update tgbotapi.Update) bool, filters []func(update tgbotapi.Update) bool) BaseHandler {
	return BaseHandler{uuid.New(), "message", callback, filters}
}
func CommandHandler(callback func(client tgbotapi.BotAPI, update tgbotapi.Update) bool, filters []func(update tgbotapi.Update) bool) BaseHandler {
	return BaseHandler{uuid.New(), "command", callback, filters}
}
func CallbackQueryHandler(callback func(client tgbotapi.BotAPI, update tgbotapi.Update) bool, filters []func(update tgbotapi.Update) bool) BaseHandler {
	return BaseHandler{uuid.New(), "callbackQuery", callback, filters}
}

// ---------------------------------------------------
// ---------------------------------------------------
// ---------------END OF TECHNICAL PART---------------
// ---------------------------------------------------
// ---------------------------------------------------

func HelpCmd(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "I understand /sayhi and /status."

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SayHi(client tgbotapi.BotAPI, update tgbotapi.Update) bool {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hi :)"

	if _, err := client.Send(msg); err != nil {
		return false
	}

	return true
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
		CommandHandler(SayHi, []func(update tgbotapi.Update) bool{func(update tgbotapi.Update) bool { return update.Message.Command() == "start" }}),
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
