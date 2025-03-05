package main

import "fmt"

func main() {
	var nums [5]int
	nums[0]=10
	nums[1]=20
	nums[2]=30
	fmt.Println(nums[0])
   fmt.Println(len(nums))
   fmt.Println(nums)
//    string array
   var names [10]string
   names[0]="akash"
   names[1]="rahul"
   names[2]="deepak"
   fmt.Println(names)
   fmt.Println(len(names))
//    bool array
   var booleans [5]bool
   booleans[0]=true
   booleans[1]=false
   booleans[2]=true
   fmt.Println(booleans[0])
   fmt.Println(len(booleans))
   // initialize array with values
   var numbers [8]int = [8]int{1,2,3,4,5}
   fmt.Println(numbers)
   // array slicing
   fmt.Println(numbers[1:3])
   // array concatenation
   fmt.Println(append(numbers[1:4], 6, 7, 8))

//    int fill with zero
// string fill with ""
// bool fill with false

// 2d array
   var twoDArray [3][3]int
   twoDArray[0][0]=1
   twoDArray[0][1]=2
   twoDArray[0][2]=3
   twoDArray[1][0]=4
   twoDArray[1][1]=5
   twoDArray[1][2]=6
   twoDArray[2][0]=7
   twoDArray[2][1]=8
   twoDArray[2][2]=9
   fmt.Println(twoDArray)

   arr:=[2][2]int{{2,1},{1,2}}
   fmt.Println(arr)



}