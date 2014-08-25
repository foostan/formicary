package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/formicary/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "ant",
		Usage:       "",
		Description: "",
		Flags:       command.AntFlags,
		Action:      command.AntCommand,
	},
	cli.Command{
		Name:        "graph",
		Usage:       "",
		Description: "",
		Flags:       command.GraphFlags,
		Action:      command.GraphCommand,
	},
}
