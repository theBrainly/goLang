package main
import "fmt"

func main() {
    fmt.Println("Hello, World!")
	// if else 
	age:=18
	if age >= 18{
        fmt.Println("You are an adult")
    }else{
        fmt.Println("You are not an adult")
    }
	// if elseif else
	if age >= 18 {
        fmt.Println("You are an adult")
    } else if age >= 13 {
        fmt.Println("You are a teenager")
    } else {
        fmt.Println("You are a child")
    }
  
    
	// switch case
	color := "red"
    switch color {
    case "red":
        fmt.Println("Color is red")
    case "green":
        fmt.Println("Color is green")
    default:
        fmt.Println("Color is neither red nor green")
    }
    // fallthrough
    color = "blue"
    switch color {
    case "red":
        fmt.Println("Color is red")
    case "green":
        fmt.Println("Color is green")
    case "blue":
        fmt.Println("Color is blue")
    default:
        fmt.Println("Color is neither red, green nor blue")
    }
   
    
}
