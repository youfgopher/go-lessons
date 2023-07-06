package main

import "fmt"

/*
## 03 Functions

**Topics:**

- **Declaration Function**
- **Functions and their parameters**
- **Returning a result from a function**
- **Function type**
- **Anonymous functions**
- **Recursive functions**
- **Closures**
- **defer and panic**
*/

func main() {

	mili := convertion(100)
	fmt.Println(mili)
	fmt.Println(convertion(100.0))
	// 62.14 | 0
	// +6.214000e+001 | 0
}

func convertion(calcMili float64) float64 {
	calcMili *= 0.6214
	return calcMili
}

// simple func

/*

1. Конвертер километров. Напишите программу, 
которая просит пользователя ввести рас­стояние в 
километрах и затем это расстояние преобразует в 
мили. Формула преобразова­ния: 

мили = километры х 0.6214.

*/

