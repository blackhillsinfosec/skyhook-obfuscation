package obfuscate

import (
	"testing"
)

func TestXor(t *testing.T) {
	key := "secret"
	xor := XOR{Key: key}
	testAlgo(&xor, t)
}
