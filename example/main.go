package main

import (
	"fmt"
	prctl "github.com/wl102/go-prctl"
	"log"
	"os"
	"os/signal"
)

func main() {
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)

	prName := "newname"

	err := prctl.SetProcessName(prName)
	if err != nil {
		log.Println(err)
	}

	name, err := prctl.GetProcessName()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(name)
	}

	<-wait
	fmt.Println("exit.")
}
