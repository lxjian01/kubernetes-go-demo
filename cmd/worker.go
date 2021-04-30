package cmd

import (
	"github.com/spf13/cobra"
	globalConf "kubernetes-go-demo/global/config"
	"kubernetes-go-demo/global/log"
	globalMachinery "kubernetes-go-demo/global/machinery"
	"os"
	"fmt"
)

func init() {
	rootCmd.AddCommand(machineryWorkerCmd)
}

var machineryWorkerCmd = &cobra.Command{
	Use:   "worker",
	Run: func(cmd *cobra.Command, args []string) {
		conf := globalConf.GetAppConfig()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Start machinery worker error by ", r)
				os.Exit(1)
			}
		}()
		log.Info("Starting init machinery server")
		globalMachinery.InitServer(conf.Machinery)
		log.Info("Init machinery server ok")

		log.Info("Starting registry machinery tasks")
		globalMachinery.RegistryTasks()
		log.Info("Registry machinery tasks ok")

		log.Info("Starting machinery worker")
		globalMachinery.StartWorker()
		log.Info("Machinery worker start ok")

		globalMachinery.RegisterScheduledTask()

	},
}