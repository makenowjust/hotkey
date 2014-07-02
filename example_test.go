// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey_test

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
)

func ExampleRegister() {
	hkey := hotkey.New()

	// Register Ctrl-A
	hkey.Register(hotkey.Ctrl, 'A', func() {
		// Here is a callback of HotKey Ctrl-A
		fmt.Println("Ctrl-A")
	})

	// Register Ctrl-Alt-B
	hkey.Register(hotkey.Ctrl+hotkey.Alt, 'B', func() {
		fmt.Println("Ctrl-Alt-B")
	})

	// Register Shift-Win-F1
	hkey.Register(hotkey.Shift+hotkey.Win, hotkey.F1, func() {
		fmt.Println("Shift-Win-F1")
	})
}

func ExampleUnregister() {
	hkey := hotkey.New()

	// Register Ctrl-A. This callback will call only once.
	var id hotkey.Id
	id, _ = hkey.Register(hotkey.Ctrl, 'A', func() {
		fmt.Println("Ctrl-A")
		hkey.Unregister(id)
	})
}
