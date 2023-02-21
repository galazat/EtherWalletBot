package commands

import (
	"encoding/hex"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	ms := inputMessage.Text
	log.Printf("INFO: [%s] %s, %v", inputMessage.From.UserName, inputMessage.Text, hex.EncodeToString([]byte(ms)))

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Привет, лучше посмотри на список доступных команд /help")

	c.bot.Send(msg)
}
