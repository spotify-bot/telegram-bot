package main

import (
	"github.com/spotify-bot/telegram/internal/config"
	"github.com/spotify-bot/telegram/internal/telegram"
)

func main() {
	tgbot := telegram.New(telegram.TGBotOptions{
		Token: config.AppConfig.TelegramBot.APIToken,
	})
	tgbot.Start()
}
