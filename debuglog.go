// This software is released under the MIT License.
// <http://opensource.org/licenses/mit-license.php>
//
// Copyright (c) 2014-2023 Hiroya Fujinami (a.k.a. TSUYUSATO "MakeNowJust" Kitsune)

package hotkey

import "log"

type debugT bool

func (d debugT) Log(fmt string, args ...interface{}) {
	if d {
		log.Printf(fmt, args...)
	}
}
