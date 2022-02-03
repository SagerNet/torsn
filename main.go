package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ooni/go-libtor"
)

func main() {
	fmt.Println("start tor instance...")

	ctx := context.Background()
	torProcess, _ := libtor.Creator.New(ctx, os.Args[1:]...)

	err := torProcess.Start()
	if err != nil {
		fmt.Println("failed to start tor: ", err)
		os.Exit(1)
	}

	err = torProcess.Wait()
	if err != nil {
		fmt.Println("tor process exited: ", err)
		os.Exit(1)
	}
}
