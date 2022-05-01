package encrypt

import (
	"crypto/sha256"
	"fmt"
)

// EncPWD serve caller to encrypt password to irent format
func EncPWD(pwd string) string {
	return fmt.Sprintf("0x%x", sha256.Sum256([]byte(pwd)))
}
