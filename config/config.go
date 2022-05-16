package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Conf Config
)

type Config struct {
	DbConnectionString string
	Env                string
	Port               string
	JwtSecret          string
	JwtExpiresIn       int
}

func init() {
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RMMBR")
}

func bindConfig() {
	Conf.DbConnectionString = viper.GetString("Database.ConnectionString")
	Conf.Env = viper.GetString("Environment")
	Conf.Port = viper.GetString("Server.Port")
	Conf.JwtSecret = viper.GetString("Jwt.Secret")
	Conf.JwtExpiresIn = viper.GetInt("Jwt.Expiration")
}

func LoadConfig(name string) {

	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")

	viper.SetConfigName(name)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if notFoundErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found. %w", notFoundErr))
		}
		panic(err)
	}

	bindConfig()

}
