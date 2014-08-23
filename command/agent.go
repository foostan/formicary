package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

var AgentFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "test",
		Usage: "",
	},
}

func AgentCommand(c *cli.Context) {
	fmt.Println("Run agent command with test: ", c.String("test"))
}
