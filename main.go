package main

import (
	"advent-of-code-2015/common"
	"advent-of-code-2015/solvers"
)

func main() {
	input := common.Input()
	p1, p2 := solvers.Solve03(input)
	common.Output(p1, p2)
}
