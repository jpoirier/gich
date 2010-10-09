// gich  a cross platform which tool written in Go

// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"strings"
)

func init() {
	flag.BoolVar(&aFlag, "a", false, allMsg)
	flag.BoolVar(&sFlag, "s", false, statMsg)
	flag.BoolVar(&hFlag, "h", false, prntMsg)
	flag.BoolVar(&helpFlag, "help", false, prntMsg)
	if syscall.OS == "windows" {
		sepChar = `;`
		// TODO: Is this necessary? windows handles forward slashes
		// to what level, ie is it different between user and
		// kernel level, and does it matter?
		sepPath = `\`
		// TODO: Is this necessary? Any difference between
		// cmd.exe and command.com?
		eol = "\r\n"
		winFlag = true
	}
}

var usage = func() {
	fmt.Print(helpMsg)
	os.Exit(0)
}

func process(files, paths, exts []string) {
	userMsg := ""
outer:	for _, file := range files {
		if strings.Index(file, `\`) >= 0 || strings.Index(file, `/`) >= 0 {
			continue
		}
inner:		for _, path := range paths {
			if len(exts) != 0 {
				f := strings.ToLower(file)
				for _, e := range exts {
					ff := path + sepPath + file
					if !strings.HasSuffix(f, e) {
						ff += e
					}
					if _, err := os.Stat(ff); err == nil {
						if sFlag {
							userMsg = "Found"
							break outer
						}
						if aFlag {
							userMsg += ff + eol
							continue
						}
						userMsg += ff + eol
						continue outer
					}
				}
			} else {
				f := path + sepPath + file
				if _, err := os.Stat(f); err == nil {
					if sFlag {
						userMsg = "Found"
						break outer
					}
					if aFlag {
						userMsg += (f + eol)
						continue
					}
					userMsg += f + eol
					continue outer
				}
			}
		}
	}
	if sFlag && userMsg == "" {
		userMsg = "None"
	}
	fmt.Print(userMsg)
}

func prolog(files []string) {
	path := os.Getenv("PATH")
	if path == "" {
		return
	}
	paths := []string{}
	exts := []string{}
	if winFlag {
// TODO: Check for functionality differences between the
// DOS (command.com) and NT (cmd.exe) shells
//		path = strings.Replace(path, `\`, `\\`, -1)
		pathext := os.Getenv("PATHEXT")
		if pathext != "" {
			exts = strings.Split(strings.ToLower(pathext), sepChar, -1)
			for i, e := range exts {
				if e == "" || e[0] != '.' {
					exts[i] = "." + e
				}
			}
		}
// TODO: Check for functionality differences between the
// DOS (command.com) and NT (cmd.exe) shells
//		paths = strings.Split(path, sepChar, -1)
//		for i, p := range paths {
//			paths[i] = `"` + p + `"`
//		}
	}
	paths = strings.Split(path, sepChar, -1)
	process(files, paths, exts)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(os.Args) == 1 || hFlag || helpFlag || aFlag && sFlag {
		usage()
	}
	x := 1
	if aFlag || sFlag {
		x += 1
	}
	if (len(os.Args) - x) < 1 {
		usage()
	}
	prolog(os.Args[x:])
	os.Exit(0)
}
