package structs

import "fmt"

func some() {

	type User struct {
		Name string
		Age  int
	}

	users := []User{
		{
			Name: "User 1",
			Age:  23,
		},
		{
			Name: "User 2",
			Age:  25,
		},
		{
			Name: "User 3",
			Age:  27,
		},
	}

	for i, user := range users {
		user.Name = "Another name"
		user.Age = 25

		fmt.Printf("User, index: %d name: %s age: %d", i, user.Name, user.Age)
	}
}
