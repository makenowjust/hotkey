// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey

import (
	"fmt"
	"runtime"

	"github.com/lxn/win"
)

import "github.com/MakeNowJust/hotkey/win"

type server interface {
	register(fsModifiers, vk uint32, handle func()) (Id, error)
	unregister(id int32)
	stop()
	isStop() bool

	// For debugging
	useDebugLog()
}

const (
	msgRegister = iota
	msgUnregister
	msgStop
)

type message struct {
	msgType         int
	id              int32
	fsModifiers, vk uint32

	// Channels as result notifier
	chId  chan Id
	chErr chan error
}

type serverImpl struct {
	chMsg     chan *message
	id2handle map[Id]func()
	threadId  uint32
	stopFlag  bool

	sendStopFlag bool

	// For debugging
	debug debugT
}

// Hotkey's id manager. It is thread safe.
var globalId = func() <-chan int32 {
	globalId := make(chan int32, 1)
	go func() {
		for i := int32(1); ; i++ {
			globalId <- i
		}
	}()
	return globalId
}()

// For test
var newServer = func() server {
	svr := new(serverImpl)
	svr.chMsg = make(chan *message, 100)
	svr.id2handle = make(map[Id]func())
	svr.debug = debugT(false)

	chThreadId := make(chan uint32)
	svr.debug.Log("Start hotkey's loop")
	go func() {
		svr.debug.Log("Lock thread for win32api")
		runtime.LockOSThread()

		svr.debug.Log("Send a thread id")
		chThreadId <- hotkey_win.GetThreadId(hotkey_win.GetCurrentThread())

		svr.debug.Log("The main of hotkey's loop")
		for {
			select {
			// Has message
			case msg := <-svr.chMsg:
				svr.debug.Log("Received message in hotkey's loop", msg)
				switch msg.msgType {
				case msgRegister:
					svr.debug.Log("Register message", msg)
					id := <-globalId
					if !hotkey_win.RegisterHotKey(0, id, msg.fsModifiers, msg.vk) {
						// TODO: Get system error message
						msg.chErr <- fmt.Errorf("failed to register hotkey {mods=%d, vk=%d}", msg.fsModifiers, msg.vk)
						break
					}
					defer func() {
						svr.debug.Log("defer Unregister", id)
						hotkey_win.UnregisterHotKey(0, id)
					}()
					msg.chId <- Id(id)
					runtime.Gosched()

				case msgUnregister:
					svr.debug.Log("Unregister message", msg)
					hotkey_win.UnregisterHotKey(0, msg.id)
					msg.chErr <- nil

				case msgStop:
					svr.debug.Log("Stop message", msg)
					svr.stopFlag = true
					svr.sendStop()
					msg.chErr <- nil
					return
				}

			// No message
			default:
				svr.debug.Log("Wait hotkey message")
				var msg win.MSG
				res := win.GetMessage(&msg, 0, 0, 0)

				if res == 0 || res == -1 {
					// TODO: Get system error message
					svr.stopFlag = true
					svr.sendStop()
					return
				}

				switch msg.Message {
				// Hotkey's command
				case win.WM_HOTKEY:
					svr.debug.Log("WM_HOTKEY", msg.WParam)
					if handle, ok := svr.id2handle[Id(msg.WParam)]; ok {
						handle()
					}

				default:
					svr.debug.Log("Other message")
					win.TranslateMessage(&msg)
					win.DispatchMessage(&msg)
				}
			}
		}
	}()

	// Recive a thread id
	svr.threadId = <-chThreadId

	return svr
}

func (svr *serverImpl) register(fsModifiers, vk uint32, handle func()) (id Id, err error) {
	if svr.stopFlag {
		err = fmt.Errorf("already stoped hotkey's loop")
		return
	}

	var msg message
	msg.msgType = msgRegister
	msg.fsModifiers = fsModifiers
	msg.vk = vk

	msg.chId = make(chan Id)
	msg.chErr = make(chan error)

	svr.chMsg <- &msg
	hotkey_win.PostThreadMessage(svr.threadId, win.WM_USER, 0, 0)

	// Wait
	select {
	case id = <-msg.chId:
		svr.debug.Log("Register success", id)
		svr.id2handle[id] = handle
	case err = <-msg.chErr:
	}
	return
}

func (svr *serverImpl) unregister(id int32) {
	if svr.stopFlag {
		return
	}

	var msg message
	msg.msgType = msgUnregister
	msg.id = id

	msg.chErr = make(chan error)

	svr.chMsg <- &msg
	hotkey_win.PostThreadMessage(svr.threadId, win.WM_USER, 0, 0)

	// Wait
	<-msg.chErr

	svr.debug.Log("Unregister done")
}

func (svr *serverImpl) stop() {
	if svr.stopFlag {
		return
	}

	var msg message
	msg.msgType = msgStop

	msg.chErr = make(chan error)

	svr.chMsg <- &msg
	hotkey_win.PostThreadMessage(svr.threadId, win.WM_USER, 0, 0)

	// Wait
	<-msg.chErr

	svr.debug.Log("Stop done")
}

func (svr *serverImpl) sendStop() {
	if svr.sendStopFlag {
		svr.debug.Log("sendStop")
		svr.sendStopFlag = true

		for len(svr.chMsg) >= 1 {
			msg := <-svr.chMsg

			if msg.msgType == msgRegister {
				msg.chErr <- fmt.Errorf("already stoped hotkey's loop")
			} else {
				msg.chErr <- nil
			}
		}
	}
}

func (svr *serverImpl) isStop() bool {
	return svr.stopFlag
}

func (svr *serverImpl) useDebugLog() {
	svr.debug = debugT(true)
}
