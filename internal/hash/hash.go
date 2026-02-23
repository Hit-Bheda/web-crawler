package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashFilename(url string) string {
	hash := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hash[:])
}
