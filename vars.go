// giche  a cross platform which tool written in Go

package main

import (
	"time"
)

const VERSION = "1.1"

var (
	compTime *time.Time = time.LocalTime()
	aFlag bool
	sFlag bool
	hFlag bool
	helpFlag bool
	winFlag = false
	eol 	= "\n"
	sepChar = `:`
	sepPath = `/`
	allMsg  = "List all executable instances found rather than just the first one."
	statMsg = "Output 'Found' if any of the executables were found and 'None' if none were found."
	prntMsg = "Print this usage message."
	compMsg = compTime.Format(time.RFC822)
	helpMsg = "Giche " + "(v" + VERSION + ")  " +
			"a cross platform which tool written in Go\n" +
			"Usage: insert [-l|-s|-h|-help] file ... \n" +
			"\t-l  "   + allMsg   + " \n" +
			"\t-s  "   + statMsg  + " \n" +
			"\t-h  "   + prntMsg  + " \n" +
			"\t-help " + prntMsg  + " \n"
)
