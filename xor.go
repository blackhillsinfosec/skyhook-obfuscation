package obfuscate

import (
	"errors"
)

// xor abstracts XOR cryption. Accepts input and key values, Returns output and an error value.
func xor(input []byte, key []byte) (output []byte, err error) {
	if len(key) == 0 {
		return input, errors.New("Key length of zero means no encryption")
	}

	output = []byte("")

	for i := 0; i < len(input); i++ {
		output = append(output, input[i]^key[i%len(key)])
	}

	return output, err
}

// XOR structure maintains a key value and satisfies Obfuscator interface.
type XOR struct {
	Key string `json:"key"`
}

func (x *XOR) ByteKey() []byte {
	return []byte(x.Key)
}

// Obfuscate XOR encrypts input bytes. Returns encrypted output and an error value.
func (x *XOR) Obfuscate(input []byte) (output []byte, err error) {
	return xor(input, x.ByteKey())
}

// Deobfuscate XOR decrypts input bytes. Returns decrypted output and an error value.
func (x *XOR) Deobfuscate(input []byte) (output []byte, err error) {
	return xor(input, x.ByteKey())
}
