// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.26rc3 command runs the go command from Go 1.26rc3.
//
// To install, run:
//
//	$ go install golang.org/dl/go1.26rc3@latest
//	$ go1.26rc3 download
//
// And then use the go1.26rc3 command as if it were your normal go
// command.
//
// See the release notes at https://tip.golang.org/doc/go1.26.
//
// File bugs at https://go.dev/issue/new.
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("go1.26rc3")
}
