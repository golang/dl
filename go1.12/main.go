// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.12 command runs the go command from Go 1.12.
//
// To install, run:
//
//     $ go install golang.org/dl/go1.12@latest
//     $ go1.12 download
//
// And then use the go1.12 command as if it were your normal go
// command.
//
// See the release notes at https://golang.org/doc/go1.12
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.12")
}
