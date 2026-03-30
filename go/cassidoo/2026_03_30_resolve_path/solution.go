package resolvePath

func ResolvePath(filesystem map[string]*string, path string) (string, error) {
	seenPaths := map[string]struct{}{}
	current := path

	for {
		next, pathExists := filesystem[current]
		if !pathExists {
			return "", ErrNoPathFound
		}
		if next == nil {
			return current, nil
		}

		_, seenBefore := seenPaths[*next]
		if seenBefore {
			return "", ErrCycleFound
		} else {
			seenPaths[*next] = struct{}{}
		}

		current = *next
	}
}
