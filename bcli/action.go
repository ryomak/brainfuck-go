package bcli

import (
	"fmt"
	"io/ioutil"
	"os"

	bf "github.com/ryomak/brainfuck-go"
	"github.com/urfave/cli"
)

func run(c *cli.Context) {
	config, err := bf.LoadConfig(c.String("config"))
	if err != nil {
		if c.String("config") != "" {
			fmt.Printf("cannot load config : %s \n", c.String("config"))
			os.Exit(0)
		}
    if c.Bool("d"){
      config = bf.DefaultConfig()
    }else{
		  config = bf.GokuConfig()
    }
	}
	if len(c.Args()) == 0 {
		fmt.Printf("no files listed")
		os.Exit(0)
	}
	src := []byte{}
	for _, v := range c.Args() {
		inputSrc, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Printf("cannot read %s", v)
			os.Exit(0)
		}
		src = append(src, inputSrc...)
	}
	state := bf.InitState(string(src), config.GetCommands(), c.Int("max"))
	state.Start()
}
