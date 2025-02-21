package main

import (
	"advent-of-code-2015/common"
	"advent-of-code-2015/solvers"
)

func main() {
	input := common.InputLines()
	p1, p2 := solvers.Solve05(input)
	common.Output(p1, p2)
}
