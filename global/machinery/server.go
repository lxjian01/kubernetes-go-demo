package machinery

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/tasks"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	appconfig "kubernetes-go-demo/config"
	mytask "kubernetes-go-demo/tasks"
)

var server *machinery.Server

func InitServer(conf *appconfig.MachineryConfig) {
	cnf := &config.Config{
		Broker: conf.Broker,
		DefaultQueue:    conf.DefaultQueue,
		ResultBackend: conf.Backend,
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	// Create server instance
	broker := redisbroker.New(cnf, conf.Broker, "", "", 0)
	backend := redisbackend.New(cnf, conf.Backend, "", "", 0)
	lock := eagerlock.New()
	server = machinery.NewServer(cnf, broker, backend, lock)
}

func RegistryTasks(){
	// Register tasks
	tasksList := map[string]interface{}{
		"add":                   mytask.Add,
		"multiply":              mytask.Multiply,
		"sum_ints":              mytask.SumInts,
		"sum_floats":            mytask.SumFloats,
		"concat":                mytask.Concat,
		"split":                 mytask.Split,
		"panic_task":            mytask.PanicTask,
	}
	err := server.RegisterTasks(tasksList)
	if err != nil {
		panic(err)
	}
}

func RegisterScheduledTask(){
	// Register tasks
	signature := &tasks.Signature{
		Name: "Multiply",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: 2,
			},
			{
				Type:  "int64",
				Value: 3,
			},
		},
	}
	err := server.RegisterPeriodicTask("0/2 * * * * ?", "periodic-task", signature)
	if err != nil {
		panic(err)
	}
}


func StartWorker(){
	workers := server.NewWorker("machinery_tasks", 10)
	err := workers.Launch()
	if err != nil {
		panic(err)
	}
}

func GetServer() *machinery.Server {
	return server
}
