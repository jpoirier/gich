/* giche a "which" utility for windows.
 *
 * This package provides and interface to the storage containers and
 * routines in gocache.  The storage routines work with anything that
 * conforms to the Entity interface.  By default the Item type will be
 * used for storage.  The package also loads data from a fixture and
 * writes it's current data into a fixture.
 *
 * Copyright 2010 by Joseph D Poirier. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package main

import (
	"flag";
	"fmt";
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s root-directory\n", os.Args[0]);
	flag.PrintDefaults();
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if *help {
		usage()
	}

	if flag.NArg() > 0 {
		tag := flag.Args()[0]
		var err os.Error
		db, err = sqlite.Open(*dbfile)
		errchk(err)
		errchk(scanDir(*picsdir, tag))
		log.Stdout("Scanning of " + *picsdir + " complete.")
		db.Close()
		return
	}

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "go main-program [arg0 [arg1 ...]]")
		os.Exit(1)
	}


	tw := twitter.NewTwitter(*username, *password);

	switch flag.Arg(0) {
	case "@", "mentions":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.Mentions()));
	case "", "friends":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.FriendsTimeline()));
	case "u", "user":
		requireLogin();
		os.Stdout.WriteString(checkForError(tw.UserTimeline()));
	case "public":
		os.Stdout.WriteString(checkForError(tw.PublicTimeline()))
	case "p", "post":
		requireLogin();
		s := "";
		for i := 1; i < flag.NArg(); i++ {
			if i > 1 {
				s += " "
			}
			s += flag.Arg(i);
		}
		err := tw.UpdateStatus(s);
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err);
			os.Exit(1);			
		}
	}
	os.Exit(0);

}
