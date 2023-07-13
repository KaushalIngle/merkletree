package merkletree

import (
	"bytes"
	"testing"
)

func TestHash(t *testing.T) {

	data := []byte("hello world")
	hash := Hash(data)
	if len(hash) != 16 {
		t.Errorf("Hash length is not 16 bytes")
	}
	if bytes.Equal(hash, []byte{223, 253, 96, 33, 187, 43, 213, 176, 175, 103, 98, 144, 128, 158, 195, 165}) {
		t.Errorf("Hash is not correct")
	}
	t.Logf("Hash was correct")
}
