package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	data int
	next *Node
}

func main() {
	numOfItems := 10
	top := Node{}
	top.next = nil
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numOfItems; i++ {
		top = addAtTop(top, rand.Intn(40))
	}
	printList(top)
	addToEnd(top, 222)
	printList(top)
	addToPosition(top, 444, 2)
	printList(top)
	top = removeAtTop(top)
	printList(top)
	removeAtEnd(&top)
	printList(top)
	removeAtPosition(&top, 8)
	printList(top)
	bubbleSort(&top)
	printList(top)
}

func addAtTop(top Node, data int) Node {
	newCell := Node{data, &top}
	return newCell
}

func addToEnd(top Node, data int) {
	before := &top
	// find the last node
	for before.next.next != nil {
		before = before.next
	}
	newCell := Node{data, before.next}
	before.next = &newCell
}

func addToPosition(top Node, data, pos int) {
	len := getLen(top)
	if pos > len-1 {
		fmt.Println("Can't add to index ", pos, "... Length of the list is less than position (", len, ")")
		return
	}
	before := &top
	curr := 0
	for curr < pos-1 {
		before = before.next
		curr++
	}
	newCell := Node{data, before.next}
	before.next = &newCell
}

func getLen(top Node) int {
	current := top
	var len int = 0
	for current.next != nil {
		len++
		current = *current.next
	}
	return len
}

func removeAtTop(top Node) Node {
	return *top.next
}

func removeAtEnd(top *Node) {
	prev := top
	for prev.next.next != nil {
		prev = prev.next
	}
	prev.next = nil
}

func removeAtPosition(top *Node, pos int) {
	len := getLen(*top)
	if pos > len-1 {
		fmt.Println("Can't remove at index ", pos, "... Length of the list is less than position (", len, ")")
		return
	}
	before := top
	curr := 0
	for curr < pos-1 {
		before = before.next
		curr++
	}
	if before.next != nil {
		before.next = before.next.next
	} else {
		before.next = nil
	}
}

func printList(top Node) {
	current := &top
	for current.next != nil {
		fmt.Print(current.data, "\t")
		current = current.next
	}
	fmt.Println()
}

func bubbleSort(top *Node) {
	if top == nil {
		return
	}
	var curr = top
	top = nil
	for curr != nil && curr.next != nil {
		var temp = curr.next
		for temp.next != nil {
			if curr.data > temp.data {
				curr.data, temp.data = temp.data, curr.data
			}
			temp = temp.next
		}
		curr = curr.next
	}
}
