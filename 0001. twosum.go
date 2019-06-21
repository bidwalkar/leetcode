//Problem number 1 of Leetcode
/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main


//One pass hashmap lookup. Takes 4 ms. Faster than 98% submissions
func twoSum(nums[] int,target int) []int {
	numbers := make(map[int] int,0)
	for i:= 0 ; i < len(nums) ; i++ {
		complement := target-nums[i]
		val,ok := numbers[complement]
		if ok && (val != i) {
			rc := make([]int,2)
			rc[0] = val
			rc[1] = i
			return(rc)
		}
		numbers[ nums[i] ] = i
	}
	return nil
}
//The original brute force method. Works but takes 44 ms
func twoSum_bruteforce(nums[] int,target int) []int {
	var rc = make([]int,2)
	rc[0]=0
	rc[1]=0
	for i := 0 ; i < len(nums) ; i++ {
		for j := i+1 ; j < len(nums) ; j++ {
			if nums[i]+nums[j]==target {
				rc[0] = i
				rc[1] = j
				return(rc)
			}
		}
	}
	return( rc )
}

func main() {
	var par = make( []int, 4)
	//par = []int{2,7,11,15}
	//par = []int{3,2,4}
	//par = []int{3,3}
	par = []int{0,4,3,0}
	rc := twoSum( par, 0 )
	for i := range(rc) {
		print( rc[i],"\n" )
	}
}
