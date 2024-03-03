package main

import "fmt"

func main() {
	nums := []int{1,2,3,4}
	target := 4
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for index, value := range nums {
        difference := target - value
        if _, ok := seen[difference]; ok {
            return []int{indexOf(difference, nums), index}
        } 
        seen[value] = index
    }
    return []int{0,0}
}

func indexOf(target int, nums []int) (int) {
   for index, value := range nums {
       if target == value {
           return index
       }
   }
   return -1
}