package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"prondru/controller"

	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload"
	"github.com/mattn/go-colorable"
)

func main() {
	// print banner
	banner.Init(colorable.NewColorableStdout(), true, false, bytes.NewBufferString(""))
	fmt.Println()

	// run CLI
	app := controller.NewApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
