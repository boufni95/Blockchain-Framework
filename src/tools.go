package gameserver

import "encoding/binary"

func intTo4Byte(b *[]byte, i int, rev bool) {
	binary.LittleEndian.PutUint32(*b, (uint32)(i))
	if rev == true {
		for i, j := 0, len(*b)-1; i < j; i, j = i+1, j-1 {
			(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
		}
	}

}
