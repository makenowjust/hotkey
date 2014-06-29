// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

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
