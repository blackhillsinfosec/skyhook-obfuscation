package obfuscate

import "testing"

func TestBlowfish(t *testing.T) {

	key := "secret"
	salt := "salt"
	b := Blowfish{
		Key:  key,
		Salt: salt,
	}
	testAlgo(&b, t)

}
