package resolvePath

import "errors"

var ErrCycleFound = errors.New("cycle found")
var ErrNoPathFound = errors.New("no path found")
