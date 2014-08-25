package ant

import (
	"fmt"
)

type Ant struct {
	config *Config
}

func Create(config *Config) (*Ant, error) {
	return &Ant{
		config: config,
	}, nil
}

func (a *Ant) Run() error {
	fmt.Println("Running ant")
	fmt.Println(a.config.Node)
	fmt.Println(a.config.NodeGroup)
	fmt.Println(a.config.Connection)

	return nil
}
