package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/formicary/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "agent",
		Usage:       "",
		Description: "",
		Flags:       command.AgentFlags,
		Action:      command.AgentCommand,
	},
	cli.Command{
		Name:        "graph",
		Usage:       "",
		Description: "",
		Flags:       command.GraphFlags,
		Action:      command.GraphCommand,
	},
}
