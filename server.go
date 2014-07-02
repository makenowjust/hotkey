// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey

type server interface {
	register(fsModifiers, vk uint32, handle func()) (Id, error)
	unregister(id int32)
	stop()
	isStop() bool

	// For debugging
	useDebugLog()
}

// For test
var newServer func() server
