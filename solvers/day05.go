package solvers

import "slices"

func hasThreeVowels(word string) bool {
	vowels := 0

	for _, char := range word {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels++
		}
		if vowels >= 3 {
			return true
		}
	}
	return false
}

func IsNice1(word string) bool {
	if !hasThreeVowels(word) {
		return false
	}

	containsDouble := false
	for i := 0; i < len(word)-1; i++ {
		first, second := word[i], word[i+1]
		if first == second {
			containsDouble = true
		}

		pair := string(first) + string(second)
		if slices.Contains([]string{"ab", "cd", "pq", "xy"}, pair) {
			return false
		}
	}

	return containsDouble
}

func hasPairTwice(word string) bool {
	for i := 0; i < len(word)-3; i++ {
		for j := i + 2; j < len(word)-1; j++ {
			if word[i] == word[j] && word[i+1] == word[j+1] {
				return true
			}
		}
	}
	return false
}

func hasPalindromeTriple(word string) bool {
	for i := 0; i < len(word)-2; i++ {
		if word[i] == word[i+2] {
			return true
		}
	}
	return false
}

func IsNice2(word string) bool {
	return hasPairTwice(word) && hasPalindromeTriple(word)
}

//goland:noinspection GoUnusedExportedFunction
func Solve05(input []string) (int, int) {
	var nice1, nice2 int

	for _, word := range input {
		if IsNice1(word) {
			nice1++
		}
		if IsNice2(word) {
			nice2++
		}
	}

	return nice1, nice2
}
