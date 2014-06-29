// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Package hotkey provides HotKey for Go Language.
package hotkey

import (
	"fmt"
	"time"

	"github.com/lxn/win"
)

import "github.com/MakeNowJust/hotkey/win"

// A hotkey.Id is a identity number of registered hotkey.
type Id int32

// A hotkey.Modifier specifies the keyboard modifier for hotkey.Register.
type Modifier uint32

// These are all members of hotkey.Modifier.
const (
	Alt   Modifier = hotkey_win.MOD_ALT
	Ctrl  Modifier = hotkey_win.MOD_CONTROL
	Shift Modifier = hotkey_win.MOD_SHIFT
	Win   Modifier = hotkey_win.MOD_WIN
)

type reservedHotKey struct {
	id              int32
	fsModifiers, vk uint32
}

var (
	started = false

	currentId = int32(1)
	id2handle = make(map[Id]func())

	reservedHotKeys = make([]reservedHotKey, 0)

	threadId     uint32
	chRegister   = make(chan reservedHotKey, 100)
	chUnregister = make(chan int32, 100)
)

// Register a hotkey with modifiers and vk.
// mods are hotkey's modifiers such as hotkey.Alt, hotkey.Ctrl|hotkey.Shift.
// vk is a hotkey's virtual key code. See also
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
func Register(mods Modifier, vk uint32, handle func()) (id Id, err error) {
	reserved := reservedHotKey{
		id:          currentId,
		fsModifiers: uint32(mods),
		vk:          uint32(vk),
	}
	id = Id(currentId)
	id2handle[id] = handle

	currentId += 1

	if started {
		chRegister <- reserved
		hotkey_win.PostThreadMessage(threadId, win.WM_USER, 0, 0)
	} else {
		reservedHotKeys = append(reservedHotKeys, reserved)
	}
	return
}

// Unregister a hotkey from id.
func Unregister(id Id) {
	if started {
		chUnregister <- int32(id)
		delete(id2handle, id)
		hotkey_win.PostThreadMessage(threadId, win.WM_USER, 0, 0)
	} else {
		for idx, reserved := range reservedHotKeys {
			if reserved.id == int32(id) {
				reservedHotKeys[idx] = reservedHotKey{0, 0, 0}
			}
		}
	}
}

// Start hotkey's loop. It is non-blocking.
func Start() <-chan error {
	chErr := make(chan error)
	chThreadId := make(chan uint32)

	go func() {
		// register and reserve to unregister hotkeys
		count := 0
		for _, reserved := range reservedHotKeys {
			if reserved.id == 0 {
				continue
			}

			if !hotkey_win.RegisterHotKey(0, reserved.id, reserved.fsModifiers, reserved.vk) {
				chErr <- fmt.Errorf("failed to register hotkey %v", reserved)
				return
			}
			defer hotkey_win.UnregisterHotKey(0, reserved.id)
			count += 1
		}

		chThreadId <- hotkey_win.GetThreadId(hotkey_win.GetCurrentThread())

		// hotkey's loop
		for {
			select {
			case <-time.After(time.Millisecond * 10):
				var msg win.MSG
				res := win.GetMessage(&msg, 0, 0, 0)

				if res == 0 || res == -1 {
					// TODO: get system error message
					chErr <- nil
					return
				}

				switch msg.Message {
				case win.WM_HOTKEY:
					if handle := id2handle[Id(msg.WParam)]; handle != nil {
						handle()
					}
				default:
					win.TranslateMessage(&msg)
					win.DispatchMessage(&msg)
				}

			case reserved := <-chRegister:
				if !hotkey_win.RegisterHotKey(0, reserved.id, reserved.fsModifiers, reserved.vk) {
					chErr <- fmt.Errorf("failed to register hotkey %v", reserved)
					return
				}
				defer hotkey_win.UnregisterHotKey(0, reserved.id)
				count += 1

			case id := <-chUnregister:
				hotkey_win.UnregisterHotKey(0, id)
				if count -= 1; count == 0 && len(chRegister) == 0 {
					chErr <- nil
					return
				}
			}
		}
	}()

	threadId = <-chThreadId
	started = true
	return chErr
}
