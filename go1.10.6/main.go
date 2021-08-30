// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.10.6 command runs the go command from Go 1.10.6.
//
// To install, run:
//
//     $ go install golang.org/dl/go1.10.6@latest
//     $ go1.10.6 download
//
// And then use the go1.10.6 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/devel/release.html#go1.10.minor
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.10.6")
}
