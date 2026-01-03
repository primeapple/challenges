package csvToList

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func CsvToListAI(csvStr string) string {
	if csvStr == "" {
		return ""
	}

	reader := csv.NewReader(strings.NewReader(csvStr))
	records, err := reader.ReadAll()
	if err != nil {
		return ""
	}

	if len(records) <= 1 {
		return ""
	}

	header := records[0]
	nameIdx, ageIdx, cityIdx := -1, -1, -1
	for i, col := range header {
		col = strings.TrimSpace(col)
		switch col {
		case "name":
			nameIdx = i
		case "age":
			ageIdx = i
		case "city":
			cityIdx = i
		}
	}

	if nameIdx == -1 || ageIdx == -1 || cityIdx == -1 {
		return ""
	}

	var result []string
	for _, row := range records[1:] {
		if len(row) <= nameIdx || len(row) <= ageIdx || len(row) <= cityIdx {
			continue
		}
		name := row[nameIdx]
		ageStr := row[ageIdx]
		city := row[cityIdx]
		line := fmt.Sprintf("- %s, age %s, from %s", name, ageStr, city)
		result = append(result, line)
	}

	return "\n" + strings.Join(result, "\n") + "\n"
}
