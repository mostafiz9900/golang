package main

import "fmt"

var address string

const IP = 23.0

func main() {
	/* 	fmt.Println("Hello World")
	   	var fName = "Md"
	   	var name string = "Mostafizur"
	   	age := 12
	   	var isStudent bool
	   	var lName string
	   	lName = "Rahman"
	   	address = "Dhaka"
	   	var a, b, c int = 1, 3, 4
	   	var aa, bb = 33, "Mirput"
	   	var (
	   		bangla string
	   		eng    int    = 990
	   		mkt    string = "Marketing"
	   	)
	   	var h, w string = "Hello", "world"
	   	fmt.Println(fName, " ", name, " ", lName)
	   	fmt.Println(name)
	   	fmt.Println(age)
	   	fmt.Println(isStudent)
	   	fmt.Println(lName)
	   	fmt.Println(address)
	   	fmt.Println(a + b + c)
	   	fmt.Println(aa)
	   	fmt.Println(bb)
	   	fmt.Println(bangla)
	   	fmt.Println(eng)
	   	fmt.Println(mkt)
	   	fmt.Println(IP)
	   	fmt.Print(h, "\n", w, "\n")
	   	fmt.Printf("%v: %T\n", name, name) */
	var student = [...]string{
		"Mostafiz", "Minhan",
	}
	arr1 := [5]int{1:10,2:40}

	
	student[0]="Showpon"
	fmt.Println(student)
	fmt.Println(arr1)
	fmt.Println(len(arr1))

}
