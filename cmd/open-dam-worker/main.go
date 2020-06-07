package main

import (
	"time"

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

	logger.Infof("configgy: %+v\n", cnf)

	server, err := machinery.NewServer(cnf)
	if err != nil {
		logger.WithError(err).Fatal("failed to start machinery server")
	}
	tasker, err := tasks.NewTasker(server, logger)
	if err != nil {
		logger.WithError(err).Fatal("failed to start tasker")
	}

	_ = server.RegisterTasks(map[string]interface{}{
		"upload":        tasker.Upload,
		"extract":       tasker.Extract,
		"soundwave":     tasker.Soundwave,
		"imageanalysis": tasker.ImageAnalysis,
		"imagecreation": tasker.ImageCreation,
	})

	worker := server.NewWorker("worker_name", 10)
	go func() {
		err = worker.Launch()
		if err != nil {
			logger.WithError(err).Fatal("failed to launch machinery worker")
		}
	}()

	logger.Info("started workers yall")
	time.Sleep(30 * time.Second)
	c := tasks.ProcessAsset("0000")
	_, err = server.SendChain(c)
	if err != nil {
		logger.WithError(err).Error("failed to chain")
	}
	// v, err := r.GetWithTimeout(time.Minute, time.Minute)
	// if err != nil {
	// 	logger.WithError(err).Error("failed to get")
	// }
	// for i, vvv := range v {
	// 	fmt.Printf("%d: %+v", i, vvv)
	// }
	select {}

}
