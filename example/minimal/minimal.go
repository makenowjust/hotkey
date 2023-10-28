// This software is released under the MIT License.
// <http://opensource.org/licenses/mit-license.php>
//
// Copyright (c) 2014-2023 Hiroya Fujinami (a.k.a. TSUYUSATO "MakeNowJust" Kitsune)

package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func main() {
	hkey := hotkey.New()

	quit := make(chan bool)

	hkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit")
		quit <- true
	})

	fmt.Println("Start hotkey's loop")
	fmt.Println("Push Ctrl-Q to escape and quit")
	<-quit
}
