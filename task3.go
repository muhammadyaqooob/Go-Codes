package main

import (
	"fmt"
	"strings"
)

type Student struct {
	name       string
	rollNumber int
	address    string
}

func newStudent(rollno int, name string, address string) *Student {

	s := new(Student)
	s.rollNumber = rollno
	s.name = name
	s.address = address
	return s
}

type studentList struct {
	list []*Student
}

func (ls *studentList) creatStudent(rollno int, name string, address string) *Student {
	st := newStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(studentList)
	student.creatStudent(24, "asim", "AAAAA")
	student.creatStudent(54, "Naveed", "BBBBBBBBB")
	fmt.Println(student.list[0].name)

	no:=1

	fmt.Printf("%s List %d %s \n", strings.Repeat("=",25),no,strings.Repeat("=",25))
	fmt.Printf("\nStudent RollNo	: %v\n",student.list[0].rollNumber);
	fmt.Printf("Student Name	: %v\n",student.list[0].name);
	fmt.Printf("Student Adress	: %v\n",student.list[0].address);
	no=no+1
	fmt.Printf("%s List %d %s \n", strings.Repeat("=",25),no,strings.Repeat("=",25))
	fmt.Printf("\nStudent RollNo	: %v\n",student.list[1].rollNumber);
	fmt.Printf("Student Name	: %v\n",student.list[1].name);
	fmt.Printf("Student Adress	: %v\n",student.list[1].address);
}