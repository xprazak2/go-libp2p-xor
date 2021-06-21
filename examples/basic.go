package main

import (
	"fmt"
	"github.com/libp2p/go-libp2p-xor/key"
	"github.com/libp2p/go-libp2p-xor/trie"
	// "encoding/binary"
)

func main() {
	var keys []key.Key

	// for _, item := range [][]byte{[]byte{7}, []byte{2}, []byte{6}, []byte{4}, []byte{5}, []byte{0}, []byte{1}} {
	// 	keys = append(keys, key.BytesKey(item))
	// }

	// for _, item := range []{7, 2, 6, 4, 5, 0, 1} {
	// 	keys = append(keys, key.BytesKey(item))
	// }

	// for _, item := range []int{7, 6} {
	// 	b := make([]byte, 4)
	// 	binary.LittleEndian.PutUint32(b, uint32(item))
	// 	keys = append(keys, key.BytesKey(b))
	// }

	for _, item := range [][]byte{[]byte{7}, []byte{2}, []byte{6}, []byte{4}, []byte{5}, []byte{0}, []byte{1}} {
		keys = append(keys, key.BytesKey(item))
	}

	// fmt.Println(keys[0].BitString())

	tree := trie.FromKeys(keys)
	fmt.Print(tree)
}
