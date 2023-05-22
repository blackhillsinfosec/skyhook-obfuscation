package obfuscate

import (
	"testing"
)

func TestBase64(t *testing.T) {
	b64 := Base64{Rounds: 4}
	testAlgo(&b64, t)
}
