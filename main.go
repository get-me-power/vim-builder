package main

import (
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	(&cli.App{
		Name: "vim-builder",
	}).Run(os.Args)
}
