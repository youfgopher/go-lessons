package main

import (
	"os"

	"github.com/01-edu/z01"
)


/*
func printStr(s string) {
  for _, r := range s {
    z01.PrintRune(r)
  }
  z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
  if even(nbr) == 1 {
    return yes
  } else {
    return no
  }
}

func main() {
  if isEven(lengthOfArg) == 1 {
    printStr(EvenMsg)
  } else {
    printStr(OddMsg)
  }
}
Usage
$ go run . "not" "odd"
I have an even number of arguments
$ go run . "not even"
I have an odd number of arguments

*/


func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	return nbr % 2
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	lengthOfArg := len(os.Args) 

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"


	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}