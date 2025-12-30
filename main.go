package main

import (
	"github.com/racibaz/go-arch/cmd"
)

func main() {
	run()
}

func run() {
	// If you want to use  cobra cli
	// Docker entry point run this
	cmd.Execute()

	// If you want to use local debugging without cobra cli, uncomment it.
	//bootstrap.Serve()
}
