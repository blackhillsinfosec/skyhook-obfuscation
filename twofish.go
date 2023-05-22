package obfuscate

import (
    "golang.org/x/crypto/twofish"
)

// Twofish structure maintains a key value and satisfies Obfuscator interface.
type Twofish struct {
    Key string `json:"key"`
}

func (t *Twofish) ByteKey() []byte {
    return []byte(t.Key)
}

// Obfuscate Twofish encrypts input bytes. Returns encrypted output and an error value.
func (t *Twofish) Obfuscate(input []byte) (output []byte, err error) {
    tfcobj, err := twofish.NewCipher(expandKey(t.ByteKey()))

    if err != nil {
        return nil, err
    }

    output = blockEncrypt(tfcobj, input)
    return output, nil
}

// Deobfuscate Twofish encrypts input bytes. Returns decrypted output and an error value.
func (t *Twofish) Deobfuscate(input []byte) (output []byte, err error) {
    tfcobj, err := twofish.NewCipher(expandKey(t.ByteKey()))

    if err != nil {
        return nil, err
    }

    output = blockDecrypt(tfcobj, input)
    return output, nil
}
