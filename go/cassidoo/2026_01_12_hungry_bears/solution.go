package hungryBears

import (
	"cassidoo/utils"
	"slices"
)

func HungryBears(bears []Bear) []string {
	average := utils.Reduce(bears, 0, func(b Bear, sum int) int { return b.hunger + sum }) / len(bears)
	filtered := utils.Filter(bears, func(b Bear) bool { return b.hunger > average })
	names := utils.Map(filtered, func(b Bear) string { return b.name })
	slices.Sort(names)
	return names
}
