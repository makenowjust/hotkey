package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func main() {
	quit := make(chan bool)

	hotkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit")
		quit <- true
	})

	hotkey.Start()

	fmt.Println("Start hotkey's loop")
	fmt.Println("Push Ctrl-Q to escape and quit")
	<-quit
}
