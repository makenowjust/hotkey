#hotkey

##About

This library add HotKey in your Go Programming Language program on Windows.
(includes win32api wrapper of `RegisterHotKey` and `UnregisterHotKey`)

##Get Started

Now run `go get github.com/MakeNowJust/hotkey`.

##Using

Such a minimal example:

`minimal.go`

```go
package main

import (
	"fmt"

	"github.com/MakeNowJust/hotkey"
)

func main() {
	quit := make(chan bool)

	hotkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit")
		quit <- true
	})

	hotkey.Start()

	fmt.Println("Start hotkey's loop")
	fmt.Println("Push Ctrl-Q to escape and quit")
	<-quit
}
```

and run `go run minimal.go`

##License

This software is released under the MIT License, see LICENSE.

