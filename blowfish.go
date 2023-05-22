package obfuscate

import (
	"errors"
	"golang.org/x/crypto/blowfish"
)

// Blowfish structure maintains a key and salt value and satisfies Obfuscator interface. Salt is optional.
type Blowfish struct {
	Key  string `json:"key"`
	Salt string `json:"salt"`
}

func (b *Blowfish) ByteKey() []byte {
	return []byte(b.Key)
}

func (b *Blowfish) ByteSalt() []byte {
	return []byte(b.Salt)
}

// Obfuscate Blowfish encrypts input bytes, optionally salted. Returns encrypted output and an error value.
func (b *Blowfish) Obfuscate(input []byte) (output []byte, err error) {
	key, salt := b.ByteKey(), b.ByteSalt()
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("Blowfish key must be between 1 and 56 bytes")
	}

	var bfcobj *blowfish.Cipher

	if salt != nil {
		bfcobj, err = blowfish.NewSaltedCipher(key, salt)
	} else {
		bfcobj, err = blowfish.NewCipher(key)
	}

	if err != nil {
		return nil, err
	}

	output = blockEncrypt(bfcobj, input)
	return output, nil
}

// Deobfuscate Blowfish decrypts input bytes, optionally salted. Returns decrypted output and an error value.
func (b *Blowfish) Deobfuscate(input []byte) (output []byte, err error) {
	key, salt := b.ByteKey(), b.ByteSalt()
	if len(key) < 1 || len(key) > 56 {
		return nil, errors.New("Blowfish key must be between 1 and 56 bytes")
	}

	var bfcobj *blowfish.Cipher

	if salt != nil {
		bfcobj, err = blowfish.NewSaltedCipher(key, salt)
	} else {
		bfcobj, err = blowfish.NewCipher(key)
	}

	if err != nil {
		return nil, err
	}

	output = blockDecrypt(bfcobj, input)
	return output, nil
}
