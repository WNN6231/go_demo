package main

import (
	"exchangeapp/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Println(config.AppConfig.App.Port)
}
