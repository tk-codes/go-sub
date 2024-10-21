package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tk-codes/go-sub/pkg/math/v2"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("env", "local v2")

	fmt.Printf("Hello World\n")

	fmt.Printf("Result 1 + 3 is %d", math.Add(1, 3))
}
