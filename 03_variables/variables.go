package main

import (
    "fmt"
)

func main() {
	// var name string="akash"
	// infer

	var name="akash"
	var isAdult bool = false
	var age int =30
	// shorthand synatx
	myName:="akash sharma"


	var nameOnly string;
	nameOnly="akash"
	// var price float64 = 0.22
	// var price =99
	price:=990.22
	fmt.Println(name,isAdult,age,myName,nameOnly,price)
    fmt.Println(name,isAdult,age,myName,nameOnly)
}

