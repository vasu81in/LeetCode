/*
*
* Vasu Mahalingam, Nov 2017
*
 */

package main

import (
	"fmt"
)

// Small go code to do swapping of nodes in a linked list
// input:   1 ----> 2 ----> 3 ----> 4 ----> 5
// result:  2 ----> 1 ----> 4 ----> 3 ----> 5

type List struct {
	elem interface{}
	next *List
}

func (l *List) new_element(e interface{}) *List {
	return &List{
		elem: e,
		next: nil,
	}
}

func (l *List) append(elem interface{}) {
	if l.next == nil {
		l.next = l.new_element(elem)
		return
	}
	walk := l.next
	for {
		if walk.next == nil {
			break
		}
		walk = walk.next
	}
	walk.next = l.new_element(elem)
}

func (l *List) print() {
	walk := l
	for walk != nil {
		fmt.Println(walk.elem)
		walk = walk.next
	}
}

// This piece of logic doesnt seem elegant :>)

func (l *List) swap() {
	var s,t *List
	head := l.next
	// zero or one node case
	if head == nil || head.next == nil {
		return
	}

	prev := head
	curr := prev.next
	l.next = curr // Fix the head
	for {
		if s != nil {
			s.next = curr
		}
		t = curr.next
		curr.next = prev
		prev.next = t
		if t != nil {
			s = prev
			prev = t
			curr = prev.next
		} else {
			break
		}
		if curr == nil {
			break
		}
	}
}

func main() {
	l1 := List{}
	// Test 1 
	l1.append(1)
	l1.append(2)
	l1.swap()
	// 2 1
	l1.print()
	// Test 2 
	l2 := List{}
	l2.append(1)
	l2.append(2)
	l2.append(3)
	l2.swap()
	// 2 1 3
	l2.print()
	// Test 3 
	l3 := List{}
	l3.append(1)
	l3.append(2)
	l3.append(3)
	l3.append(4)
	l3.swap()
	// 2 1 4 3 
	l3.print()
	// Test 4 
	l4 := List{}
	l4.append(1)
	l4.append(2)
	l4.append(3)
	l4.append(4)
	l4.append(5)
	l4.swap()
	// 2 1 4 3 5
	l4.print()
	// Test 5 
	l5 := List{}
	l5.swap()
	// nil 
	l5.print()
	// Test 6 
	l6 := List{}
	l6.append(1)
	l6.swap()
	// 1 
	l6.print()
}

