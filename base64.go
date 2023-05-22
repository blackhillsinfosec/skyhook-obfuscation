package obfuscate

import (
    "encoding/base64"
)

// Base64 structure maintains a rounds value and satisfies Obfuscator interface.
type Base64 struct {
    Rounds uint `json:"rounds"`
}

// Obfuscate Base64 encodes input bytes by a number of rounds. Returns encoded output and an error value.
func (b *Base64) Obfuscate(input []byte) (output []byte, err error) {
    return b.run(input, "obf")
}

// Deobfuscate Base64 decodes input bytes by a number of rounds. Returns decoded output and an error value.
func (b *Base64) Deobfuscate(input []byte) (output []byte, err error) {
    return b.run(input, "deobf")
}

// run enables the "rounds" loop to be implemented only once.
//
// The act parameter determines if obfuscation or deobfuscation occurs.
//
// - obf - Indicates obfuscation.
// - deobf - Indicates deobfuscation.
func (b *Base64) run(input []byte, act string) (output []byte, err error) {
    for r := b.Rounds; r > 0; r-- {

        if act == "obf" {
            output = Base64Encode(input)
        } else {
            if output, err = Base64Decode(input); err != nil {
                break
            }
        }

        if r != 1 {
            input = make([]byte, len(output))
            copy(input, output)
        }

    }
    return output, err
}

func Base64Encode(in []byte) (out []byte) {
    out = make([]byte, base64.StdEncoding.EncodedLen(len(in)))
    base64.StdEncoding.Encode(out, in)
    return out
}

func Base64Decode(in []byte) (out []byte, err error) {
    out = make([]byte, base64.StdEncoding.DecodedLen(len(in)))
    var n int
    if n, err = base64.StdEncoding.Decode(out, in); err == nil {
        out = out[:n]
    }
    return out, err
}
