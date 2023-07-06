package piscine

import (
	"fmt"
	"strings"
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
    const base = 10
	var sres []string

	res := combine(base, k)

	for _, arr := range res {
		sres = append(sres, aggregator(arr))
	}
    
	fmt.Printf("%v\n\n", trimArtefact(sres))
}

func aggregator(arr []int) string {
	length := len(arr)
	var bufString string 


	for i := 0; i < length; i++ {
		bufString += fmt.Sprintf("%v", arr[i])
	}

	bufString += ","

	return bufString
}

func trimArtefact(s []string) (res string) {
	res = strings.Join(s, " ")
	after, ok := strings.CutPrefix(res, "[")
	if ok {
		res = after
	}

	after, ok = strings.CutSuffix(res, ",]")
	if ok {
		res = after
	}

	after, ok = strings.CutSuffix(res, ",")
	if ok {
		res = after
	}

	return res
}


func combine(n int, k int) [][]int {
    results := [][]int{}
    if k > n {
        return results
    }
    
    dfs([]int{}, n, k, 0, &results)
    return results
}

func dfs(buf []int, n, k, idx int, results *[][]int) {
    if k == 0 {
        tmp := make([]int, len(buf))
        copy(tmp, buf)
        *results = append(*results, tmp)
    }
    
    for i := idx; i <= n; i++ {
        buf = append(buf, i)
        
        dfs(buf, n, k-1, i+1, results)
        
        buf = buf[:len(buf)-1]
    }
}