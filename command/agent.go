package command

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/foostan/formicary/command/agent"
)

var AgentFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config-dir",
		Usage: "directory of configuration files to load",
	},
}

func AgentCommand(c *cli.Context) {
	paths := []string{c.String("config-dir")}
	config, err := agent.ReadConfigPaths(paths)
	if err != nil {
		log.Fatalf("Error reading '%s': %s", paths, err)
	}
	config = agent.MergeConfig(agent.DefaultConfig(), config)

	fmt.Println(config.Node)
	fmt.Println(config.NodeGroup)
	fmt.Println(config.Connection)

	agent, err := agent.Create(config)
	fmt.Println(agent)
}
