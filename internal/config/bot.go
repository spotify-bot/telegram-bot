package config

type Config struct {
	APIServerAddress string `mapstructure:"API_SERVER_ADDRESS"`
	TelegramAPIToken string `mapstructure:"TELEGRAM_API_TOKEN"`
}

var AppConfig Config

func initMamalConfig() {
	loadConfig(&AppConfig)
}
