package main

import (
	"fmt"

	"github.com/jmesserli/vros/config"
)

func main() {
	config := config.ReadConfig("./config.json")
	fmt.Printf("%+v", config)
}
