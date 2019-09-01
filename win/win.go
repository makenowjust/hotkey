// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// +build windows

// Package hotkey_win is win32api wrapper for hotkey.
package hotkey_win

import (
	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var (
	registerHotKey    *windows.LazyProc
	unregisterHotKey  *windows.LazyProc
	postThreadMessage *windows.LazyProc

	getCurrentThread *windows.LazyProc
	getThreadId      *windows.LazyProc
)

func init() {
	// Library
	libuser32 := windows.NewLazySystemDLL("user32.dll")
	libkernel32 := windows.NewLazySystemDLL("kernel32.dll")

	// Functions
	registerHotKey = libuser32.NewProc("RegisterHotKey")
	unregisterHotKey = libuser32.NewProc("UnregisterHotKey")
	postThreadMessage = libuser32.NewProc("PostThreadMessageW")

	getCurrentThread = libkernel32.NewProc("GetCurrentThread")
	getThreadId = libkernel32.NewProc("GetThreadId")
}

func RegisterHotKey(hwnd win.HWND, id int32, fsModifiers, vk uint32) bool {
	ret, _, _ := registerHotKey.Call(
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk))

	return ret != 0
}

func PostThreadMessage(idThread uint32, msg uint32, wParam, lParam int32) bool {
	ret, _, _ := postThreadMessage.Call(
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam))
	return ret != 0
}

func UnregisterHotKey(hwnd win.HWND, id int32) bool {
	ret, _, _ := unregisterHotKey.Call(
		uintptr(hwnd),
		uintptr(id))

	return ret != 0
}

func GetCurrentThread() win.HANDLE {
	ret, _, _ := getCurrentThread.Call()
	return win.HANDLE(ret)
}

func GetThreadId(thread win.HANDLE) uint32 {
	ret, _, _ := getThreadId.Call(uintptr(thread))
	return uint32(ret)
}
