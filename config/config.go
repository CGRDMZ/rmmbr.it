package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var (
	Conf Config
)

type OAuth map[string]*OAuthConfig

type OAuthConfig struct {
	ClientId     string
	ClientSecret string
}

type Config struct {
	DbConnectionString string
	Env                string
	Port               string
	JwtSecret          string
	JwtExpiresIn       int
	OAuth              OAuth
	IsHeroku           bool
}

func init() {
	Conf.OAuth = make(OAuth)
	
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

	Conf.OAuth["google"] = &OAuthConfig{
		ClientId: viper.GetString("OAuth.Google.ClientId"),
		ClientSecret: viper.GetString("Oauth.Google.ClientSecret"),
	} 
}

func LoadConfig() {

	viper.AddConfigPath("./config/")
	viper.SetConfigType("json")

	viper.SetConfigName("config")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	viper.BindEnv("Server.Port", "PORT")
	viper.BindEnv("Database.ConnectionString", "DATABASE_URL")

	err := viper.ReadInConfig()
	if err != nil {
		if notFoundErr, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("config file not found. %w", notFoundErr))
		}
		panic(err)
	}

	bindConfig()

}
