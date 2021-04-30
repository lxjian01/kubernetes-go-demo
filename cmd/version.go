package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	globalConf "kubernetes-go-demo/global/config"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Run: func(cmd *cobra.Command, args []string) {
		conf := globalConf.GetAppConfig()
		fmt.Println("kubernetes go demo version ", conf.Version)
	},
}