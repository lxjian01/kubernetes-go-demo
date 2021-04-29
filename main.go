package main

import (
	"fmt"
	"kubernetes-go-demo/cmd"
	"os"
)

func main() {
	// start httpd server
	err := cmd.HttpdCmdExecute()
	if err != nil {
		fmt.Println("Start httpd server error by ",err)
		os.Exit(1)
	}
}





