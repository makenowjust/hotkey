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
	libuser32 uintptr

	registerHotKey   uintptr
	unregisterHotKey uintptr
)

func init() {
	// Library
	libuser32 = MustLoadLibrary("user32.dll")

	// Functions
	registerHotKey = MustGetProcAddress(libuser32, "RegisterHotKey")
	unregisterHotKey = MustGetProcAddress(libuser32, "UnregisterHotKey")
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

func UnregisterHotKey(hwnd HWND, id int32) bool {
	ret, _, _ := syscall.Syscall(unregisterHotKey, 2,
		uintptr(hwnd),
		uintptr(id),
		0)

	return ret != 0
}
