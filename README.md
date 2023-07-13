# Merkle Tree


This is an implementation of the Merkle Tree in Golang. It builds a Merkle tree from a list of data, generate Merkle proofs (Path) for leaf nodes, verifies Merkle proofs against the root hash, and updates the data of leaf nodes. 
The hash function for the merkle tree has been abstracted. This allows custom implementations of the hash functions to be used in the merkle tree. 
It uses 128 bit truncated SHA-256 hash algorithm as default. 

## Basic Usage

The Module can be used to add merkle trees to your application. 
 
The follwoing is a basic example

```
package main

import (

    "fmt"
    "github.com/kaushalingle/merkletree"
)

func main() {

    // Data for merkle trees
    data := [][]byte{[]byte("hello"), []byte("world"), []byte("foo"), []byte("bar")}

    testTree := merkleTree.NewMerkleTree(data)

    //to View the Tree
    fmt.Print(testTree)

    //to generate Path to leaf node
    testPath := testTree.GetMerklePath([]byte("hello"))

    //to verify the merkle path
    rootHash := testTree.GetRootHash()
    verified := merkleTree.VerifyMerklePath(rootHash, []byte("hello"), testPath)

    //to update a data
    testTree.UpdateMerkleTree([]byte("hello"), []byte("hi"))

    //to add new data at the rightmost node
    testTree.AddData([]byte("bye"))

}

```


## Issues

If you discover any issues, please email kaushalingle@gmail.com.


## Dependencies

This project depends on the crypto/sha256 package in Go to compute SHA256 hashes.

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
