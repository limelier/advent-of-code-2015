package tests

import (
	"advent-of-code-2015/solvers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNice2(t *testing.T) {
	assert.True(t, solvers.IsNice2("qjhvhtzxzqqjkmpb"), "qjhvhtzxzqqjkmpb should be nice")
	assert.True(t, solvers.IsNice2("xxyxx"), "xxyxx should be nice")
	assert.False(t, solvers.IsNice2("uurcxstgmygtbstg"), "uurcxstgmygtbstg should be naughty")
	assert.False(t, solvers.IsNice2("ieodomkazucvgmuy"), "ieodomkazucvgmuy should be naughty")
}
