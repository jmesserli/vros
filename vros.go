package main

import (
	"fmt"

	"github.com/jmesserli/vros/config"
)

func main() {
	config := config.ReadConfig("./config.yaml")
	fmt.Printf("%+v", config)
}
