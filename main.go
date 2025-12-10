package main

import (
	"github.com/racibaz/go-arch/cmd"
)

func main() {
	run()
}

func run() {
	cmd.Execute() // if you want  use  cobra cli
	//bootstrap.Serve() //if you want changed it for local debugging
}
