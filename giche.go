// giche  a cross platform which tool written in Go

package main

import (
	"flag";
	"fmt";
	"os"
	"syscall"
	"strings"
	"time"
)

var compTime *time.Time = time.LocalTime()
var version string = "1.0"
var sepChar string = `:`
var sepPath string = `/`
var eol string = "\n"
var aFlag bool
var sFlag bool
var hFlag bool
var helpFlag bool
var winFlag bool = false
var allMsg string = "List all executable instances found rather than just the first one."
var statusMsg string = "Output 'Found' if any of the executables were found and 'None' if none were found"
var helpMsg string = "Print this usage message"

func init() {
	flag.BoolVar(&aFlag, "a", false, allMsg)
	flag.BoolVar(&sFlag, "s", false, statusMsg)
	flag.BoolVar(&hFlag, "h", false, helpMsg)
	flag.BoolVar(&helpFlag, "help", false, helpMsg)
	if syscall.OS == "windows" {
		sepChar = `;`
		sepPath = `\`
		eol = "\r\n"
		winFlag = true
	}
}

var usage = func() {
	fmt.Fprintf(os.Stderr, "Giche  a cross platform which tool written in Go\n")
	fmt.Fprintf(os.Stderr, "v%s, %s\n", version, compTime.Format(time.RFC822))
	fmt.Fprintf(os.Stderr, "\nUsage: %s [-l|-s|-h|-help] file ... \n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t-l  %s \n", allMsg)
	fmt.Fprintf(os.Stderr, "\t-s  %s \n", statusMsg)
	fmt.Fprintf(os.Stderr, "\t-h  %s \n", helpMsg)
	fmt.Fprintf(os.Stderr, "\t-help  %s \n", helpMsg)
	os.Exit(0)
}

func chkStat(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
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
					if strings.HasSuffix(f, e) {
						ff := path + sepPath + file
						if chkStat(ff) {
							if sFlag {
								userMsg = "Found"
								break outer
							}
							if aFlag {
								userMsg += file + eol
								continue
							}
							userMsg += file + eol
							continue outer
						}
					}
				}
				for _, e := range exts {
					ff := path + sepPath + file + e
					if chkStat(ff) {
						if sFlag {
							userMsg = "Found"
							break outer
						}
						if aFlag {
							userMsg += (file + eol)
							continue
						}
						userMsg += (file + eol)
						continue outer
					}
				}
			} else {
				f := path + sepPath + file
				if chkStat(f) {
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
	fmt.Println(userMsg)
}

func prolog(files []string) {
	path := os.Getenv("PATH")

	pathext := os.Getenv("PATHEXT")
	exts := []string{}
	if pathext != "" {
		exts = strings.Split(strings.ToLower(pathext), sepChar, -1)
		for i, e := range exts {
			if len(e) < 1 || e[0] != '.' {
				exts[i] = `.` + e
			}
		}
	}
	paths := []string{}
	if path != "" {
		paths = strings.Split(path, sepChar, -1)
	}
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
