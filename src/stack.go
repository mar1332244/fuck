package main

type Node struct {
    next *Node
    value int
}

type Stack struct {
    top *Node
    Size int
}

func (s *Stack) Push(value int) {
    node := new(Node)
    node.value = value
    node.next = s.top
    s.top = node
    s.Size++
}

func (s *Stack) Pop() {
    if s.top == nil {
        return
    }
    s.top = s.top.next
    s.Size -= 1
}

func (s *Stack) Peek() int {
    if s.top == nil {
        return -1
    }
    return s.top.value
}
