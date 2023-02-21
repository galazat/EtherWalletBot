package commands

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandData struct {
	Offset int `json:"offset"`
}

type Commander struct {
	bot    *tgbotapi.BotAPI
	client *ethclient.Client
}

func NewCommander(bot *tgbotapi.BotAPI, client *ethclient.Client) *Commander {
	return &Commander{
		bot:    bot,
		client: client,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recover from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		update.Message = update.CallbackQuery.Message
		switchCommand(update.CallbackQuery.Data, c, update)
		return
	}

	if update.Message == nil { // If we got a message
		return
	}

	switchCommand(update.Message.Command(), c, update)

}

func switchCommand(arg string, c *Commander, update tgbotapi.Update) {
	switch arg {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	case "start":
		c.Start(update.Message)
	default:
		c.Default(update.Message)
	}
}
