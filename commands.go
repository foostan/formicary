package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAgent,
	commandGraph,
}

var commandAgent = cli.Command{
	Name:  "agent",
	Usage: "",
	Description: `
`,
	Action: doAgent,
}

var commandGraph = cli.Command{
	Name:  "graph",
	Usage: "",
	Description: `
`,
	Action: doGraph,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doAgent(c *cli.Context) {
}

func doGraph(c *cli.Context) {
}
