// gich  a cross platform which tool written in Go

// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package main

import (
)

const VERSION = "1.2"

var (
	aFlag 	 bool
	sFlag 	 bool
	hFlag 	 bool
	helpFlag bool
	winFlag = false
	eol 	= "\n"
	sepChar = `:`
	sepPath = `/`
	allMsg  = "List all executable instances found rather than just the first one."
	statMsg = "Output 'Found' if any of the executables were found and 'None' if none were found."
	prntMsg = "Print this usage message."
	helpMsg = "gich " + "(v" + VERSION + ")  " +
			"a cross platform which tool written in Go\n" +
			"Usage: gich [-l|-s|-h|-help] file ... \n" +
			"\t-l  "   + allMsg   + " \n" +
			"\t-s  "   + statMsg  + " \n" +
			"\t-h  "   + prntMsg  + " \n" +
			"\t-help " + prntMsg  + " \n"
)
