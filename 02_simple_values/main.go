package main

import (
    "fmt"
    "math"
)

func main(){
	fmt.Println(1)
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Pow(2, 4))
	// string 
	fmt.Println("Hello, " + "World!")
    // bool+
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
    // integer
    fmt.Println(1 + 2)
    fmt.Println(1 - 2)
    fmt.Println(1 * 2)
    fmt.Println(1 / 2)
    fmt.Println(1 % 2)
    // float
    fmt.Println(1.0 + 2.0)
    fmt.Println(1.0 - 2.0)
    fmt.Println(1.0 * 2.0)
    fmt.Println(1.0 / 2.0)
    // complex
    fmt.Println(1 + 2i)
}