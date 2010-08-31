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

