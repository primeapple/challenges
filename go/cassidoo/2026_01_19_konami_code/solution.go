package konamiCode

import "fmt"

func KonamiCode(sequence string) (map[string]string, error) {
	runeSequence := []rune(sequence)
	for i := 0; i < len(runeSequence)-9; i++ {
		mapping, err := getKonamiMapping(runeSequence[i : i+10])
		if err != nil {
			continue
		}
		return mapping, nil
	}
	return nil, ErrNoMappingFound
}

func getKonamiMapping(seq []rune) (map[string]string, error) {
	if len(seq) != 10 {
		panic(fmt.Sprintf("Invalid length of sequence given, expected 10, got %d", len(seq)))
	}

	// Check the same characters
	if seq[0] != seq[1] {
		return nil, ErrNoMappingFound
	}
	if seq[2] != seq[3] {
		return nil, ErrNoMappingFound
	}
	if seq[4] != seq[6] {
		return nil, ErrNoMappingFound
	}
	if seq[5] != seq[7] {
		return nil, ErrNoMappingFound
	}

	uMapping := seq[0]
	dMapping := seq[2]
	lMapping := seq[4]
	rMapping := seq[5]
	bMapping := seq[8]
	aMapping := seq[9]

	mappings := []rune{
		uMapping,
		dMapping,
		lMapping,
		rMapping,
		bMapping,
		aMapping,
	}

	// check that all are different
	for i, value := range mappings {
		for j := i + 1; j < len(mappings); j++ {
			if value == mappings[j] {
				return nil, ErrNoMappingFound
			}
		}
	}

	return map[string]string{
		string(uMapping): "U",
		string(dMapping): "D",
		string(lMapping): "L",
		string(rMapping): "R",
		string(bMapping): "B",
		string(aMapping): "A",
	}, nil
}
