package config

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	APIServerAddress string `mapstructure:"API_SERVER_ADDRESS"`
	TelegramAPIToken string `mapstructure:"TELEGRAM_API_TOKEN"`
}

var (
	ConfigName string
	AppConfig  Config
)

func init() {
	// Register flags
	flag.StringVar(&ConfigName, "config-name", "", "Path to config directory")
	flag.Parse()

	loadConfig(&AppConfig)
}

func loadConfig(configStruct interface{}) {
	viperInstance := viper.New()

	// If Configname is passed, read the config from file
	if ConfigName != "" {
		viperInstance.SetConfigName(ConfigName)
		viperInstance.AddConfigPath(".")
		viperInstance.SetConfigType("env")

		err := viperInstance.ReadInConfig()
		if err != nil {
			log.Panic("Config file is not set", err)
		}

		err = viperInstance.Unmarshal(configStruct)
		if err != nil {
			log.Panic("Failed to Unmarshal Config file: ", err)
		}
		log.Printf("loaded config: %v\n", ConfigName)
	} else {
		viperInstance.AutomaticEnv()
		serverAddr := viperInstance.GetString("API_SERVER_ADDRESS")
		apiToken := viperInstance.GetString("TELEGRAM_API_TOKEN")
		if len(serverAddr) > 0 && len(apiToken) > 0 {
			AppConfig = Config{
				APIServerAddress: serverAddr,
				TelegramAPIToken: apiToken,
			}
		} else {
			log.Panic("ENV Vars should be set")
		}
	}
}
