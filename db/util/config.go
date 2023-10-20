package util

import (
	"time"

	"github.com/spf13/viper"
)

// It stores all the configuration vars of the app.
// The values are read by viper from a config file or env vars.
// We use mapstructure tags to specify the name of each config field (that is, env vars in this case).
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

// To get the values and store in struct we use unmarshalling freature of viper.
// viper uses the mapstructure pkg under the hood, for unmarshalling values - https://github.com/mitchellh/mapstructure

// reads configuration from env vars
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Doubt
	viper.AutomaticEnv()

	// Doubt
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
