package whatKindOfNumber

import "fmt"

type KindOfNumber int

const (
	Perfect KindOfNumber = iota
	Abundant
	Deficient
)

func (c KindOfNumber) String() string {
	switch c {
	case 0:
		return "Perfect"
	case 1:
		return "Abundant"
	case 2:
		return "Deficient"
	}
	panic(fmt.Sprintf("unknown number %d", c))
}
