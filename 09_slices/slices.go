package main

import (
	"fmt"
	"slices"
)


func main() {
	// var nums []int
	// // nums = append(nums, 1, 2, 3)
	// fmt.Println(nums)
	// fmt.Println(nums==nil)
	// nums = make([]int, 0, 5)

	// nums = append(nums, 1, 2, 3)
    // fmt.Println(nums)
    // fmt.Println(nums==nil)

    // nums := make([]int, 3, 5)
    // nums[0] = 1
    // nums[1] = 2
	// nums = append(nums, 3)


	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))
	// nums:=make([]int,0,5)
	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums2 := make([]int,len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums2)
	
	// fmt.Println(nums)
	// fmt.Println(nums2)
	// copy(nums2, nums)
	// fmt.Println(nums)
	// fmt.Println(nums2)
	// 
	// 
	// 	// slice operator
	    // nums := []int{1, 2, 3, 4, 5}
        // fmt.Println(nums[:2])
        // fmt.Println(nums[2:])
        // fmt.Println(nums[1:3])
        // fmt.Println(nums[:3])
        // fmt.Println(nums[3:])
        // fmt.Println(nums[:])

        // // slice reassignment
        // nums := []int{1, 2, 3, 4, 5}
        // nums[1:3] = []int{7, 8}
        // fmt.Println(nums)
        // nums = append(nums[:1], nums[2:]...)
        // fmt.Println(nums)
		// 

		// slice
		num1:=[]int{1,2,3,4}
		num2:=[]int{1,2,3,4}
		num4:=[]int{1,2,3,4}
		num3:=append(append(num1,num2...),num4...)
		fmt.Println(num3)
		fmt.Println(slices.Equal(num1,num2))


}
