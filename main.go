package main

import "github.com/racibaz/go-arch/pkg/bootstrap"

func main() {
	run()
}

func run() {
	// If you want to use  cobra cli
	// Docker entry point run this command
	// cmd.Execute()

	// If you want to use local debugging without cobra cli
	bootstrap.Serve()
}
