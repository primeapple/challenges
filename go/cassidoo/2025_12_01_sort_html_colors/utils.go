package sortHtmlColors

import (
	"math/rand"
	"slices"
)

var HtmlColors = []string{
	"000000",
	"000080",
	"0000FF",
	"008000",
	"008080",
	"00FF00",
	"00FFFF",
	"800000",
	"800080",
	"808000",
	"808080",
	"C0C0C0",
	"FF0000",
	"FF00FF",
	"FFFF00",
	"FFFFFF",
}

func GetColorsInRandomOrder() []string {
	copiedColors := slices.Clone(HtmlColors)

	rand.Shuffle(len(HtmlColors), func(i, j int) {
		copiedColors[i], copiedColors[j] = copiedColors[j], copiedColors[i]
	})

	return copiedColors
}
