package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	// read server configuration
	fmt.Println("Server port:: ", viper.GetInt("server.port"))
	fmt.Println("Server port:: ", viper.GetString("security.jwt.key"))

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}

	fmt.Println("Config Port:", config.Server.Port)

	for index, db := range config.Databases {
		fmt.Printf("%v: Database User: %s, password: %s, Host: %s \n", index+1, db.User, db.Password, db.Host)
	}
}
