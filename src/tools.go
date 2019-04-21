// Review Remark: Should be a separate module (folder or smth..)
// Review Remark: Specify what functionality is provided given the responsibility and clear name.

package gameserver

import (
	"encoding/binary"
	"math/rand"
)

// Review Remark: ByteToIntConverter.go
func intTo4Byte(b *[]byte, i int, rev bool) {
	binary.LittleEndian.PutUint32(*b, (uint32)(i))
	if rev == true {
		for i, j := 0, len(*b)-1; i < j; i, j = i+1, j-1 {
			(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
		}
	}

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Review Remark: ByteRandomizer.go
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
