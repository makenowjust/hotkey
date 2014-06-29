// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

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
