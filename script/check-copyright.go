// Copyright (C) 2015 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// +build ignore

// Checks for files missing copyright notice
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// File extensions to check
var checkExts = map[string]bool{
	".go": true,
}

// Valid copyright headers, searched for in the top five lines in each file.
var copyrightRegexps = []string{
	`Copyright`,
	`package auto`,
	`automatically generated by genxdr`,
	`generated by protoc`,
}

var copyrightRe = regexp.MustCompile(strings.Join(copyrightRegexps, "|"))

func main() {
	flag.Parse()
	for _, dir := range flag.Args() {
		err := filepath.Walk(dir, checkCopyright)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func checkCopyright(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !info.Mode().IsRegular() {
		return nil
	}
	if !checkExts[filepath.Ext(path)] {
		return nil
	}

	fd, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for i := 0; scanner.Scan() && i < 5; i++ {
		if copyrightRe.MatchString(scanner.Text()) {
			return nil
		}
	}

	return fmt.Errorf("Missing copyright in %s?", path)
}
