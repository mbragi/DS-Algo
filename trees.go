package main

import (
	"fmt"
)

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func (n *Tree) Insert(value int) {
	if n.Left == nil {
		n.Left = &Tree{value, nil, nil}
	} else if n.Right == nil {
		n.Right = &Tree{value, nil, nil}
	} else {
		n.Left.Insert(value)
	}
}

func (n *Tree) BFS() {
	list := []*Tree{n}
	for len(list) > 0 {
		root := list[0]
		list = list[1:]
		fmt.Println(root.Data)
		if root.Left != nil {
			list = append(list, root.Left)
		}
		if root.Right != nil {
			list = append(list, root.Right)
		}
	}
}

func (n *Tree) DFS() {
	if n == nil {
		return
	}
	fmt.Println(n.Data)
	n.Left.DFS()
	n.Right.DFS()
}
