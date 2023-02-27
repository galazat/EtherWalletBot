package commands

import (
	"fmt"
	"log"

	//"github.com/galazat/go-telegram-bot/internal/service/currency"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	vote "github.com/galazat/EtherWalletBot/internal/vote"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	outputMsgText := "Лидер голосования: \n\n"

	address := common.HexToAddress("0x3B93b909F99E453f91E32Ed448E0F7f3e600BE1E")
	instance, err := vote.NewVote(address, c.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	version, err := instance.WinnerName(nil)
	if err != nil {
		log.Println(err)
	}

	outputMsgText += fmt.Sprintf("Имя кандидата: %s", version)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "account"),
		),
	)

	c.bot.Send(msg)
}
