/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"leetcode/linklist"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、链表的创建，传进来一个数组，自动生成链表，返回头结点
func CreateLinkList(nums []int) *ListNode {
	p1Head := &ListNode{-1, nil}
	tail := p1Head
	/*
	for i:=0; i <= len(nums); i ++ {
		currNode := &ListNode{nums[i], nil}
		tail.Next = currNode
		tail = currNode
	}
	*/
	for _, value := range nums{
		currNode := &ListNode{value, nil}
		tail.Next = currNode
		tail = currNode
	}
	return p1Head
}

// 2、链表的显示
func ShowLinkList(pHead *ListNode) {
	ret := ""
	arrow := " -> "
	pTempNode := pHead
	for pTempNode != nil {
		ret +=  strconv.Itoa(pTempNode.Val) + arrow
		pTempNode = pTempNode.Next
	}
	ret = strings.TrimRight(ret, arrow)
	fmt.Println(ret)
}

// 3、链表做和
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Temp := l1.Next
	l2Temp := l2.Next
	lRet := &ListNode{-1, nil}
	lLoop := lRet
	// 1、都为空的情况
	if l1Temp == nil && l2Temp == nil {
		return lRet
	}
	carry := 0
	// 2、都不为空的情况
	for l1Temp != nil && l2Temp != nil {
		temp_sum := (l1Temp.Val + l2Temp.Val + carry) % 10
		carry = (l1Temp.Val + l2Temp.Val + carry) / 10
		l1Temp = l1Temp.Next
		l2Temp = l2Temp.Next
		lRetTemp := &ListNode{temp_sum, nil}
		lLoop.Next = lRetTemp
		lLoop = lRetTemp
	}
	// 3、l1Temp不为空的情况
	for l1Temp != nil {
		temp_sum := (l1Temp.Val + carry) % 10
		carry = (l1Temp.Val + carry) / 10
		l1Temp = l1Temp.Next
		lRetTemp := &ListNode{temp_sum, nil}
		lLoop.Next = lRetTemp
		lLoop = lRetTemp
	}
	// 4、l2Temp不为空的情况
	for l2Temp != nil {
		temp_sum := (l2Temp.Val + carry) % 10
		carry = (l2Temp.Val + carry) / 10
		l2Temp = l2Temp.Next
		lRetTemp := &ListNode{temp_sum, nil}
		lLoop.Next = lRetTemp
		lLoop = lRetTemp
	}

	if carry != 0 {
		lRetTemp := &ListNode{carry, nil}
		lLoop.Next = lRetTemp
		lLoop = lRetTemp
	}

	return lRet.Next
}

// 学习如何调用其它的package
func useLinkList() {
	fmt.Println("learn how to use other package")
	// 调用基础的链表的方法
	// linklist.RunLinkList()
	// 构造初始的链表1
	l1List := linklist.CreateList()
	l1List.Insert(1, 2)
	l1List.Insert(2, 4)
	l1List.Insert(3, 3)
	linklist.PrintList(l1List)

	// 构造初始的链表2
	l2List := linklist.CreateList()
	l2List.Insert(1, 5)
	l2List.Insert(2, 6)
	l2List.Insert(3, 4)
	linklist.PrintList(l2List)
}

func main() {
	// 1、调用其它的函数
	// useLinkList()
	// 2、自己再实现一套链表pHead1，头结点设为 -1
	l1List := []int{2, 4, 3}
	// fmt.Println("l1List = ", l1List)
	p1Head := CreateLinkList(l1List)
	ShowLinkList(p1Head)

	// 3、pHead2的实现
	l2List := []int{5, 6, 4}
	// fmt.Println("l2List = ", l2List)
	p2Head := CreateLinkList(l2List)
	ShowLinkList(p2Head)

	// 4、链表做和
	lRet := addTwoNumbers(p1Head, p2Head)
	ShowLinkList(lRet)

}
