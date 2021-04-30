package machinery

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	appconfig "kubernetes-go-demo/config"
	"kubernetes-go-demo/tasks"
)

var server *machinery.Server

func InitServer(conf *appconfig.MachineryConfig) {
	cnf := &config.Config{
		DefaultQueue:    conf.DefaultQueue,
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
		"add":                   tasks.Add,
		"multiply":              tasks.Multiply,
		"sum_ints":              tasks.SumInts,
		"sum_floats":            tasks.SumFloats,
		"concat":                tasks.Concat,
		"split":                 tasks.Split,
		"panic_task":            tasks.PanicTask,
	}
	err := server.RegisterTasks(tasksList)
	if err != nil {
		panic(err)
	}
}

func GetServer() *machinery.Server {
	return server
}
