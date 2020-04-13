/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var ret []int
	/*
		for key, value := range nums {
			fmt.Println(key, " => ", value)
		}
	*/
	arrLen := len(nums)
	indexMin := 0
	indexMax := arrLen - 1

	for indexMin < indexMax {
		if nums[indexMin]+nums[indexMax] == target {
			ret = append(ret, indexMin)
			ret = append(ret, indexMax)
			break
		} else if nums[indexMin]+nums[indexMax] < target {
			indexMin++
		} else {
			indexMax--
		}
	}
	return ret
}

func bestTwoSum(nums []int, target int) []int {
	var ret []int
	// 1、使用make初始化一个map
	hashMap := make(map[int]int)
	// hashMap := map[int]int{}
	for key, value := range nums {
		hashMap[value] = key
	}
	// fmt.Println("hashMap = ", hashMap)
	for key, value := range nums {
		newKey, ok := hashMap[target-value]
		if ok == true && newKey != key {
			return []int{key, newKey}
		}
	}
	return ret
}

func main() {
	fmt.Println("LeetCode - two - sum -- main")
	allNums := []int{3, 2, 4, 15}
	target := 6
	fmt.Println(bestTwoSum(allNums, target))

}
