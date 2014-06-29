// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func registerOnce(mods hotkey.Modifier, vk uint32, handle func()) {
	var id hotkey.Id
	id, _ = hotkey.Register(mods, vk, func() {
		handle()
		hotkey.Unregister(id)
	})
}

func main() {
	for _, vk := range "QUIT" {
		registerOnce(hotkey.Ctrl, uint32(vk), func(vk rune) func() {
			return func() {
				fmt.Printf("Push Ctrl-%c\n", vk)
			}
		}(vk))
	}

	chErr := hotkey.Start()

	fmt.Println(`
Start hotkey's loop.
Push Ctrl-Q, U, I and T to quit.
`[1:])

	if err := <-chErr; err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("QUIT!")
}
