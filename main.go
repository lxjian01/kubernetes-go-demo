package main

import (
	"fmt"
	"kubernetes-go-demo/cmd"
	"os"
)

func main() {


	// start httpd server
	err1 := cmd.RootCmdExecute()
	if err1 != nil {
		fmt.Println("Start httpd server error by ",err1)
		os.Exit(1)
	}
}





