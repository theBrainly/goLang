package main


import     "fmt"

// for loop =>only construct in go for looping

func main(){
	fmt.Println("hello world")
	// while loop =>
	// i:=1
	// for i<=4{
    //     fmt.Println(i)
    //     i++
    // }
	// infinte loop
	for{
		fmt.Println("hello world")
        break // to break the loop if condition is met
        // continue // to skip current iteration and continue with next iteration
	}
	// for loop with range
	for i:=0; i<4; i++{
		fmt.Println(i)
	}

}