package initialize

import (
	"fmt"

	"example.com/m/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}
