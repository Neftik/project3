package main

import (
	"github.com/Neftik/project3/internal/orchestrator"
)

func main() {
	app := orchestrator.New()

	app.Run()
}
