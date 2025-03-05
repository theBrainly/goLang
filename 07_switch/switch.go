package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hello, World!")
    // fwt.PrintMessage()
    // fmt.Println(fwt.Add(2, 3))
    // fmt.Println(fwt.Subtract(5, 3))
    // fmt.Println(fwt.Multiply(2, 3))
    // fmt.Println(fwt.Divide(6, 3))
    // fmt.Println(fwt.Exponent(2, 3))
    // fmt.Println(fwt.Factorial(5))
    // fmt.Println(fwt.IsPrime(7))
    // fmt.Println(fwt.Fibonacci(10))
    // fmt.Println(fwt.ReverseString("Hello, World!"))
    // fmt.Println(fwt.PalindromeCheck("madam"))
	// switch
	i:=3
	switch i{
		case 1:
            fmt.Println("one")
        case 2:
            fmt.Println("two")
        case 3:
            fmt.Println("three")
        default:
            fmt.Println("other")
		}
	// multiple condition switch
    switch time.Now().Weekday(){
	case time.Sunday,time.Saturday:
		fmt.Println("weekend")
    default:
		fmt.Println("working day")
	}
	// type switch
	 whoAmi:= func(i interface{}){
		switch v:=i.(type){
			case int:
                fmt.Println("int")
            case string:
                fmt.Println("string")
            default:
                fmt.Println("unknown ",v)
		}

	 }
	 whoAmi("akash")

}