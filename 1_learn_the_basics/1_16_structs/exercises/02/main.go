package main

import "fmt"

type Tree struct {
	Value int
	Left, Right *Tree
}

func add(t* Tree, value int) *Tree {
	if t == nil {
		tree := Tree{value, nil, nil}
		return &tree 
	}

	if t.Value > value {
		t.Left = add(t.Left, value)
	} else {	
		t.Right = add(t.Right, value)
	}

	return t
}

func show(t* Tree) {
	if t == nil {
		return
	}

	show(t.Left)
	fmt.Printf("%d ", t.Value)
	show(t.Right)
}

func main() {	
	
	var tree Tree
	var arr = []int{2,10,7,12,5,15}

	for _,v := range arr {
		tree = *add(&tree, v)	
	}

	show(&tree)
}

// [10, 7, 5, 0, 2]

// add(nil, 10) -> [nil,10,nil]
// add([nil,10,nil], 7) -> [add(nil,7),10,nil]
//.........................[[nil,7,nil],10,nil]
// add([[nil,7,nil],10,nil],5) -> [add([nil,7,nil],5)]
//................................[]