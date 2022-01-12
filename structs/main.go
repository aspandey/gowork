package main

import "fmt"

type employee struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email  string
	mobile uint64
}

func main() {
	e1 := employee{
		firstName: "Tom",
		lastName:  "Hardy",
		contact: contactInfo{
			email:  "tom@gmail.com",
			mobile: 9009900990,
		},
	}
	fmt.Println(e1)
	fmt.Printf("%+v", e1)
}
