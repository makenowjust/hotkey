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

func main() {
	quit := make(chan bool)

	count := 0
	for _, vk := range "QUIT" {
		registerOnce(hotkey.Ctrl, uint32(vk), func(vk rune) func() {
			return func() {
				fmt.Printf("Push Ctrl-%c\n", vk)
				if count += 1; count == 4 {
					quit <- true
				}
			}
		}(vk))
	}

	fmt.Println(`
Start hotkey's loop.
Push Ctrl-Q, U, I and T to quit.
`[1:])

	<-quit
	fmt.Println("QUIT!")
}
