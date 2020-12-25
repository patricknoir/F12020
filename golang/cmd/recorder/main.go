package main

import (
	"fmt"
	"github.com/patricknoir/F12020/pkg/common/config"
	"log"
	"os"
)

func main() {
	var cfg config.Config
	err := config.FromEnv(&cfg)
	exitOnError(err)
	fmt.Printf("Config = %+v\n", cfg)
}

func exitOnError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
