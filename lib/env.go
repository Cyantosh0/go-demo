package lib

import (
	"fmt"

	"github.com/spf13/viper"
)

type Env struct {
	SeverPort string `mapstructure:"SERVER_PORT"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

var env Env

func GetEnv() *Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read cofiguration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		fmt.Errorf("unable to decode into struct, %v", err)
	}
	return &env
}
