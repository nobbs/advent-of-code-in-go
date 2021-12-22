package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/nobbs/advent-of-code-in-go/util"
)

type Node struct {
	val    int
	level  int
	left   *Node
	right  *Node
	parent *Node
}

// parse the given string by counting brackets starting from
// left and splitting the string at the right ','.
func splitPair(s string) (string, string) {
	ctr, cut := 0, 0

	for i, c := range s {
		switch c {
		case '[':
			ctr++
		case ']':
			ctr--
		case ',':
			if ctr == 1 {
				cut = i
				break
			}
		}
	}

	return s[1:cut], s[cut+1 : len(s)-1]
}

func parseNode(s string, parent *Node) *Node {
	// split the node string into left and right element of the pair
	leftString, rightString := splitPair(s)

	// helper function
	parse := func(s string, p *Node) (node *Node) {
		if strings.Contains(s, "[") {
			// node with children
			node = parseNode(s, p)
		} else {
			// leaf node
			val := util.ParseInt(s)
			node = &Node{val: val, parent: p, level: p.level + 1}
		}
		return node
	}

	// create the node
	newNode := &Node{parent: parent}
	if parent != nil {
		newNode.level = parent.level + 1
	} else {
		newNode.level = 0
	}

	// recursively parse child nodes if there are any
	newNode.left = parse(leftString, newNode)
	newNode.right = parse(rightString, newNode)

	return newNode
}

// pre-order traversal with a custom function.
// mostly for fun
func (n *Node) preorderTraversal(fn func(*Node)) {
	fn(n)

	if n.left != nil {
		n.left.preorderTraversal(fn)
	}
	if n.right != nil {
		n.right.preorderTraversal(fn)
	}
}

// in-order traversal with a custom function.
// mostly for fun
func (n *Node) inorderTraversal(fn func(*Node)) {
	if n.left != nil {
		n.left.preorderTraversal(fn)
	}

	fn(n)

	if n.right != nil {
		n.right.preorderTraversal(fn)
	}
}

// post-order traversal with a custom function.
// mostly for fun
func (n *Node) postorderTraversal(fn func(*Node)) {
	if n.left != nil {
		n.left.preorderTraversal(fn)
	}
	if n.right != nil {
		n.right.preorderTraversal(fn)
	}

	fn(n)
}

// explode a node with two value-children
func (N *Node) explode() {
	// looking for left sibling: move up the tree until there's a left node that's
	// not the starting node
	leftSibling := N
	for leftSibling.parent != nil {
		old := leftSibling
		leftSibling = leftSibling.parent

		if leftSibling.left != old {
			leftSibling = leftSibling.left
			break
		}
	}

	// now move down to the right child until there's no more
	if leftSibling.parent != nil {
		for leftSibling.right != nil {
			leftSibling = leftSibling.right
		}
	} else {
		leftSibling = nil
	}

	// looking for right sibling: move up the tree until there's a right node that's
	// not the starting node
	rightSibling := N
	for rightSibling.parent != nil {
		old := rightSibling
		rightSibling = rightSibling.parent

		if rightSibling.right != old {
			rightSibling = rightSibling.right
			break
		}
	}

	// now move down to the left child until there's no more
	if rightSibling.parent != nil {
		for rightSibling.left != nil {
			rightSibling = rightSibling.left
		}
	} else {
		rightSibling = nil
	}

	// add the values to the siblings if they exist
	if leftSibling != nil {
		leftSibling.val += N.left.val
	}
	if rightSibling != nil {
		rightSibling.val += N.right.val
	}

	// clear the exploded node
	N.val = 0
	N.left = nil
	N.right = nil
}

// split a node into two more child nodes
func (n *Node) split() {
	if n.val != 0 {
		leftNode := &Node{
			val:    int(math.Floor(float64(n.val) / 2.0)),
			parent: n,
			level:  n.level + 1,
		}
		rightNode := &Node{
			val:    int(math.Ceil(float64(n.val) / 2.0)),
			parent: n,
			level:  n.level + 1,
		}

		n.val = 0
		n.left = leftNode
		n.right = rightNode
	}
}

func (n *Node) reduce() bool {
	// look for explodable pairs
	hadExplosion := false
	var traverseAndExplode func(*Node)
	traverseAndExplode = func(m *Node) {
		// check if current is explodable
		if !hadExplosion && m.level >= 4 {
			if (m.left != nil && m.left.left == nil && m.left.right == nil) && (m.right != nil && m.right.left == nil && m.right.right == nil) {
				m.explode()
				// fmt.Printf("explode: %v\n", n)
				hadExplosion = true
			}
		}

		if !hadExplosion && m.left != nil {
			traverseAndExplode(m.left)
		}

		if !hadExplosion && m.right != nil {
			traverseAndExplode(m.right)
		}
	}

	traverseAndExplode(n)
	if hadExplosion {
		return true
	}

	// second part, if no explosion occured: look for numbers greater than 10
	// to split

	hadSplit := false
	var traverseAndSplit func(*Node)
	traverseAndSplit = func(m *Node) {
		// check if current is explodable
		if !hadSplit && m.left == nil && m.right == nil && m.val >= 10 {
			m.split()
			// fmt.Printf("split:   %v\n", n)
			hadSplit = true
		}

		if !hadSplit && m.left != nil {
			traverseAndSplit(m.left)
		}

		if !hadSplit && m.right != nil {
			traverseAndSplit(m.right)
		}
	}

	traverseAndSplit(n)
	if hadSplit {
		return true
	}

	return false
}

func (n *Node) String() string {
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d", n.val)
	} else {
		return fmt.Sprintf("[%s,%s]", n.left, n.right)
	}
}

// recursive magnitute calculation
func (n *Node) magnitude() (mag int) {
	if n.left != nil {
		mag += 3 * n.left.magnitude()
	}

	if n.right != nil {
		mag += 2 * n.right.magnitude()
	}

	if n.val != 0 {
		mag += n.val
	}

	return mag
}

func partOne(lines []string) int {
	a := lines[0]
	var r *Node

	for _, b := range lines[1:] {
		number := "[" + a + "," + b + "]"
		r = parseNode(number, nil)
		for repeat := r.reduce(); repeat; {
			repeat = r.reduce()
		}
		a = fmt.Sprintf("%s", r)
	}

	return r.magnitude()
}

func partTwo(lines []string) int {
	maxMagnitude := 0
	for i, a := range lines {
		for j, b := range lines {
			if i != j {
				// a+b
				number := "[" + a + "," + b + "]"
				r := parseNode(number, nil)
				for repeat := r.reduce(); repeat; {
					repeat = r.reduce()
				}
				if mag := r.magnitude(); mag > maxMagnitude {
					maxMagnitude = mag
				}

				// b+a
				number = "[" + b + "," + a + "]"
				r = parseNode(number, nil)
				for repeat := r.reduce(); repeat; {
					repeat = r.reduce()
				}
				if mag := r.magnitude(); mag > maxMagnitude {
					maxMagnitude = mag
				}
			}
		}
	}
	return maxMagnitude
}

func main() {
	lines := util.ReadInputFile("./input.txt")

	solutionOne := partOne(lines)
	fmt.Println("Solution for part 1:", solutionOne)

	solutionTwo := partTwo(lines)
	fmt.Println("Solution for part 2:", solutionTwo)
}
