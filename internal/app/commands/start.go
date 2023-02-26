package commands

import (
	"fmt"
	"log"

	//"github.com/galazat/go-telegram-bot/internal/service/currency"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/galazat/EtherWalletBot/internal/service/blockchain"
)

var START_CURRENCIES = []string{"USD", "EUR", "GBP", "UAH"}
var CHAT_ID int64

func (c *Commander) Start(inputMessage *tgbotapi.Message) {

	text := fmt.Sprintf("Привет, @%s ! \xF0\x9F\x91\x8B \n\n", inputMessage.Chat.UserName)
	text += fmt.Sprintf("Данный бот позволяет взаимодействовать с аккаунтом в приватной blockchain Etherium сети \xf0\x9f\x8c\x90. ")

	CHAT_ID = inputMessage.Chat.ID

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		text,
	)

	c.bot.Send(msg)

	var err error
	c.client, err = blockchain.ConnectionInit()
	if err != nil {
		text = "Ошибка. Не удалось подключиться к сети " + blockchain.BLOCHCHAIN_NET_ADRR
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)
		log.Println(err)
	} else {
		text = "Подключено к сети : " + blockchain.BLOCHCHAIN_NET_ADRR
		msg = tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Авторизация", "auth"),
			),
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonData("Полный список курсов валют", "list"),
			// ),
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonData("Конвертировать валюту", "convert"),
			// ),
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonData("Курс определённой валюты", "get"),
			// ),
		)
	}

	c.bot.Send(msg)
}
