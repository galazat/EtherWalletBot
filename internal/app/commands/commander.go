package commands

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/galazat/EtherWalletBot/internal/service/blockchain"
)

type CommandData struct {
	Offset int `json:"offset"`
}

type Commander struct {
	currentStage string
	bot          *tgbotapi.BotAPI
	client       *ethclient.Client
	account      *blockchain.Account
}

func NewCommander(bot *tgbotapi.BotAPI, client *ethclient.Client) *Commander {
	return &Commander{
		currentStage: "",
		bot:          bot,
		client:       client,
		account:      blockchain.InitAccount(),
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
	case "auth":
		c.Auth(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	case "start":
		c.Start(update.Message)
	case "account":
		c.Auth(update.Message)
	default:
		c.Default(update.Message)
	}
}
