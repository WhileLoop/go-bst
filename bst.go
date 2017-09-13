package main

import (
  "fmt"
)

type Node struct {
	Value  int
	Left  *Node
	Right *Node
}

type Tree struct {
  Root *Node
}

func (t *Tree) Insert(value int) {
  // First insert, set root.
  if t.Root == nil {
    t.Root = &Node{Value: value}
  } else {
    t.Root.Insert(value)
  }
}

func (t *Tree) Print() {
  t.Root.PrintTree()
}

func (t *Tree) Find(value int) bool {
  return t.Root.Find(value)
}

func (n *Node) Insert(value int) {
  // Duplicates are no-op.
  switch {
  case value < n.Value:
    // At left leaf, create new child.
    if n.Left == nil {
      n.Left = &Node{Value: value}
    } else {
      // Recurse left subtree.
      n.Left.Insert(value)
    }
  case value > n.Value:
    // At right leaf, create new child.
    if n.Right == nil {
      n.Right = &Node{Value: value}
    } else {
      // Recurse right subtree.
      n.Right.Insert(value)
    }
  }
}

// Print the tree in order
func (n *Node) PrintTree() {
  if n.Left != nil {
    n.Left.PrintTree()
  }
  fmt.Println(n.Value)
  if n.Right != nil {
    n.Right.PrintTree()
  }
}

func (n *Node) Find(value int) bool {
  if n.Value > value {
    if n.Left == nil {
      return false
    } else {
      return n.Left.Find(value)
    }
  } else if n.Value < value {
    if n.Right == nil {
      return false
    } else {
      return n.Right.Find(value)
    }
  }
  return true
}

func (n *Node) Height() int {
  var left, right int = 0, 0

  if n.Left != nil {
    left = n.Left.Height()
  }

  if n.Right != nil {
    right = n.Right.Height()
  }

  return 1 + maxInt(left, right)
}

func maxInt(a int, b int) int {
  if a >= b {
    return a
  }
  return b
}

func main() {
  t := Tree{}
  t.Insert(7)
  t.Insert(1)
  t.Insert(14)
  t.Insert(4)
  t.Insert(9)
  t.Insert(2)
  t.Insert(0)
  t.Print()
  fmt.Println(t.Find(4))
  fmt.Println(t.Root.Height())
}
