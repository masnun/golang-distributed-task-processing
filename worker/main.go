package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/errors"
)

func main() {

	var cnf = config.Config{
		Broker:        "redis://127.0.0.1:6379",
		ResultBackend: "redis://127.0.0.1:6379",
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		errors.Fail(err, "Could not create server")
	}

	server.RegisterTask("Say", Say)

	worker := server.NewWorker("worker-1")
	err = worker.Launch()
	if err != nil {
		errors.Fail(err, "Could not launch worker!")
	}

}
