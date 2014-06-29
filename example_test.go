// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey_test

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
)

func ExampleRegister() {
	// Register Ctrl-A
	hotkey.Register(hotkey.Ctrl, 'A', func () {
		// Here is a callback of HotKey Ctrl-A
		fmt.Println("Ctrl-A")
	})

	// Register Ctrl-Alt-B
	hotkey.Register(hotkey.Ctrl+hotkey.Alt, 'B', func () {
		fmt.Println("Ctrl-Alt-B")
	})

	// Register Shift-Win-F1
	hotkey.Register(hotkey.Shift+hotkey.Win, hotkey.F1, func () {
		fmt.Println("Shift-Win-F1")
	})
}

func ExampleUnregister() {
	// Register Ctrl-A. This callback will call only once.
	var id hotkey.Id
	id = hotkey.Register(hotkey.Ctrl, 'A', func () {
		fmt.Println("Ctrl-A")
		hotkey.Unregister(id)
	})
}


