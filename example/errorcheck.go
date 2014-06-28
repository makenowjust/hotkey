package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func main() {
	hotkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit 1")
	})
	hotkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit 2")
	})

	chErr := hotkey.Start()

	if err := <-chErr; err != nil {
		fmt.Println(err)
	}
}
