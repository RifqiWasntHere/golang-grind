package util

import "github.com/spf13/viper"

// Define the value of the environment credentials within this struct, and tag it with mapstructure:"<env_variable_name>"
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS" `
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)  // Adds the config.go path
	viper.SetConfigName("app") // Named "app" because of app.env
	viper.SetConfigType("env") // Available to JSON, and INI

	viper.AutomaticEnv() // Crosschecks .env content with the defined struct

	err = viper.ReadInConfig() // Discover and load the config file
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // Unmarshals said config into the defined struct
	return
}
