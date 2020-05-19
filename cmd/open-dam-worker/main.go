package main

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/log"
	"github.com/open-dam/open-dam-worker/pkg/opendam"
	"github.com/open-dam/open-dam-worker/pkg/tasks"
)

func main() {
	logger := opendam.Logger()
	log.Set(logger)

	cnf, err := config.NewFromEnvironment(true)
	if err != nil {
		logger.WithError(err).Fatal("failed to build machinery config from environment")
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		logger.WithError(err).Fatal("failed to start machinery server")
	}
	_ = server.RegisterTasks(map[string]interface{}{
		"add":      tasks.Add,
		"multiply": tasks.Multiply,
	})

	worker := server.NewWorker("worker_name", 10)
	err = worker.Launch()
	if err != nil {
		logger.WithError(err).Fatal("failed to launch machinery worker")
	}

}
