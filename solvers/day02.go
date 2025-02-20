package solvers

import (
	"slices"
	"strconv"
	"strings"
)

type box struct {
	length, width, height int
}

func boxFromInput(line string) box {
	parts := strings.Split(line, "x")
	l, _ := strconv.Atoi(parts[0])
	w, _ := strconv.Atoi(parts[1])
	h, _ := strconv.Atoi(parts[2])

	return box{l, w, h}
}

func (b *box) wrappingPaperArea() int {
	sideAreas := []int{b.length * b.width, b.width * b.height, b.length * b.height}
	return 2*(sideAreas[0]+sideAreas[1]+sideAreas[2]) + slices.Min(sideAreas)
}

func (b *box) volume() int {
	return b.length * b.width * b.height
}

func (b *box) ribbonLength() int {
	dimensions := []int{b.length, b.width, b.height}
	smallestSidePerimeter := 2 * ((b.length + b.width + b.height) - slices.Max(dimensions))
	return smallestSidePerimeter + b.volume()
}

//goland:noinspection GoUnusedExportedFunction
func Solve02(input []string) (int, int) {
	totalPaperArea := 0
	totalRibbonLength := 0

	for _, line := range input {
		box := boxFromInput(line)
		totalPaperArea += box.wrappingPaperArea()
		totalRibbonLength += box.ribbonLength()
	}

	return totalPaperArea, totalRibbonLength
}
