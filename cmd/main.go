package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("env", "local")

	fmt.Printf("Hello World\n")
}
