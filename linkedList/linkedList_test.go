package linkedlist

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	node := &Linkedlist{
		1,
		&Linkedlist{
			2,
			&Linkedlist{
				3,
				&Linkedlist{
					4,
					&Linkedlist{
						5,
						nil,
					},
				},
			},
		},
	}
	ShowNode(node)
	fmt.Printf("%p\r\n", node)
	node = ReverseLinkedList(node)
	ShowNode(node)
	fmt.Printf("%p\r\n", node)
}
