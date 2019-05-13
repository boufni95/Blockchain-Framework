package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

//HashSha256 : returns the hash of a slice of bytes
func HashSha256(tohash []byte) string {
	h := sha256.New()
	h.Write(tohash)
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
