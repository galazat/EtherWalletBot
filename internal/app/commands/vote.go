package commands

import (
	"fmt"
	"log"
	"math/big"
	"strconv"

	//"github.com/galazat/go-telegram-bot/internal/service/currency"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	vote "github.com/galazat/EtherWalletBot/internal/vote"
)

func (c *Commander) Vote(inputMessage *tgbotapi.Message) {
	outputMsgText := "Голос зачтён! \n\n"

	address := common.HexToAddress("0x3B93b909F99E453f91E32Ed448E0F7f3e600BE1E")
	instance, err := vote.NewVote(address, c.client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	arg := inputMessage.CommandArguments()
	fmt.Printf("TEST: %+v", arg)
	num, _ := strconv.Atoi(arg)
	fmt.Println("TEST2: ", arg)
	_, err = instance.Vote(nil, big.NewInt(int64(num)))
	if err != nil {
		log.Println(err)
	}
	fmt.Println("TEST3: ", arg)

	version2, err2 := instance.Candidats(nil, big.NewInt(int64(num)))
	if err != nil {
		log.Println(err2)
	}

	outputMsgText += fmt.Sprintf("Вы проголосовали за %s\n", version2.Name)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "account"),
		),
	)

	c.bot.Send(msg)
}
