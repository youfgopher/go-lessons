package piscine

import (
	"github.com/01-edu/z01"
)

/*
	Write a function that prints all possible combinations of n different digits in ascending order.
	n will be defined as : 0 < n < 10
	Below are the references for the printing format expected.
	(for n = 1) '0, 1, 2, 3, ..., 8, 9'
	(for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'
*/

// PrintCombN prints all possible combinations of n different digits in ascending order.
func PrintCombN(k int) {
	const base = 9
	var sres []rune

	res := combine(base, k)

	for _, arr := range res {
		sres = append(sres, aggregator(arr)...)
	}

	writer(sres)
}

func writer(data []rune) {
	sres := make([]rune, len(data))
	newLine := '\n'
	suffix := 2

	for _, s := range data {
		sres = append(sres, s)
	}

	sres = sres[:len(sres)-suffix]
	sres = append(sres, newLine)

	for _, r := range sres {
		z01.PrintRune(r)
	}
}

func aggregator(arr []int) (bufRune []rune) {
	length := len(arr)
	delim := ','
	space := ' '

	for i := 0; i < length; i++ {
		switch arr[i] {
		case 0:
			bufRune = append(bufRune, '\u0030')
		case 1:
			bufRune = append(bufRune, '\u0031')
		case 2:
			bufRune = append(bufRune, '\u0032')
		case 3:
			bufRune = append(bufRune, '\u0033')
		case 4:
			bufRune = append(bufRune, '\u0034')
		case 5:
			bufRune = append(bufRune, '\u0035')
		case 6:
			bufRune = append(bufRune, '\u0036')
		case 7:
			bufRune = append(bufRune, '\u0037')
		case 8:
			bufRune = append(bufRune, '\u0038')
		case 9:
			bufRune = append(bufRune, '\u0039')
		}
	}
	bufRune = append(bufRune, delim)
	bufRune = append(bufRune, space)

	return bufRune
}

func combine(n int, k int) [][]int {
	return combineRecursive(0, n, k)
}

func combineRecursive(min, max, k int) [][]int {
	var out [][]int

	if k == 1 {
		for i := min; i <= max; i++ {
			out = append(out, []int{i})
		}

		return out
	}

	for i := min; i < max; i++ {
		for _, line := range combineRecursive(i+1, max, k-1) {
			out = append(out, append([]int{i}, line...))
		}
	}

	return out
}
