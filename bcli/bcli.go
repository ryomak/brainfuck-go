package bcli

import (
	"github.com/urfave/cli"
)

/*
MIT License

Copyright (c) 2016 Jeremy Saenz & Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.<Paste>
*/

func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = getCommands()
	return app
}

func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "run",
			Usage:  "run bf file",
			Action: run,
			Flags:  getFlags(),
		},
	}
}

func getFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "c, config",
			Value: "",
			Usage: "config for setting bf token",
		},
		cli.IntFlag{
			Name:  "m,max",
			Value: 1000,
			Usage: "set max order num",
		},
    cli.BoolFlag{
			Name:  "d,default",
			Usage: "set default brainfxxk token",
		},
	}
}
