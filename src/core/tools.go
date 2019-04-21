package core

import (
	"encoding/binary"
	"math/rand"
)

//IntTo4Byte : convert an in into a slice of bytes
func IntTo4Byte(b *[]byte, i int, rev bool) {
	binary.LittleEndian.PutUint32(*b, (uint32)(i))
	if rev == true {
		for i, j := 0, len(*b)-1; i < j; i, j = i+1, j-1 {
			(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
		}
	}

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
