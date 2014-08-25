package command

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/foostan/formicary/command/ant"
)

var AntFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "config-dir",
		Usage: "directory of configuration files to load",
	},
}

func AntCommand(c *cli.Context) {
	config_dir := c.String("config-dir")
	if config_dir == "" {
		log.Fatalf("Error missing flag 'config-dir'")
	}

	cnf, err := ant.ReadConfigPaths([]string{config_dir})
	if err != nil {
		log.Fatalf("Error reading '%s': %s", config_dir, err)
	}

	agt, err := ant.Create(ant.MergeConfig(ant.DefaultConfig(), cnf))
	if err != nil {
		log.Fatalf("Error creating ant: %s", err)
	}

	err = agt.Run()
	if err != nil {
		log.Fatalf("Error running ant: %s", err)
	}
}
