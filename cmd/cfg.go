package cmd

import "github.com/spf13/viper"

var (
	cfgFile string
)

func parseConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile) // Use config file from the flag.
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(".")      // optionally look for config in the working directory
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignoring error
		} else {
			// incorrect config, etc
			panic(err)
		}
	}

	Token = viper.GetString("token")
	if viper.GetString("project-id") != "" {
		projectId = viper.GetString("project-id")
	}
}
