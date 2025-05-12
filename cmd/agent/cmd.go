package main

import (
	"github.com/Neftik/project3/internal/agent"
	"github.com/Neftik/project3/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	agent := agent.New(cfg)
	agent.Run()
}
