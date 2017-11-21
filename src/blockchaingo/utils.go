package blockchainGo

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// ComputeHashSha256 create a hash
func ComputeHashSha256(bytes []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(bytes))
}

// PseudoUUID create pseudo uuid
func PseudoUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return ""
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// StringSet struct of map
type StringSet struct {
	set map[string]bool
}

// NewStringSet return a struct of map string bool
func NewStringSet() StringSet {
	return StringSet{make(map[string]bool)}
}

// Add stringset on map
func (set *StringSet) Add(str string) bool {
	_, found := set.set[str]
	set.set[str] = true
	return !found
}

// Keys return keys
func (set *StringSet) Keys() []string {
	var keys []string
	for k := range set.set {
		keys = append(keys, k)
	}
	return keys
}
