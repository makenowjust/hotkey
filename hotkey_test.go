// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey_test

import (
	"testing"
)

import "github.com/MakeNowJust/hotkey"

type testRegister struct {
	fsModifiers hotkey.Modifier
	vk uint32
}

var	registers = []testRegister {
	{hotkey.Ctrl , 'A'},
	{hotkey.Ctrl+hotkey.Alt , 'B'},
	{hotkey.Shift, 'C'},
	{hotkey.Shift*hotkey.Win, 'D'},
	{hotkey.Win  , hotkey.F1},
}


func TestRegister(t *testing.T) {
	count := 0
	hotkey.MockRegister = func (fsModifiers, vk uint32, handle func ()) (hotkey.Id, error) {
		if reg := registers[count]; reg.fsModifiers != hotkey.Modifier(fsModifiers) || reg.vk != vk {
			t.Errorf("unexpected to register call (%d, %d)", fsModifiers, vk)
		}
		count += 1
		return hotkey.Id(count), nil
	}

	man := hotkey.New()
	for _, reg := range registers {
		man.Register(reg.fsModifiers, reg.vk, func () {
			// nothing
		})
	}
	if count != len(registers) {
		t.Error("no enough to call count register")
	}
}

func TestUnregister(t *testing.T) {
	ids := make([]hotkey.Id, len(registers)/2)
	count := 0
	hotkey.MockRegister = func (fsModifiers, vk uint32, handle func ()) (hotkey.Id, error) {
		if reg := registers[count]; reg.fsModifiers != hotkey.Modifier(fsModifiers) || reg.vk != vk {
			t.Errorf("unexpected to register call (%d, %d)", fsModifiers, vk)
		}
		count += 1
		return hotkey.Id(count), nil
	}
	idx1 := 0
	hotkey.MockUnregister = func (id int32) {
		if ids[idx1] != hotkey.Id(id) {
			t.Errorf("unexpected to unregister call (%d)", id)
		}
		idx1 += 1
	}

	idx2 := 0
	man := hotkey.New()
	for _, reg := range registers {
		id, _ := man.Register(reg.fsModifiers, reg.vk, func () {
			// nothing
		})
		if idx2 < len(ids) {
			ids[idx2] = id
			idx2 += 1
		}
	}
	if count != len(registers) {
		t.Error("no enough to call count register")
	}

	for _, id := range ids {
		man.Unregister(id)
	}
	if idx1 != idx2 {
		t.Error("no enough to call count unregister")		
	}
}
