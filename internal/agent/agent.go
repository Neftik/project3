package agent

import (
	"log"

	"github.com/Neftik/project3/internal/config"
	"github.com/Neftik/project3/internal/models"
)

type Agent struct {
	config config.Config
}

type Task struct {
	ID   int
	Arg1 string
	Arg2 string
	Type string
}

var (
	resultsCh = make(chan *models.Result)
	tasksCh   = make(chan *Task)
)

func New(cfg config.Config) *Agent {
	return &Agent{config: cfg}
}

func (a *Agent) Run() {
	go a.Connect()

	for i := range a.config.ComputingPower {
		log.Printf("worker %d starting...", i+1)
		go worker(a.config)
	}

	select {}
}
