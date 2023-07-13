package merkletree

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

// MerkleTree holds the root node of a Merkle tree and a hash function.
type MerkleTree struct {
	Root         *Node
	HashFunction func([]byte) []byte
}

// Node represents a Merkle tree node.
type Node struct {
	Left  *Node
	Right *Node
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data.
func NewMerkleTree(data [][]byte, hashDetails ...func([]byte) []byte) *MerkleTree {
	// Default hash function is SHA256.
	hashFunction := Hash
	merkleTree := MerkleTree{HashFunction: hashFunction}

	// If a hash function is provided, use it instead of the default.
	if len(hashDetails) != 0 {
		merkleTree.HashFunction = hashDetails[0]
	}
	hashData := [][]byte{}
	for _, datum := range data {
		hashData = append(hashData, merkleTree.HashFunction(datum))
	}
	merkleTree.Root = merkleTree.buildTree(hashData)
	return &merkleTree
}

// buildTree builds a Merkle tree from a sequence of data.
func (merkleTree *MerkleTree) buildTree(data [][]byte) *Node {
	// Create leaf nodes.
	var nodes []Node
	for _, datum := range data {
		node := Node{Data: datum}
		nodes = append(nodes, node)
	}
	// Create parent nodes by combining leaf nodes.
	for len(nodes) > 1 {

		// Combine two nodes.
		var level []Node
		for j := 0; j < len(nodes); j += 2 {
			if (j + 1) == len(nodes) {
				level = append(level, nodes[j])
				break
			}
			node := Node{Left: &nodes[j], Right: &nodes[j+1]}
			hash := merkleTree.HashFunction(append(node.Left.Data, node.Right.Data...))
			node.Data = hash
			level = append(level, node)
		}
		nodes = level
	}

	// Return the root node.
	return &nodes[0]
}

func (merkleTree *MerkleTree) GetRootHash() []byte {
	return merkleTree.Root.Data
}

// GetMerklePath returns the merkle path for a given data
func (merkleTree *MerkleTree) GetMerklePath(data []byte) [][]byte {
	dataHash := merkleTree.HashFunction(data)
	var merklePath [][]byte
	merklePath = merkleTree.Root.getHashPath(dataHash, merklePath)
	return merklePath
}

// getHashPath returns the merkle path for a given data by traversing the tree recursively
func (node *Node) getHashPath(data []byte, merklePath [][]byte) [][]byte {
	if bytes.Equal(node.Data, data) {
		return merklePath
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	leftPath := node.Left.getHashPath(data, append([][]byte{node.Right.Data}, merklePath...))
	if leftPath != nil {
		return leftPath
	}
	rightPath := node.Right.getHashPath(data, append([][]byte{node.Left.Data}, merklePath...))
	if rightPath != nil {
		return rightPath
	}
	return nil
}

// VerifyMerklePath verifies the merkle path for a given data
func VerifyMerklePath(rootHash []byte, data []byte, merklePath [][]byte) bool {
	hash := Hash(data)
	for _, path := range merklePath {
		hash = Hash(append(hash, path...))
	}
	return bytes.Equal(hash, rootHash)
}

// UpdateMerkleTree updates the merkle tree with new data
func (merkleTree *MerkleTree) UpdateMerkleTree(newData []byte, oldData []byte) bool {
	newHash := merkleTree.HashFunction(newData)
	oldHash := merkleTree.HashFunction(oldData)
	return merkleTree.Root.updateNode(newHash, oldHash, merkleTree.HashFunction)
}

// updateNode updates the nodes of a Merkle tree recursively.
func (node *Node) updateNode(newHash []byte, oldHash []byte, hashFunc func([]byte) []byte) bool {
	if bytes.Equal(node.Data, oldHash) {
		node.Data = newHash
		return true
	}
	if node.Left == nil && node.Right == nil {
		return false
	}
	if node.Left.updateNode(newHash, oldHash, hashFunc) || node.Right.updateNode(newHash, oldHash, hashFunc) {
		node.Data = hashFunc(append(node.Left.Data, node.Right.Data...))
		return true
	}
	return false
}

// AddData adds new data to the Merkle tree and updates the root node. Data is added to the rightmost leaf node.
func (merkleTree *MerkleTree) AddData(data []byte) {
	dataHash := merkleTree.HashFunction(data)
	leafNodeData := merkleTree.GetLeafNodeData()
	leafNodeData = append(leafNodeData, dataHash)
	merkleTree.Root = merkleTree.buildTree(leafNodeData)
}

// GetLeafNodeData returns the data of all leaf nodes of the Merkle tree.
func (merkletree *MerkleTree) GetLeafNodeData() [][]byte {
	var leafNodeData [][]byte
	merkletree.Root.getLeafNodeData(&leafNodeData)
	return leafNodeData
}

// getLeafNodeData returns the data of all leaf nodes of the Merkle tree recursively.
func (node *Node) getLeafNodeData(leafNodeData *[][]byte) {
	if node.Left == nil && node.Right == nil {
		*leafNodeData = append(*leafNodeData, node.Data)
		return
	}
	node.Left.getLeafNodeData(leafNodeData)
	node.Right.getLeafNodeData(leafNodeData)
}

// GetMaxDepth returns the maximum depth of the Merkle tree.
func (merkletree *MerkleTree) GetMaxDepth() int {
	return merkletree.Root.getMaxDepth()
}

// getMaxDepth returns the maximum depth of the Merkle tree recursively.
func (node *Node) getMaxDepth() int {
	if node.Left == nil && node.Right == nil {
		return 1
	}
	leftDepth := node.Left.getMaxDepth()
	rightDepth := node.Right.getMaxDepth()
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// String returns a string representation of a Merkle tree.
func (merkleTree *MerkleTree) String() string {
	return merkleTree.Root.String()
}

// String returns a string representation of a node.
func (node *Node) String() string {
	return node.stringHelper(1)
}

// stringHelper is a helper function for String().
func (node *Node) stringHelper(depth int) string {
	encodedString := hex.EncodeToString(node.Data)
	if node.Left == nil && node.Right == nil {
		return fmt.Sprintf("%s", encodedString)
	}
	return fmt.Sprintf("%s\n%s%s\n%s%s", encodedString, indent(depth), node.Left.stringHelper(depth+1), indent(depth), node.Right.stringHelper(depth+1))
}
