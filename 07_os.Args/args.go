package args

import (
	"fmt"
	"os"
)

func PrintProgramName() {
	programName := os.Args[0]
	argOne := os.Args[1]
	argTwo := os.Args[2]

	fmt.Printf("Program name: %s\n\n", programName)
	fmt.Println("argOne: ", argOne)
	fmt.Println("argTwo: ", argTwo)
}
