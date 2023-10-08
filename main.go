package main

import (
	"fmt"
)

type Node struct {
  data any
	next *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (linkedList *LinkedList) pop_front() any {
  node := linkedList.head
  linkedList.head = node.next
  return node.data
}

func (linkedList *LinkedList) value_at(index int) any {
 node := linkedList.head
 i := 0
 for node.next != nil {
  if index == i {
    return node.data
  } else {
    node = node.next
    i+=1
  }
 }
 return 0
}

func (linkedList *LinkedList) push_back(value any) {
  newNode := &Node{data: value, next: nil}
  if linkedList.tail == nil && linkedList.head == nil {
    linkedList.tail = newNode
    linkedList.head = newNode
  } else {
    linkedList.tail.next = newNode
    linkedList.tail = newNode
  }
}


func (linkedList *LinkedList) push_front(value any) {
  newNode := &Node{data: value, next: linkedList.head}
  if linkedList.head == nil && linkedList.tail == nil {
		linkedList.head = newNode
    linkedList.tail = newNode
	} else {
    newNode.next = linkedList.head
		linkedList.head = newNode
	}
}


func (linkedlist *LinkedList) pop_back() {
  size := linkedlist.size() - 3
  head := linkedlist.head

  for size >= 0 {
    head = head.next
    size = size - 1
  }

}

func (linkedList *LinkedList) size() int {
	total_nodes := 1
	node := linkedList.head
	for node.next != nil {
		total_nodes++
		node = node.next
	}  
	return total_nodes 
}


func (linkedlist *LinkedList) empty() bool {
  node := linkedlist.head
  if (node.next != nil) {
    return false
  } else {
    return true
  }
}

func (linkedlist *LinkedList) search(target any) *any  {
	head := linkedlist.head
	var value any;
	for head != nil {
		if head.data == target {
			value = head.data
		}
		head = head.next
	}
	return &value
}

func (linkedList *LinkedList) front() any {
  return linkedList.head.data
}

func (linkedList *LinkedList) back() any {
  return linkedList.tail.data
}

func (linkedlist *LinkedList) insert(index int, value int) {
  count := 1
  node := linkedlist.head
  if index == 1 {
    node.data = value
  } else {
    
    for index >= count {
      node = node.next
      count = count + 1
      node.data = value
    }

  }
  
}

func (linkedlist *LinkedList) erase(index int) {
  
  forward_node := linkedlist.head
  prev_node := linkedlist.head
  
  prev_index := index - 1


  forward_count := 1
  for index >= forward_count {
    forward_node = forward_node.next
    forward_count+=1
  }
  
  prev_count := 1
  for prev_index > prev_count {
    prev_node = prev_node.next
    prev_count+=1
  }

  
  prev_node.next = forward_node

}

/*
func (linkedlist *LinkedList) reverse() {
  node := linkedlist.head 
  for node != nil {
    prev :=   
  }
}
*/
 



func main() {
  linkedList := LinkedList{}

  
  linkedList.push_back(2)
  linkedList.push_back(3)
  linkedList.push_back(4)
  linkedList.push_back(5)
  linkedList.push_back(6)
  linkedList.push_back("XD")
  linkedList.pop_back();
/*  
  size := linkedList.size()
  linkedList.pop_front()
*/

  //linkedList.pop_back()

  //linkedList.insert(2, 121)
  fmt.Println(linkedList.value_at(6))

  //fmt.Println(size)

}



