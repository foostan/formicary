package agent

import (
)

type Agent struct {
	config *Config
}

func Create(config *Config) (*Agent, error) {

	return new(Agent), nil
}
