package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("env", "local v2")

	fmt.Printf("Hello World\n")
}
