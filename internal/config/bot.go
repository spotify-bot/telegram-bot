package config

type Config struct {
	Webserver   Webserver   `mapstructure:"webserver"`
	TelegramBot TelegramBot `mapstructure:"telegramBot"`
}

type Webserver struct {
	Address string `mapstructure:"address"`
}

type TelegramBot struct {
	APIToken string `mapstructure:"apiToken"`
}

var AppConfig Config

func initMamalConfig() {
	loadConfig(&AppConfig)
}
