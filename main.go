package main

import (
	"fmt"
	"kubernetes-go-demo/cmd"
	"os"
)

func main() {
	// start server
	err1 := cmd.Execute()
	if err1 != nil {
		fmt.Println("Start server error by ",err1)
		os.Exit(1)
	}
}





