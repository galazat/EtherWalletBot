package commands

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Auth(inputMessage *tgbotapi.Message) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recover from panic: %v", panicValue)
		}
	}()

	if c.currentStage != "auth" {
		c.currentStage = "auth"
		outputMsgText := "Введите private key кошелька: \n\n"

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
		c.bot.Send(msg)
	} else {
		outputMsgText := ""
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

		inputPrivateKey := inputMessage.Text
		privateKey, err := crypto.HexToECDSA(inputPrivateKey)
		log.Println("TEST DATA ---  LEN: ", privateKey, "ERR: ", err)

		if err != nil {
			c.currentStage = "reject"
			outputMsgText = "Некорректный private key кошелька"
			msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
			msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Авторизация", "auth"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Назад", "start"),
				),
			)
			log.Println(err)
			c.bot.Send(msg)
		} else {
			c.account.SetPrivateKey(privateKey)
			c.account.PublicFromPrivate()
			c.account.AddressFromPublic()
			log.Println("TEST DATA ---  USER KEY: ", c.account, "\nADDR: ", c.account.GetAddress(), "\nPublic: ", c.account.GetPublicKey(), "\nPrivate: ", c.account.GetPrivateKey())

			outputMsgText = "Авторизация пройдена"
			msg = tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)
			c.bot.Send(msg)

			c.Account(inputMessage)
		}

	}
}
