// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.13.1 command runs the go command from Go 1.13.1.
//
// To install, run:
//
//     $ go get golang.org/dl/go1.13.1
//     $ go1.13.1 download
//
// And then use the go1.13.1 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/devel/release.html#go1.13.minor
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.13.1")
}
