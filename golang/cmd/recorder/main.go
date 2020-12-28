package main

import (
	"fmt"
	"github.com/patricknoir/F12020/cmd/recorder/recorder"
	"github.com/patricknoir/F12020/pkg/common/config"
	"log"
	"os"
)

func main() {
	var cfg config.Config
	err := config.FromEnv(&cfg)
	exitOnError(err)
	fmt.Println("Running recorder")
	rec, err := recorder.New(cfg.Server.Host, cfg.Server.Port, "./data/")
	exitOnError(err)
	recorder.Record(rec)
}

func exitOnError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
