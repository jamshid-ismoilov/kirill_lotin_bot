package handler

import (
	"log"

	"github.com/Jamshid-Ismoilov/kirill_lotin/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// BotState represents the state of the user's interaction with the bot
type BotState int

const (
	StateIdle BotState = iota
	StateWaitingForLatinText
	StateWaitingForCyrillicText
	StateWaitingForEchoText
)

var currentState = make(map[int64]BotState) // Mapping of chat IDs to their respective states

func RunTgBot(botToken string) {
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set up updates configuration
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// Get updates from the bot
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Process updates
	for update := range updates {
		// Check if the update has a command
		if update.Message != nil && update.Message.IsCommand() {
			handleCommand(bot, update.Message)
		} else if update.Message != nil {
			handleMessage(bot, update.Message)
		}
	}
}

// Handle commands
func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		handleStartCommand(bot, message)
	case "kirill":
		handleCyrillicCommand(bot, message)
	case "lotin":
		handleLatinCommand(bot, message)
	case "cyrillic":
		handleCyrillicCommand(bot, message)
	case "latin":
		handleLatinCommand(bot, message)
	default:
		reply := `Bizda bunday buyruq mavjud emas :(, mavjud buyruqlar: /start /lotin /latin /kirill /cyrillic 
		Бизда бундай буйруқ мавжуд эмас :(, мавжуд буйруқлар: /start /lotin /latin /kirill /cyrillic`
		sendMessage(bot, message.Chat.ID, reply)
	}
}

// Handle the /start command
func handleStartCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := "Assalomu alaykum, " + message.From.FirstName + "!\nKirill-lotin bot hizmatingizda @kirill_lotin_iobot\n\nAссалому алайкум, " + message.From.FirstName + "!\nКирилл-лотин бот ҳизматингизда @kirill_lotin_iobo"
	sendMessage(bot, message.Chat.ID, reply)
}

// Handle the /latin command
func handleLatinCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := "Lotincha matnni yuboring"
	sendMessage(bot, message.Chat.ID, reply)
	currentState[message.Chat.ID] = StateWaitingForLatinText
}

// Handle the /cyrillic command
func handleCyrillicCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := "Кириллча матнни юборинг"
	sendMessage(bot, message.Chat.ID, reply)
	currentState[message.Chat.ID] = StateWaitingForCyrillicText
}

// Handle regular messages
func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	state := currentState[message.Chat.ID]

	switch state {
	case StateWaitingForCyrillicText:
		handleCyrillicMessage(bot, message)
	case StateWaitingForLatinText:
		handleLatinMessage(bot, message)
	default:
		// Default behavior
		handleSimpleMessage(bot, message)
	}
}

// Handle the cyrillic message
func handleCyrillicMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := utils.CyrillicToLatin(message.Text)
	sendMessage(bot, message.Chat.ID, reply)
	currentState[message.Chat.ID] = StateIdle
}

// Handle the latin message
func handleLatinMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	reply := utils.LatinToCyrillic(message.Text)
	sendMessage(bot, message.Chat.ID, reply)
	currentState[message.Chat.ID] = StateIdle
}

// Handle the simple message
func handleSimpleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	if utils.IsCyrillic(message.Text) {
		reply := utils.CyrillicToLatin(message.Text)
		sendMessage(bot, message.Chat.ID, reply)
	}
	if utils.IsLatin(message.Text) {
		reply := utils.LatinToCyrillic(message.Text)
		sendMessage(bot, message.Chat.ID, reply)
	}
}

// Send a message using the bot
func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
