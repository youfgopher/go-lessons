package piscine

/*
	Write a function that prints all possible combinations of n different digits in ascending order.
	n will be defined as : 0 < n < 10
	Below are the references for the printing format expected.
	(for n = 1) '0, 1, 2, 3, ..., 8, 9'
	(for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'
*/

// PrintCombN prints all possible combinations of n different digits in ascending order.
func PrintCombN(k int) {
	const base = 10
	var sres []string

	res := combine(base, k)

	for _, arr := range res {
		sres = append(sres, aggregator(arr))
	}

	writer(sres)
}

func writer(data []string) {
	var sres string
	newLine := "\n"
	suffix := 2

	for _, s := range data {
		sres += s + " "
	}

	sres = sres[:len(sres)-suffix]
	sres += newLine

	print(sres)
}

func aggregator(arr []int) (bufString string) {
	length := len(arr)
	delim := ","
	for i := 0; i < length; i++ {
		switch arr[i] {
		case 0:
			bufString += "0"
		case 1:
			bufString += "1"
		case 2:
			bufString += "2"
		case 3:
			bufString += "3"
		case 4:
			bufString += "4"
		case 5:
			bufString += "5"
		case 6:
			bufString += "6"
		case 7:
			bufString += "7"
		case 8:
			bufString += "8"
		case 9:
			bufString += "9"
		case 10:
			bufString += "10"

		}
	}
	bufString += delim

	return bufString
}

func combine(n int, k int) [][]int {
	return combineRecursive(1, n, k)
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
