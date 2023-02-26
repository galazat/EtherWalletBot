package main

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"github.com/galazat/EtherWalletBot/internal/app/commands"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
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

	// ctx := context.Background()
	// deploy.Deploy(client, ctx)

	for update := range updates {
		commader.HandleUpdate(update)
	}
}

//"github.com/galazat/EtherWalletBot/internal/app/commands"
