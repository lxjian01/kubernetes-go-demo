package machinery

import (
	"github.com/RichardKnop/machinery/v2"
	backendsiface "github.com/RichardKnop/machinery/v2/backends/iface"
	brokersiface "github.com/RichardKnop/machinery/v2/brokers/iface"
	machineryconfig "github.com/RichardKnop/machinery/v2/config"
	locksiface "github.com/RichardKnop/machinery/v2/locks/iface"
	"github.com/RichardKnop/machinery/v2/tasks"
	appConfig "kubernetes-go-demo/config"
)

func test(conf appConfig.MachineryConfig){
	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 1,
			},
			{
				Type:  "int64",
				Value: 1,
			},
		},
	}
	var broker brokersiface.Broker
	var backend backendsiface.Backend
	var lock locksiface.Lock
	cnf := machineryconfig.Config{
		Broker: conf.Broker,
		DefaultQueue: conf.DefaultQueue,
		ResultBackend: conf.ResultBackend,
	}
	server := machinery.NewServer(&cnf, broker, backend, lock)
	err := server.RegisterPeriodicTask("0 6 * * ?", "periodic-task", signature)
	if err != nil {
		// failed to register periodic task
	}
}
