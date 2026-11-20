package main

import (
	"fmt"
	"os"
	"tig/app/bootstrap"
)

func main() {
	app, err := bootstrap.InitializeApp()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Initialization error: %v\n", err)
		os.Exit(1)
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
