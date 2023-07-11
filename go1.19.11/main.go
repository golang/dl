// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.19.11 command runs the go command from Go 1.19.11.
//
// To install, run:
//
//	$ go install golang.org/dl/go1.19.11@latest
//	$ go1.19.11 download
//
// And then use the go1.19.11 command as if it were your normal go
// command.
//
// See the release notes at https://go.dev/doc/devel/release#go1.19.11.
//
// File bugs at https://go.dev/issue/new.
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.19.11")
}
