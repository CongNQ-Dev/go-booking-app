package main

import "fmt"

type Student struct {
	Code  string
	Name  string
	Email string
}

func checkExistedStudent(student1 Student, student2 Student) {
	if student1 == student2 {
		fmt.Println("Existed Student")
	} else {
		fmt.Println("Not Existed Student")
	}
}
func main() {
	student := Student{
		Code:  "SE150087",
		Name:  "Cong",
		Email: "cong@gmail.com",
	}
	var anonymousStudent = struct {
		Code  string
		Name  string
		Email string
	}{
		Code:  "SE150087",
		Name:  "Cong",
		Email: "cong@gmail.com",
	}
	checkExistedStudent(student, anonymousStudent)

}
