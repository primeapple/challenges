package main

import "strings"

func GetCharAfterVimCommandsAI(text string, commands string) string {
	lines := strings.Split(text, "\n")

	x := 0
	y := 0

	for _, rune := range commands {
		switch string(rune) {
		case "h":
			x = max(0, x-1)
		case "j":
			if y+1 < len(lines) {
				y++
				// Adjust x if new line is shorter
				if x >= len(lines[y]) {
					x = len(lines[y]) - 1
				}
			}
		case "k":
			y = max(0, y-1)
			// Adjust x if new line is shorter
			if x >= len(lines[y]) {
				x = len(lines[y]) - 1
			}
		case "l":
			if x+1 < len(lines[y]) {
				x++
			}
		}
	}

	return string(lines[y][x])
}
