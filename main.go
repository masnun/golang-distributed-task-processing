package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/errors"
	"github.com/RichardKnop/machinery/v1/signatures"
)

func main() {

	var cnf = config.Config{
		Broker:        "redis://127.0.0.1:6379",
		ResultBackend: "redis://127.0.0.1:6379",
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		errors.Fail(err, "Can not create server!")
	}

	sayTask := signatures.TaskSignature{
		Name: "Say",
		Args: []signatures.TaskArg{
			signatures.TaskArg{
				Type:  "string",
				Value: "masnun",
			},
		},
	}

	server.SendTask(&sayTask)

}
