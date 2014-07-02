// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

var hkey = hotkey.New()

func registerOnce(mods hotkey.Modifier, vk uint32, handle func()) {
	var id hotkey.Id
	id, _ = hkey.Register(mods, vk, func() {
		handle()
		hkey.Unregister(id)
	})
}

func registerLoop(str []rune, idx int, finish chan bool) {
	if idx < len(str) {
		registerOnce(hotkey.Ctrl, uint32(str[idx]), func() {
			fmt.Printf("Push Ctrl-%c\n", str[idx])
			registerLoop(str, idx+1, finish)
		})
	} else {
		finish <- true
	}
}

func main() {
	quit := make(chan bool)

	registerLoop([]rune("QUIT"), 0, quit)

	fmt.Println(`
Start hotkey's loop.
Push Ctrl-Q, U, I and T to quit.
`[1:])

	<-quit
	fmt.Println("QUIT!")
}
