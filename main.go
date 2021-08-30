package main

import (
	"github.com/chutommy/prondru/controller"
	"log"
	"os"
)

func main() {
	// run CLI
	app := controller.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
