// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.20.13 command runs the go command from Go 1.20.13.
//
// To install, run:
//
//	$ go install golang.org/dl/go1.20.13@latest
//	$ go1.20.13 download
//
// And then use the go1.20.13 command as if it were your normal go
// command.
//
// See the release notes at https://go.dev/doc/devel/release#go1.20.13.
//
// File bugs at https://go.dev/issue/new.
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.20.13")
}
