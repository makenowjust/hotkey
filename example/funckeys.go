// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func main() {
	for i := uint32(0); i < 12; i++ {
		hotkey.Register(hotkey.Alt+hotkey.Shift, hotkey.F1+i, func(i uint32) func() {
			return func() {
				fmt.Printf("Push Alt-Shift-F%d\n", i)
			}
		}(i+1))
	}

	quit := make(chan bool)
	hotkey.Register(hotkey.Ctrl, 'Q', func() {
		quit <- true
	})

	hotkey.Start()

	fmt.Println(`
Start hotkey's loop.

Alt-Shift-F1..F12: Print key name.
Ctrl-Q:            Quit`[1:])

	<-quit
}
