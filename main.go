package main

import (
	"advent-of-code-2015/common"
	"advent-of-code-2015/day02"
)

func main() {
	input := common.InputLines()
	p1, p2 := day02.Solve(input)
	common.Output(p1, p2)
}
