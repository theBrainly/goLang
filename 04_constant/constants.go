package main

import  "fmt"

func main() {
	// :=
	// name := "akash"
    // fmt.Println(name)

    // name="new name" // this will throw an error as name is declared as const
    // fmt.Println(name)

    // const name string="akash"
    // name = "new name" // this will throw an error as name is declared as const
    // fmt.Println(name)

    const (
        a = 5
        b = 10
        c = a + b
    )
    fmt.Println(c)
	const (
		port =5500
		host = "localhost"
	)
 fmt.Printf("Port: %d, Host: %s\n", port, host)
    // name := "akash" // this will throw an error as name is declared as const and redeclared
    // fmt.Println(name)

    // const name string = "akash"
    // fmt.Println(name)

    // name = "new name" // this will throw an error as name is declared as const and redeclared
	const name string="cant change name"
    fmt.Println(name)
}
    