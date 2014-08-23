package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var GraphFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "test",
		Usage: "",
	},
}

func GraphCommand(c *cli.Context) {
	fmt.Println("Run graph command with test: ", c.String("test"))
}
