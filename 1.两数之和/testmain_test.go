package main

import (
	"testing"
)

// 比较两个数组是否相等的方法, 缺少了排序算法等
func StringSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Test_bestTwoSum_1(t *testing.T) {
	allNums := []int{2, 7, 11, 15}
	target := 9
	ret := []int{0, 1}
	if i := bestTwoSum(allNums, target); false == StringSliceEqual(i, ret) {
		t.Error("两数之和cases不通过")
	} else {
		t.Log("第一个case通过")
	}
}

func Test_bestTwoSum_2(t *testing.T) {
	allNums := []int{2, 5, 11, 15}
	target := 13
	ret := []int{0, 2}
	if i := bestTwoSum(allNums, target); false == StringSliceEqual(i, ret) {
		t.Error("两数之和cases不通过")
	} else {
		t.Log("第一个case通过")
	}
}

func Test_bestTwoSum_3(t *testing.T) {
	allNums := []int{2, 5, 11, 15}
	target := 130
	var ret []int
	if i := bestTwoSum(allNums, target); false == StringSliceEqual(i, ret) {
		t.Error("两数之和cases不通过")
	} else {
		t.Log("第一个case通过")
	}
}

func Test_bestTwoSum_4(t *testing.T) {
	allNums := []int{3, 2, 4}
	target := 6
	ret := []int{1, 2}
	if i := bestTwoSum(allNums, target); false == StringSliceEqual(i, ret) {
		t.Error("两数之和cases不通过")
	} else {
		t.Log("第一个case通过")
	}
}
