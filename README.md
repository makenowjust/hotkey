# hotkey [![Go Reference](https://pkg.go.dev/badge/github.com/MakeNowJust/hotkey.svg)](https://pkg.go.dev/github.com/MakeNowJust/hotkey)

## About

This library provides HotKey for Go Language on Windows (including win32api wrapper, such as `RegisterHotKey`, `UnregisterHotKey`).

## Get Started

```console
$ go get github.com/MakeNowJust/hotkey
```

## Import

```go
import "github.com/MakeNowJust/hotkey"
```

## Usage

The below is a minimal example.

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

Let's see the [`example/`](example/) for more examples.

## License

This software is released under the MIT License.
<http://opensource.org/licenses/mit-license.php>

Copyright (c) 2014-2023 Hiroya Fujinami (a.k.a. TSUYUSATO "MakeNowJust" Kitsune)
