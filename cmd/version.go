package cmd

import (
	"fmt"
	globalConf "kubernetes-go-demo/global/config"
	"github.com/spf13/cobra"
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