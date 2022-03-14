package bst

import (
	"errors"
	"fmt"
)

var errKeyExists = errors.New("key is present in tree")

// ErrKeyExists used to determine if returned error is because key is already present
func ErrKeyExists() error {
	return errKeyExists
}

var errTreeEmpty = errors.New("tree is empty")

// ErrTreeEmpty used to determine if returned error is because key is already present
func ErrTreeEmpty() error {
	return errKeyExists
}

// node a single node
type node struct {
	key   int
	left  *node
	right *node
	// if tree needs to be traversed add parent node to increase performance
	parent *node

	depth int
}

func newNode(key int, depth int, parent *node) *node {
	n := new(node)
	n.key = key
	n.depth = depth
	n.parent = parent
	return n
}

func (n *node) String() string {
	return fmt.Sprintf("%v", n.key)
}

// Tree structure to encapsulate BST functionality
type Tree struct {
	root *node
}

func FromSlice(arr []int) *Tree {
	tree := NewTree()
	for _, v := range arr {
		if err := tree.Insert(v); err != nil {
			fmt.Printf("duplicate key: %v", v)
		}
	}
	return tree
}

// NewTree poduces a new empty tree
func NewTree() *Tree {
	return new(Tree)
}

// Empty returns true if tree has not been populated
func (t *Tree) Empty() bool {
	return t.root == nil
}

// DeepestNodes returns a list of deepest nodes and the depth of the nodes
func (t *Tree) DeepestNodes() ([]int, int, error) {
	if t.Empty() {
		return nil, 0, errTreeEmpty
	}

	currentLevel := []*node{t.root}
	// all children of current level
	children := []*node{}

	index := 0 // index into the queue
	for {
		current := currentLevel[index]
		if current.left != nil {
			children = append(children, current.left)
		}
		if current.right != nil {
			children = append(children, current.right)
		}
		index++
		if index >= len(currentLevel) { // finished evaluating current level
			if len(children) == 0 { // current level is the leaf node level
				break
			}

			// reset tracking variables
			index = 0
			currentLevel = children
			children = []*node{}
		}
	}
	depth := currentLevel[0].depth
	values := make([]int, 0)
	for _, v := range currentLevel {
		values = append(values, v.key)
	}
	return values, depth, nil
}

// Insert insert new key into BST if key already exists error is returned
func (t *Tree) Insert(key int) error {
	if t.Empty() {
		t.root = newNode(key, 0, nil)
		return nil
	}
	node := t.root
	for {
		var leftNode bool
		parent := node
		if key < node.key {
			node = node.left
			leftNode = true
		} else if key > node.key {
			node = node.right
			leftNode = false
		} else {
			// value is already in tree do nothing
			return errKeyExists
		}

		if node == nil {
			if leftNode {
				parent.left = newNode(key, parent.depth+1, parent)
				break
			} else {
				parent.right = newNode(key, parent.depth+1, parent)
				break
			}
		}
	}
	return nil
}

func (t *Tree) Find(key int) (*node, bool) {
	node := t.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {

			return node, true
		}
	}
	return nil, false
}
func (t *Tree) ToSlice() []int {
	if t.Empty() {
		return []int{}
	}
	slice := []int{}
	queue := []*node{t.root}
	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:] // remove first element
		slice = append(slice, front.key)
		if front.left != nil {
			queue = append(queue, front.left)
		}
		if front.right != nil {
			queue = append(queue, front.right)
		}
	}
	return slice
}

func (t *Tree) Delete(key int) {
	node, found := t.Find(key)
	if !found {
		return
	}

	if node.right != nil {
		child := node.right
		if node.parent.left == node {
			node.parent.left = child
		} else if node.parent.right == node {
			node.parent.right = child
		}

		// insert left children back into tree
		for {
			if child.left != nil {
				child = child.left
			} else {
				child.left = node.left
				return
			}
		}
	} else if node.left != nil {
		child := node.left
		if node.parent.right == node {
			node.parent.right = child
		} else if node.parent.left == node {
			node.parent.left = child
		}
	} else { // child is a leaf
		if node.parent.right == node {
			node.parent.right = nil
		} else {
			node.parent.left = nil
		}
	}
}
