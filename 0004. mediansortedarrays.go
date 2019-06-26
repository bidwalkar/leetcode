/*
Problem numbr 4 : Classified as hard problem
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	//"sort"
	"lib"
)
//First, let's do brute force and readymade functions to sort
//1. Merge arrays
//2. Find Mean
//This still runs between 8ms and 16ms and I am between top 1.5% and top 17%. Definitely room to optimize. But at least, all test cases pass and
//I know the functionality
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//n1 := len(nums1)
	//n2 := len(nums2)
	//n := n1 + n2
	//for _,val := range(nums2) {
		//nums1 = append(nums1,val)
	//}
	//sort.Ints(nums1)

	//This is the approach using my own merge/sort. I once came down to 8ms and top 1.5%
	a := make([][]int,2)
	a[0] = nums1
	a[1] = nums2
	nums1 = lib.MergeSortedArrays(a)
	n := len(nums1)
	rc := 0.0
	if n % 2 == 1 { //Odd
		mn := (n+1)/2 - 1
		rc = float64(nums1[mn])
	} else { //Even
		mn := n/2 - 1
		rc = float64(nums1[mn]+nums1[mn+1]) / 2
	}
	return rc
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	par1 := strings.Split(os.Args[1],",")
	num1 := make([]int,0)
	for _,val := range(par1) {
		vv,_ := strconv.Atoi(val)
		num1 = append( num1, vv)
	}
	par2 := strings.Split(os.Args[2],",")
	num2 := make([]int,0)
	for _,val := range(par2) {
		vv,_ := strconv.Atoi(val)
		num2 = append( num2, vv)
	}
	fmt.Printf("%d AND %d %f\n", num1, num2, findMedianSortedArrays(num1,num2))
}
