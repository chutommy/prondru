package main

import (
	"log"
	"os"
	"prondru/controller"
)

func main() {
	app := controller.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
