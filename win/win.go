// Package hotkey_win is win32api wrapper for hotkey.
package hotkey_win

import (
	. "github.com/lxn/win"
	"syscall"
)

const (
	MOD_ALT     = 1
	MOD_CONTROL = 2
	MOD_SHIFT   = 4
	MOD_WIN     = 8
)

var (
	libuser32   uintptr
	libkernel32 uintptr

	registerHotKey    uintptr
	unregisterHotKey  uintptr
	postThreadMessage uintptr

	getCurrentThread uintptr
	getThreadId      uintptr
)

func init() {
	// Library
	libuser32 = MustLoadLibrary("user32.dll")
	libkernel32 = MustLoadLibrary("kernel32.dll")

	// Functions
	registerHotKey = MustGetProcAddress(libuser32, "RegisterHotKey")
	unregisterHotKey = MustGetProcAddress(libuser32, "UnregisterHotKey")
	postThreadMessage = MustGetProcAddress(libuser32, "PostThreadMessageW")

	getCurrentThread = MustGetProcAddress(libkernel32, "GetCurrentThread")
	getThreadId = MustGetProcAddress(libkernel32, "GetThreadId")
}

func RegisterHotKey(hwnd HWND, id int32, fsModifiers, vk uint32) bool {
	ret, _, _ := syscall.Syscall6(registerHotKey, 4,
		uintptr(hwnd),
		uintptr(id),
		uintptr(fsModifiers),
		uintptr(vk),
		0, 0)

	return ret != 0
}

func PostThreadMessage(idThread uint32, msg uint32, wParam, lParam int32) bool {
	ret, _, _ := syscall.Syscall6(postThreadMessage, 4,
		uintptr(idThread),
		uintptr(msg),
		uintptr(wParam),
		uintptr(lParam),
		0, 0)
	return ret != 0
}

func UnregisterHotKey(hwnd HWND, id int32) bool {
	ret, _, _ := syscall.Syscall(unregisterHotKey, 2,
		uintptr(hwnd),
		uintptr(id),
		0)

	return ret != 0
}

func GetCurrentThread() HANDLE {
	ret, _, _ := syscall.Syscall(getCurrentThread, 0, 0, 0, 0)
	return HANDLE(ret)
}

func GetThreadId(thread HANDLE) uint32 {
	ret, _, _ := syscall.Syscall(getThreadId, 1, uintptr(thread), 0, 0)
	return uint32(ret)
}
