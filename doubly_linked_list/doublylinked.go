package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	start *Node
	tail  *Node
}

func main() {
	numOfItems := 10
	list := &List{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numOfItems; i++ {
		insert(list, rand.Intn(40))
	}
	insert(list, 666)
	printList(list)
	printListReversed(list)
	list2 := &List{}
	printList(list2)
	printListReversed(list2)
}

func insert(list *List, data int) {
	newNode := &Node{data, nil, nil}
	if list.start == nil {
		list.start = newNode
		list.tail = newNode
	} else {
		current := list.start
		for current.next != nil {
			current = current.next
		}
		newNode.prev = current
		current.next = newNode
		list.tail = newNode
	}
}

func printList(list *List) {
	top := list.start
	if top == nil {
		fmt.Println("The list is empty.")
		return
	}
	for top != nil {
		fmt.Print(top.data, "\t")
		top = top.next
	}
	fmt.Println()
}

func printListReversed(list *List) {
	tail := list.tail
	if tail == nil {
		fmt.Println("The list is empty.")
		return
	}
	for tail != nil {
		fmt.Print(tail.data, "\t")
		tail = tail.prev
	}
	fmt.Println()
}
