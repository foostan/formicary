package command

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/foostan/formicary/command/agent"
)

var AgentFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "test",
		Usage: "",
	},
}

func AgentCommand(c *cli.Context) {
	fmt.Println("Run agent command with test: ", c.String("test"))

	paths := []string{"config"}
	config, err := agent.ReadConfigPaths(paths)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", paths, err)
	}
	config = MergeConfig(DefaultConfig(), config)
	fmt.Println(config.Node)
	fmt.Println(config.NodeGroup)
	fmt.Println(config.Connection)

	agent, err := agent.Create(config)
	fmt.Println(agent)
}
