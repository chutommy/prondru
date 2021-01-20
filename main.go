package main

import (
	"log"
	"os"
	"prondru/controller"
)

func main() {
	app := controller.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
