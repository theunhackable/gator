package main

import (
	"fmt"
	"os"

	"github.com/theunhackable/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config file. %w", err)
		os.Exit(-1)
	}
	cfg.SetUser()

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Println("Error while reading updated config: %w", err)
		os.Exit(-1)
	}

	fmt.Println(updatedCfg.DBUrl)

}
