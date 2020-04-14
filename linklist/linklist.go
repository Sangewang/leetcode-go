package linklist

import (
	"fmt"
)

// 创建节点结构、类型 => 数据域 和 指针域
type Node struct {
	Data interface{}
	Next *Node
}

// 创建链表结构
type LList struct {
	Head   *Node
	Length int
}

// a 设计接口
type Method interface {
	Insert(i int, v interface{}) // 增
	Delete(i int)                // 删
	GetLength() int              // 获取长度
	Search(v interface{}) int    // 查
	isNull() bool                // 判断是否为空
}

// b 初始化函数
// 创建节点
func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}

// 创建空链表
func CreateList() *LList {
	return &LList{CreateNode(nil), 0}
}

// 基于链表结构体实现Method中的方法
// 在i处插入节点(前插 => 即插入到原来的第i个节点之前，成为现在的第i个节点)
func (list *LList) Insert(i int, v interface{}) {
	// 根据传入的参数v创建一个节点
	newNode := CreateNode(v)
	curNode := list.Head
	for count := 0; count <= i; count++ {
		if count == i-1 {
			newNode.Next = curNode.Next
			curNode.Next = newNode
			list.Length++
		}
		curNode = curNode.Next
	}
}

// 删除第i个节点
func (list *LList) Delete(i int) {
	curNode := list.Head
	for count := 0; count <= i-1; count++ {
		nextNode := curNode.Next
		if count == i-1 {
			curNode.Next = nextNode.Next
			list.Length--
		}
		curNode = curNode.Next
	}
}

// 返回链表的长度
func (list *LList) GetLength() int {
	return list.Length
}

// 查找v所在的位置
func (list *LList) Search(v interface{}) int {
	curNode := list.Head.Next
	for i := 1; i <= list.Length; i++ {
		if curNode.Data == v {
			return i
		}
		curNode = curNode.Next
	}
	return 0
}

// 判空
func (list *LList) isNull() bool {
	curNode := list.Head.Next
	if curNode == nil {
		return true
	}
	return false
}

// d 设计打印链表的输出函数
func PrintList(list *LList) {
	curNode := list.Head.Next
	fmt.Println("List Show as Follows: ...", list.Length)
	for i := 0; i < list.Length; i++ {
		fmt.Printf("%v\n", curNode.Data)
		curNode = curNode.Next
	}
}

// 想要被外部调用必须大写 类似于Public
func RunLinkList() {
	fmt.Println("practice link list")
	lList := CreateList()
	fmt.Println("lList is null = ", lList.isNull())
	var M Method
	M = lList // 接口类型的变量可以存储所有实现该接口的类型变量
	M.Insert(1, 3)
	M.Insert(2, 6)
	M.Insert(1, 5)

	PrintList(lList)
	fmt.Println("List length is: ", lList.Length)
	fmt.Println("元素6在位置：", M.Search(6))
	fmt.Println("元素100在位置：", M.Search(100))
	fmt.Println("List is null: ", lList.isNull())

	M.Delete(2)
	PrintList(lList)
	fmt.Println("List length is: ", lList.Length)
}
