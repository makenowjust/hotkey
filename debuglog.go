// Copyright (c) 2014 TSUYUSATO Kitsune
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package hotkey

import "log"

type debugT bool

func (d debugT) Log(fmt string, args ...interface{}) {
	if d {
		log.Printf(fmt, args...)
	}
}
