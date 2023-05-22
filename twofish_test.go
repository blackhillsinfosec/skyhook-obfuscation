package obfuscate

import "testing"

func TestTwofish(t *testing.T) {
	tf := Twofish{Key: "secret"}
	testAlgo(&tf, t)
}
