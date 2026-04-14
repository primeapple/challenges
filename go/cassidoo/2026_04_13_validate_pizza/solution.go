package validatePizza

func ValidatePizza(layers []string, rules [][]string) (bool, []string) {
	layerIndexes := make(map[string]int)
	for index, layer := range layers {
		layerIndexes[layer] = index
	}

	for _, rule := range rules {
		firstIndex := layerIndexes[rule[0]]
		secondIndex := layerIndexes[rule[1]]
		if firstIndex > secondIndex {
			return false, rule
		}
	}
	return true, nil
}
