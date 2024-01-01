package merkle

import "crypto/sha256"

// Node represent a Merkle tree node
type Node struct {
	Left, Right *Node
	Data        []byte
}

// Tree represent a Merkle tree
type Tree struct {
	RootNode *Node
}

// NewNode creates a new Merkle tree node
func NewNode(left, right *Node, data []byte) *Node {
	node := Node{}

	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		node.Data = hash[:]
	} else {
		prevHash := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHash)
		node.Data = hash[:]
	}

	node.Left = left
	node.Right = right

	return &node
}

// NewTree creates a new Merkle tree from a sequence of data
func NewTree(data [][]byte) *Tree {
	var nodes []Node

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2; i++ {
		var newLev []Node

		for j := 0; j < len(nodes); j += 2 {
			node := NewNode(&nodes[j], &nodes[j+1], nil)
			newLev = append(newLev, *node)
		}

		nodes = newLev
	}

	return &Tree{&nodes[0]}
}
