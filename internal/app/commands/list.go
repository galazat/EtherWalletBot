package commands

import (
	"fmt"
	"log"

	//"github.com/galazat/go-telegram-bot/internal/service/currency"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	vote "github.com/galazat/EtherWalletBot/internal/vote"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Cписок кандидатов: \n\n"

	address := common.HexToAddress("0x3B93b909F99E453f91E32Ed448E0F7f3e600BE1E")
	instance, err := vote.NewVote(address, c.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	totalCand, err := instance.TotalCand(nil)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < totalCand; i++ {
		version, err := instance.Candidats(nil, i)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("Имя кандидата:\n%s\nКоличество голосов:\n%s\n\n", version.Name, version.VoteCount)
		outputMsgText += fmt.Sprintf("Имя кандидата:\n%s\nКоличество голосов:\n%s\n\n", version.Name, version.VoteCount)

	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "account"),
		),
	)

	c.bot.Send(msg)
}
