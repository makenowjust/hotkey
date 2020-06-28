// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

// Package hotkey provides HotKey for Go Language.
package hotkey

import "github.com/MakeNowJust/hotkey/win"

// A hotkey.Id is a identity number of registered hotkey.
type Id int32

// A hotkey.Modifier specifies the keyboard modifier for hotkey.Register.
type Modifier uint32

// These are all members of hotkey.Modifier.
const (
	None  Modifier = hotkey_win.MOD_NONE
	Alt   Modifier = hotkey_win.MOD_ALT
	Ctrl  Modifier = hotkey_win.MOD_CONTROL
	Shift Modifier = hotkey_win.MOD_SHIFT
	Win   Modifier = hotkey_win.MOD_WIN
)

// A hotkey's manager.
type Manager struct {
	svr server
}

// Create hotkey's manager and Start hotkey's loop. It is non-blocking.
func New() (man *Manager) {
	man = new(Manager)
	man.svr = newServer()

	return
}

// Register a hotkey with modifiers and vk on man.
//
// mods are hotkey's modifiers such as hotkey.Alt, hotkey.Ctrl+hotkey.Shift.
//
// vk is a hotkey's virtual key code. See also
// http://msdn.microsoft.com/en-us/library/windows/desktop/dd375731(v=vs.85).aspx
func (man *Manager) Register(mods Modifier, vk uint32, handle func()) (id Id, err error) {
	id, err = man.svr.register(uint32(mods), vk, handle)
	return
}

// Unregister a hotkey from id.
func (man *Manager) Unregister(id Id) {
	man.svr.unregister(int32(id))
}

// Stop hotkey's loop.
func (man *Manager) Stop() {
	man.svr.stop()
}

// Check if hotkey's loop is stopping.
func (man *Manager) IsStop() bool {
	return man.svr.isStop()
}

// For debugging.
func (man *Manager) UseDebugLog() *Manager {
	man.svr.useDebugLog()
	return man
}
