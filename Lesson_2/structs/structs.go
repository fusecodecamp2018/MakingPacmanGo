package main

import "fmt"

//lets say i want to save multiple pieces of information that are related in one place
//if youve done java, think POJO's
type myname struct {
	firstname string
	lastname string
	age int
}

func main()  {
	person := myname{firstname: "bob", lastname: "smith", age: 55}
	fmt.Println(person.firstname)
	fmt.Println(person.lastname)
	fmt.Println(person.age)

}