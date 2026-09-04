// Command create_mod produces a Go module source zip from a directory, using
// golang.org/x/mod/zip — Go's own module-zip implementation, so the resulting
// member set matches what proxy.golang.org serves by construction (no .git, no
// vendor/, no nested modules).
//
// It lives under .github/ so that the seal-build workflow's staging step, which
// excludes .github from the tree it zips, keeps it out of the published module
// zip. It is compiled from its own throwaway module (see the workflow) so that
// golang.org/x/mod never becomes a dependency of golang.org/x/net itself.
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
