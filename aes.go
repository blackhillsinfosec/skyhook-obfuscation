package obfuscate

import (
    "crypto/aes"
)

// expandKey expands/truncates an AES key to be 16, 24, or 32 bytes in length.
func expandKey(key []byte) []byte {
    keylen := len(key)
    var keyout []byte = make([]byte, keylen)

    switch keylen {
    case 16:
    case 24:
    case 32:
        return key
    }

    copy(keyout, key)

    if keylen < 16 {
        for i := 0; i < 16-keylen; i++ {
            keyout = append(keyout, key[i%keylen])
        }
    } else if keylen < 24 {
        for i := 0; i < 24-keylen; i++ {
            keyout = append(keyout, key[i%keylen])
        }
    } else if keylen < 32 {
        for i := 0; i < 32-keylen; i++ {
            keyout = append(keyout, key[i%keylen])
        }
    } else {
        keyout = make([]byte, 32)
        copy(keyout, key[:32])
    }

    return keyout
}

// AES structure maintains a key value and satisfies Obfuscator interface.
type AES struct {
    Key string `json:"key"`
}

// ByteKey returns Key as a byte slice.
func (a *AES) ByteKey() []byte {
    return []byte(a.Key)
}

// Obfuscate AES encrypts input bytes. Returns encrypted output and an error value.
func (a *AES) Obfuscate(input []byte) (output []byte, err error) {
    aesobj, err := aes.NewCipher(expandKey(a.ByteKey()))
    output = []byte("")

    if err != nil {
        return nil, err
    }

    output = blockEncrypt(aesobj, input)
    return output, nil
}

// Deobfuscate AES decrypts input bytes. Returns decrypted output and an error value.
func (a *AES) Deobfuscate(input []byte) (output []byte, err error) {
    aesobj, err := aes.NewCipher(expandKey(a.ByteKey()))

    if err != nil {
        return nil, err
    }

    output = blockDecrypt(aesobj, input)
    return output, nil
}
