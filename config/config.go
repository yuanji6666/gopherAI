package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct{
	MainConfig
	MysqlConfig
	RedisConfig
	EmailConfig
	JwtConfig
}

type MysqlConfig struct{
	MysqlHost		string
	MysqlPort		int
	MysqlUser		string
	MysqlPassword	string
	MysqlDBName		string
	MysqlCharset	string
}
type RedisConfig struct{
	RedisHost		string
	RedisPort		int
	RedisDB			int
	RedisPassword	string
}

type EmailConfig struct{
	Authcode 	string
	Email		string
}

type MainConfig struct{
	Host 			string
	Port 			int
	AppName 		string
}


type RedisKeyConfig struct {
	CaptchaPrefix string
}

var DefaultRedisKeyConfig = RedisKeyConfig{
	CaptchaPrefix: "captcha:%s",
}

type JwtConfig struct{
	ExpireDuration		int
	Issuer				string
	Subject				string
	Key 				string
}

var config *Config

func InitConfig(){
	config = new(Config)
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig() ; err != nil {
		log.Fatalf("Read in config error : %v", err.Error())
	}
	if err := viper.Unmarshal(config) ; err != nil {
		log.Fatalf("Unmarshal config error : %v", err.Error())
	}
}

func GetConfig() *Config{
	if config == nil {
		config = new(Config)
		InitConfig()
	}
	return config
}

