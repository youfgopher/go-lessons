package main

import (
	"fmt"

	"github.com/youfgopher/go-lessons/04_tasks"
)

func main() {
	userData := map[string]string{
		"David":  "9876543",
		"Daniil": "1234567",
	}
	name := " -a _l E x-An )Der "
	phone := " + 123 ( 234) _1 )2-1 _2-12 "

	err := tasks.AddUniqueUserToMap(userData, name, phone) // output: User, username: alexander, phone: 123234121212
	if err != nil {
		fmt.Println(err)
	}
}
