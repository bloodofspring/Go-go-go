package telegramBots

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"log"
	"os"
)

type Filter func(update tgbotapi.Update) bool

type Callback interface {
	run(update tgbotapi.Update) error
	getName() string
}

type Handler interface {
	checkType(update tgbotapi.Update) bool
	checkFilters(update tgbotapi.Update) bool
	run(update tgbotapi.Update) (bool, error)
	getId() uuid.UUID
}

type BaseHandler struct {
	uuid      uuid.UUID
	queryType string
	callback  Callback
	filters   []Filter
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

func (h BaseHandler) run(update tgbotapi.Update) (bool, error) {
	if h.checkType(update) && h.checkFilters(update) {
		return true, h.callback.run(update)
	}

	return false, nil
}

type ActiveHandlers struct {
	handlers []Handler
}

func (hl ActiveHandlers) handleAll(update tgbotapi.Update) map[uuid.UUID]bool {
	result := make(map[uuid.UUID]bool)

	for _, h := range hl.handlers {
		runResult, err := h.run(update)

		if err != nil {
			panic(err)
		}

		result[h.getId()] = runResult
	}

	return result
}

type handlerProducer struct {
	handlerType string
}

func (p handlerProducer) product(callback Callback, filters []Filter) BaseHandler {
	return BaseHandler{
		uuid:      uuid.New(),
		queryType: p.handlerType,
		callback:  callback,
		filters:   filters,
	}
}

var MessageHandler = handlerProducer{"message"}
var CommandHandler = handlerProducer{"command"}
var CallbackQueryHandler = handlerProducer{"callbackQuery"}

// ---------------------------------------------------
// ---------------------------------------------------
// ---------------END OF TECHNICAL PART---------------
// ---------------------------------------------------
// ---------------------------------------------------

type SayHi struct {
	name   string
	client tgbotapi.BotAPI
}

func (e SayHi) fabricateAnswer(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Hi :)"

	return msg
}

func (e SayHi) run(update tgbotapi.Update) error {
	if _, err := e.client.Send(e.fabricateAnswer(update)); err != nil {
		return err
	}

	return nil
}

func (e SayHi) getName() string {
	return e.name
}

func BotV2Main() {
	// bot init
	bot, err := tgbotapi.NewBotAPI(os.Getenv("API_KEY"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Successfully authorized on account @%s", bot.Self.UserName)

	startFilter := func(update tgbotapi.Update) bool { return update.Message.Command() == "start" }

	actions := ActiveHandlers{handlers: []Handler{ // All th actions bot will react
		CommandHandler.product(SayHi{"start-cmd", *bot}, []Filter{startFilter}),
	}}

	// start bot
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// get updates
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		runRes := actions.handleAll(update)
		fmt.Println("Run results: [ID|called]")
		fmt.Println(runRes)
	}
}
