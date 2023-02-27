package commands

import (
	"context"
	"fmt"
	"log"
	"math/big"

	//"github.com/galazat/go-telegram-bot/internal/service/currency"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Account(inputMessage *tgbotapi.Message) {

	outputMsgText := "Данные кошелька: \n\n"
	outputMsgText = fmt.Sprintf("Адрес кошелька: %s \nБаланс: %s wai", c.account.GetAddress(), GetBalance(c))
	//log.Println("TEST DATA ---  USER KEY: ", c.account, "\nADDR: ", c.account.GetAddress(), "\nPublic: ", c.account.GetPublicKey(), "\nPrivate: ", c.account.GetPrivateKey())

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Список кандидатов", "list"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Проголосовать", "Vote"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Показать лидера", "get"),
		),
	)

	c.bot.Send(msg)
}

func GetBalance(c *Commander) *big.Int {
	balance, err := c.client.BalanceAt(context.Background(), c.account.Address, nil)
	if err != nil {
		log.Println(err)
		return balance
	}

	return balance
}
