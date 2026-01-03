package csvToList

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CsvToList(csv string) string {
	lines := strings.Split(csv, "\n")
	if len(lines) <= 1 {
		return ""
	}

	posName, posAge, posCity := 0, 1, 2
	persons := []Person{}
	for idx, line := range lines {
		if idx == 0 {
			posName = slices.Index(strings.Split(line, ","), "name")
			posAge = slices.Index(strings.Split(line, ","), "age")
			posCity = slices.Index(strings.Split(line, ","), "city")
			if posName == -1 || posAge == -1 || posCity == -1 {
				panic("Couldn't find all fields in csv header")
			}
		} else {
			person := parseRow(line, posName, posAge, posCity)
			persons = append(persons, person)
		}
	}

	return buildResultStr(persons)
}

func parseRow(row string, posName, posAge, posCity int) Person {
	values := []string{}
	currentValue := ""
	isQuotes := false
	for _, char := range row {
		switch string(char) {
		case ",":
			if isQuotes {
				currentValue += string(char)
			} else {
				values = append(values, currentValue)
				currentValue = ""
			}
		case "\"":
			isQuotes = !isQuotes
		default:
			currentValue += string(char)
		}
	}
	values = append(values, currentValue)

	name, age, city := "", 0, ""
	for idx := range 3 {
		switch idx {
		case posName:
			name = values[idx]
		case posAge:
			ageInt, err := strconv.Atoi(values[idx])
			if err != nil {
				panic(fmt.Errorf("Couldn't convert age string to int: %q, %w", values[idx], err))
			}
			age = ageInt
		case posCity:
			city = values[idx]
		}
	}

	return Person{name, age, city}
}

func buildResultStr(persons []Person) string {
	var builder strings.Builder

	for _, person := range persons {
		builder.WriteString("\n- ")
		builder.WriteString(person.name)
		builder.WriteString(", age ")
		builder.WriteString(fmt.Sprint(person.age))
		builder.WriteString(", from ")
		builder.WriteString(person.city)
	}
	builder.WriteString("\n")

	return builder.String()
}
