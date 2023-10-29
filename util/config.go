package util

import (
	"os"
)

type Config struct {
	SlackSigningSecret string `mapstructure:"SLACK_SIGNING_SECRET"`
	SlackBotToken      string `mapstructure:"SLACK_BOT_TOKEN"`
	// SlackSocketToken   string `mapstructure:"SLACK_SOCKET_TOKEN"`
}

func LoadConfig() (Config, error) {
	// return loadConfigByViper("./", "env")
	config := Config{}
	config.SlackSigningSecret = os.Getenv("SLACK_SIGNING_SECRET")
	config.SlackBotToken = os.Getenv("SLACK_BOT_TOKEN")
	return config, nil
}

// func loadConfigByViper(path, appName string) (config Config, err error) {
// 	viper.AddConfigPath(path)
// 	viper.SetConfigName(appName)
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	err = viper.ReadInConfig()
// 	if err != nil {
// 		return
// 	}
// 	err = viper.Unmarshal(&config)
// 	return
// }
