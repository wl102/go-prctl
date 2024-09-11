package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	prctl "github.com/wl102/go-prctl"
)

func main() {
	prName := "newname"
	// make sure first running
	err := prctl.SetProcessName(prName)
	if err != nil {
		log.Println(err)
	}

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)

	name, err := prctl.GetProcessName()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(name)
	}

	<-wait
	fmt.Println("exit.")
}
