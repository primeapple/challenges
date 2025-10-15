package main

import "strings"

func GetCharAfterVimCommands(text string, commands string) string {
	lines := strings.Split(text, "\n")

	x := 0
	y := 0

	for _, rune := range commands {
		switch string(rune) {
		case "h":
			x = max(0, x-1)
		case "j":
			y = min(len(lines)-1, y+1)
			x = min(len(lines[y])-1, x)
		case "k":
			y = max(0, y-1)
			x = min(len(lines[y])-1, x)
		case "l":
			x = min(len(lines[y])-1, x+1)
		}
	}

	return string(lines[y][x])
}
