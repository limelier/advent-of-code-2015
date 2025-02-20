package solvers

import (
	. "advent-of-code-2015/common/int2"
	"log"
)

func Solve03(input string) (int, int) {
	visited1 := make(map[Pos]bool)
	visited2 := make(map[Pos]bool)
	var santa1, santa2, roboSanta Pos
	roboMoves := false
	visited1[santa1] = true
	visited2[santa2] = true

	for _, char := range input {
		var delta Delta
		switch char {
		case '>':
			delta = Delta{X: 1}
		case '<':
			delta = Delta{X: -1}
		case '^':
			delta = Delta{Y: -1}
		case 'v':
			delta = Delta{Y: 1}
		default:
			log.Fatalf("Unexpected char %b in input string", char)
		}

		santa1 = santa1.Moved(delta)
		visited1[santa1] = true

		if roboMoves {
			roboSanta = roboSanta.Moved(delta)
			visited2[roboSanta] = true
		} else {
			santa2 = santa2.Moved(delta)
			visited2[santa2] = true
		}
		roboMoves = !roboMoves
	}

	return len(visited1), len(visited2)
}
