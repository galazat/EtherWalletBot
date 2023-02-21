package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/galazat/EtherWalletBot/internal/app/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	//token := os.Getenv("TOKEN")

	// Del in prod
	token1 := ""

	bot, err := tgbotapi.NewBotAPI(token1)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	var client *ethclient.Client

	commader := commands.NewCommander(bot, client)

	for update := range updates {
		commader.HandleUpdate(update)
	}
}

//"github.com/galazat/EtherWalletBot/internal/app/commands"
