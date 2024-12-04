package main

import (
	"bytes"
	"fmt"
	"os"
)

func checkForXmas(b1, b2, b3, b4 byte) bool {
	return (b1 == 88 && b2 == 77 && b3 == 65 && b4 == 83) || (b1 == 83 && b2 == 65 && b3 == 77 && b4 == 88)
}

func XmasCount(lines [][]byte) int {
	numLines := len(lines)
	numCol := len(lines[0])

	count := 0
	// checking lines
	for _, line := range lines {
		for i := 0; i < numCol-3; i++ {
			if checkForXmas(line[i], line[i+1], line[i+2], line[i+3]) {
				count++
			}
		}
	}

	// checking columns
	for lineIndex := 0; lineIndex < numLines-3; lineIndex++ {
		for colIndex := 0; colIndex < numCol; colIndex++ {
			if checkForXmas(lines[lineIndex][colIndex], lines[lineIndex+1][colIndex], lines[lineIndex+2][colIndex], lines[lineIndex+3][colIndex]) {
				count++
			}
		}
	}

	// checking for diagonals from top left to bottom right
	for lineIndex := 0; lineIndex < numLines-3; lineIndex++ {
		for colIndex := 0; colIndex < numCol-3; colIndex++ {
			if checkForXmas(lines[lineIndex][colIndex], lines[lineIndex+1][colIndex+1], lines[lineIndex+2][colIndex+2], lines[lineIndex+3][colIndex+3]) {
				count++
			}
		}
	}

	// checking for diagonals from top right to bottom left
	for lineIndex := 0; lineIndex < numLines-3; lineIndex++ {
		for colIndex := 3; colIndex < numCol; colIndex++ {
			if checkForXmas(lines[lineIndex][colIndex], lines[lineIndex+1][colIndex-1], lines[lineIndex+2][colIndex-2], lines[lineIndex+3][colIndex-3]) {
				count++
			}
		}
	}

	return count
}

func checkForMas(threeByThree [][]byte) bool {
	if len(threeByThree) != 3 {
		panic(fmt.Sprintf("Not 3 lines in %v", threeByThree))
	}

	isMasFromTopLeft := checkForXmas(threeByThree[0][0], threeByThree[1][1], threeByThree[2][2], 88) || checkForXmas(88, threeByThree[0][0], threeByThree[1][1], threeByThree[2][2])
	isMasFromTopRight := checkForXmas(threeByThree[0][2], threeByThree[1][1], threeByThree[2][0], 88) || checkForXmas(88, threeByThree[0][2], threeByThree[1][1], threeByThree[2][0])
	return isMasFromTopLeft && isMasFromTopRight
}

func XMasCount(lines [][]byte) int {
	numLines := len(lines)
	numCol := len(lines[0])

	count := 0
	for lineIndex := 0; lineIndex < numLines-2; lineIndex++ {
		for colIndex := 0; colIndex < numCol-2; colIndex++ {
			threeByThree := [][]byte{
				lines[lineIndex][colIndex : colIndex+3],
				lines[lineIndex+1][colIndex : colIndex+3],
				lines[lineIndex+2][colIndex : colIndex+3],
			}
			if checkForMas(threeByThree) {
				count++
			}
		}
	}

	return count
}

func main() {
	body, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(body, []byte("\n"))
	// remove last empty line
	lines = lines[:len(lines)-1]

	solution1 := XmasCount(lines)
	solution2 := XMasCount(lines)
	fmt.Println("Solution puzzle 1:", solution1)
	fmt.Println("Solution puzzle 2:", solution2)
}
