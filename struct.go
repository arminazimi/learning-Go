package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := person{firstName: "x", lastName: "y", age: 10}

	var z person
	z.age = 21
	z.firstName = "w"
	z.lastName = "q"

	fmt.Println(p, z)

}
