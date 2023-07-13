package merkletree

import (
	"bytes"
	"testing"
)

func TestNewMerkleTree(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world")}
	merkleTree := NewMerkleTree(data)
	if merkleTree.Root == nil {
		t.Errorf("Merkle tree root is nil")
	}
	if merkleTree.Root.Data == nil {
		t.Errorf("Merkle tree root data is nil")
	}
	if merkleTree.Root.Left == nil {
		t.Errorf("Merkle tree root left is nil")
	}
	if merkleTree.Root.Left.Data == nil {
		t.Errorf("Merkle tree root left data is nil")
	}
	if merkleTree.Root.Right == nil {
		t.Errorf("Merkle tree root right is nil")
	}
	if merkleTree.Root.Right.Data == nil {
		t.Errorf("Merkle tree root right data is nil")
	}
	if !bytes.Equal(merkleTree.Root.Left.Data[:], Hash([]byte("hello")[:])) {
		t.Errorf("Merkle tree root left data is not correct")
	}
	if !bytes.Equal(merkleTree.Root.Right.Data[:], Hash([]byte("world")[:])) {
		t.Errorf("Merkle tree root right data is not correct")
	}
	if !bytes.Equal(merkleTree.Root.Data[:], Hash(append(Hash([]byte("hello")), Hash([]byte("world"))...))) {
		t.Errorf("Merkle tree root data is not correct")
	}
	t.Logf("Merkle tree root data was correct")
}

func TestGetMerklePath(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar")}
	merkleTree := NewMerkleTree(data)
	merklePath := merkleTree.GetMerklePath([]byte("hello"))
	if len(merklePath) != 2 {
		t.Errorf("Merkle path length is not 2")
	}
	if !bytes.Equal(merklePath[0][:], Hash([]byte("world"))) {
		t.Errorf("Merkle path is not correct")
	}
	if !bytes.Equal(merklePath[1][:], Hash(append(Hash([]byte("foo")[:]), Hash([]byte("bar")[:])...))) {
		t.Errorf("Merkle path is not correct")
	}
}

func TestVerifyMerklePath(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar")}
	merkleTree := NewMerkleTree(data)
	merklePath := merkleTree.GetMerklePath([]byte("hello"))
	rootHash := merkleTree.GetRootHash()
	if !VerifyMerklePath(rootHash, []byte("hello"), merklePath) {
		t.Errorf("Merkle path verification failed")
	}
}

func TestUpdateMerkleTree(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar")}
	merkleTree := NewMerkleTree(data)
	merkleTree.UpdateMerkleTree([]byte("hello"), []byte("world"))
	editedData := [][]byte{[]byte("hello"), []byte("hello"), []byte("foo"), []byte("bar")}
	editedMerkleTree := NewMerkleTree(editedData)
	if !bytes.Equal(merkleTree.Root.Data[:], editedMerkleTree.Root.Data[:]) {
		t.Errorf("Merkle tree update failed")
	}
}

func TestAddData(t *testing.T) {
	data := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar")}
	merkleTree := NewMerkleTree(data)
	merkleTree.AddData([]byte("hello"))
	editedData := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar"), []byte("hello")}
	editedMerkleTree := NewMerkleTree(editedData)
	if !bytes.Equal(merkleTree.Root.Data[:], editedMerkleTree.Root.Data[:]) {
		t.Errorf("Merkle tree addition failed")
	}
}
