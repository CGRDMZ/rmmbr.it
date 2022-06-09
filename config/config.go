package config

import (
	"fmt"
	"strings"
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
	IsHeroku		   bool
}

func init() {
	viper.SetConfigType("json")
	viper.SetEnvPrefix("RMMBR")
}

func bindConfig() {
	Conf.DbConnectionString = viper.GetString("Database.ConnectionString")
	Conf.Env = viper.GetString("Environment")
	Conf.Port = viper.GetString("Server.Port")
	Conf.JwtSecret = viper.GetString("Jwt.Secret")
	Conf.JwtExpiresIn = viper.GetInt("Jwt.Expiration")
	Conf.IsHeroku = viper.GetBool("IsHeroku")
}

func LoadConfig() {

	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")

	viper.SetConfigName("config")
	
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	viper.BindEnv("Server.Port", "PORT")
	
	err := viper.ReadInConfig()
	if err != nil {
		if notFoundErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found. %w", notFoundErr))
		}
		panic(err)
	}
	
	bindConfig()

}
