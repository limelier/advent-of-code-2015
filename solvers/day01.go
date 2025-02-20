package solvers

import (
	"log"
)

//goland:noinspection GoUnusedExportedFunction
func Solve01(input string) (int, int) {
	floor := 0
	var firstBasementEntrance int
	for pos, char := range input {
		switch char {
		case '(':
			floor++
		case ')':
			{
				floor--
				if floor < 0 && firstBasementEntrance == 0 {
					firstBasementEntrance = pos + 1
				}
			}
		default:
			log.Fatalf("Input is invalid: found character %s in input string.", char)
		}
	}

	return floor, firstBasementEntrance
}
