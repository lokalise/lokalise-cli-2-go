package cmd

import "github.com/spf13/viper"

func init() {
	parseConfig()
}

func parseConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignoring error
		} else {
			// incorrect config, etc
			panic(err)
		}
	}
}
