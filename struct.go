package main

import "fmt"

type Student struct {
	Code string
}
type Person struct {
	Name  string
	Email string
	Student
}

func checkExistedStudent(person1 Person, person2 Person) {
	if person1 == person2 {
		fmt.Println("Existed Student")
	} else {
		fmt.Println("Not Existed Student")
	}
}
func main() {
	person := Person{
		Name:    "Cong",
		Email:   "cong@gmail.com",
		Student: Student{Code: "SE150087"},
	}
	var anonymousPerson = struct {
		Name  string
		Email string
		Student
	}{
		Student: Student{Code: "SE150087"},
		Name:    "Cong",
		Email:   "cong@gmail.com",
	}
	checkExistedStudent(person, anonymousPerson)

}
