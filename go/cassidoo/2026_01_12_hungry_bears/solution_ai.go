package hungryBears

import (
	"cassidoo/utils"
	"sort"
)

func HungryBearsAI(bears []Bear) []string {
	if len(bears) == 0 {
		return []string{}
	}
	sum := utils.Reduce(bears, 0, func(b Bear, acc int) int { return acc + b.hunger })
	average := float64(sum) / float64(len(bears))
	hungry := utils.Filter(bears, func(b Bear) bool { return float64(b.hunger) > average })
	names := utils.Map(hungry, func(b Bear) string { return b.name })
	sort.Strings(names)
	return names
}
