package main

import (
	"log"

	"github.com/aimamit/hostand-cli/api"
	"github.com/aimamit/hostand-cli/cmd"
	"github.com/aimamit/hostand-cli/config"
)

func main() {
	config.Init()
	err := api.Init()
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute()
}
