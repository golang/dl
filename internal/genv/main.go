// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The genv command generates version-specific go command source files.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: genv <version>...")
	os.Exit(2)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}
	for _, version := range os.Args[1:] {
		if !strings.HasPrefix(version, "go") {
			failf("version names should have the 'go' prefix")
		}
		var buf bytes.Buffer
		if err := mainTmpl.Execute(&buf, struct {
			Year                int
			Version             string // "go1.5.3rc2"
			VersionNoPatch      string // "go1.5"
			CapitalSpaceVersion string // "Go 1.5"
			DocHost             string // "golang.org" or "tip.golang.org" for rc/beta
		}{
			Year:                time.Now().Year(),
			Version:             version,
			VersionNoPatch:      versionNoPatch(version),
			DocHost:             docHost(version),
			CapitalSpaceVersion: strings.Replace(version, "go", "Go ", 1),
		}); err != nil {
			failf("mainTmpl.execute: %v", err)
		}
		path := filepath.Join(os.Getenv("GOPATH"), "src/golang.org/dl", version, "main.go")
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			failf("%v", err)
		}
		if err := ioutil.WriteFile(path, buf.Bytes(), 0666); err != nil {
			failf("ioutil.WriteFile: %v", err)
		}
		fmt.Println("Wrote", path)
		if err := exec.Command("gofmt", "-w", path).Run(); err != nil {
			failf("could not gofmt file %q: %v", path, err)
		}
	}
}

func docHost(ver string) string {
	if strings.Contains(ver, "rc") || strings.Contains(ver, "beta") {
		return "tip.golang.org"
	}
	return "golang.org"
}

func versionNoPatch(ver string) string {
	rx := regexp.MustCompile(`^(go\d+\.\d+)($|rc|beta|\.)`)
	m := rx.FindStringSubmatch(ver)
	if len(m) < 2 {
		failf("unrecognized version %q", ver)
	}
	return m[1]
}

func failf(format string, args ...interface{}) {
	if len(format) == 0 || format[len(format)-1] != '\n' {
		format += "\n"
	}
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

var mainTmpl = template.Must(template.New("main").Parse(`// Copyright {{.Year}} The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The {{.Version}} command runs the go command from {{.CapitalSpaceVersion}}.
//
// To install, run:
//
//     $ go get golang.org/dl/{{.Version}}
//     $ {{.Version}} download
//
// And then use the {{.Version}} command as if it were your normal go
// command.
//
// See the release notes at https://{{.DocHost}}/doc/{{.VersionNoPatch}}
//
// File bugs at https://golang.org/issues/new
package main

import "golang.org/dl/internal/version"

func main() {
	version.Run("{{.Version}}")
}
`))
