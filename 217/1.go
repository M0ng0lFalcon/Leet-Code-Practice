/*
 * @Author: F4lc0n
 * @Date: 2021-11-06 23:42:53
 * @Last Modified by: F4lc0n
 * @Last Modified time: 2021-11-06 23:45:47
 * @Leet Code: 217. Contains Duplicate
 */
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	for _, num := range nums {
		if mp[num] == true {
			return true
		} else {
			mp[num] = true
		}
	}
	return false
}

func main() {
	arr0 := []int{1, 2, 3, 1}
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(arr0))
	fmt.Println(containsDuplicate(arr1))
	fmt.Println(containsDuplicate(arr2))
}
