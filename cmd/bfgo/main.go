package main

import (
	"os"

	"github.com/ryomak/brainfuck-go/bcli"
)

func main() {
	app := bcli.New("brainfuck-go", "", "1.0.0")
	app.Run(os.Args)
}
