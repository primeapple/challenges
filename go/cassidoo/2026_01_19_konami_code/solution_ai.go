package konamiCode

func KonamiCodeAI(str string) (map[string]string, error) {
	n := len(str)
	for i := 0; i <= n-10; i++ {
		sub := str[i : i+10]
		if sub[0] != sub[1] || sub[2] != sub[3] || sub[4] != sub[6] || sub[5] != sub[7] {
			continue
		}
		chars := make(map[byte]bool)
		chars[sub[0]] = true
		chars[sub[2]] = true
		chars[sub[4]] = true
		chars[sub[5]] = true
		chars[sub[8]] = true
		chars[sub[9]] = true
		if len(chars) != 6 {
			continue
		}
		mapping := make(map[string]string)
		mapping[string(sub[0])] = "U"
		mapping[string(sub[2])] = "D"
		mapping[string(sub[4])] = "L"
		mapping[string(sub[5])] = "R"
		mapping[string(sub[8])] = "B"
		mapping[string(sub[9])] = "A"
		return mapping, nil
	}
	return nil, ErrNoMappingFound
}
