package linkedlist

import "fmt"

//Linkedlist 链表
type Linkedlist struct {
	value int
	next  *Linkedlist
}

//ShowNode 打印节点
func ShowNode(head *Linkedlist) {
	for head != nil {
		fmt.Printf("%d\t", head.value)
		head = head.next
	}
}

//ReverseLinkedList 反转链表
func ReverseLinkedList(head *Linkedlist) *Linkedlist {
	var reverseNode *Linkedlist
	var node = head
	for node != nil {
		next := node.next
		node.next = reverseNode
		reverseNode = node
		node = next
	}
	return reverseNode
}
