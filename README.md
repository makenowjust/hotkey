#hotkey

##About

This library provides HotKey for Go Language on Windows.
(includes win32api wrapper of `RegisterHotKey`, `UnregisterHotKey` and more)

##Get Started

Now run `go get github.com/MakeNowJust/hotkey`.

##Import

```go
import "github.com/MakeNowJust/hotkey"
```

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
  hkey := hotkey.New()

	quit := make(chan bool)

	hkey.Register(hotkey.Ctrl, 'Q', func() {
		fmt.Println("Quit")
		quit <- true
	})

	fmt.Println("Start hotkey's loop")
	fmt.Println("Push Ctrl-Q to escape and quit")
	<-quit
}
```

and run `go run minimal.go`

More examples exists `example` directory. Let's see.

##Document

See <https://godoc.org/github.com/MakeNowJust/hotkey>.

##License

This software is released under the MIT License, see LICENSE.

