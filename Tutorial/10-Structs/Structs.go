package main

import "fmt"

type person struct {
	name  string
	email string
	phone int
}

type result struct {
	person
	Is_pass bool
}

func main() {
	person_1 := person{
		name:  "pradip",
		email: "pradip@gmail.cin",
		phone: 9876541230,
	}
	fmt.Println(person_1)

	result_1 := result{
		person:  person_1,
		Is_pass: true,
	}
	fmt.Println(result_1)
	result_2 := result{
		person: person{
			name:  "avin",
			email: "avin@gmail.com",
			phone: 987654123,
		},
		Is_pass: false,
	}
	fmt.Println(result_2)
}
