package obfuscate

import "testing"

func TestAES(t *testing.T) {
	a := AES{Key: "secret"}
	testAlgo(&a, t)
}
