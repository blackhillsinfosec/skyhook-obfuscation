package obfuscate

// Obfuscator interface allows for homogenous obfuscation chaining.
type Obfuscator interface {
	Obfuscate(input []byte) (output []byte, err error)
	Deobfuscate(input []byte) (output []byte, err error)
}

// Obfuscate operates on a chain of Obfuscator objects and returns the obfuscated output
func Obfuscate(input []byte, chain []Obfuscator) (out []byte, err error) {
	out = input
	for _, algo := range chain {
		out, err = algo.Obfuscate(out)
	}
	out = Base64Encode(out)
	return out, err
}

// Deobfuscate operates on a chain of Obfuscator objects and returns the deobfuscated output
func Deobfuscate(input []byte, chain []Obfuscator) (out []byte, err error) {
	if input, err = Base64Decode(input); err == nil {
		out = input
		for i := len(chain) - 1; i >= 0; i-- {
			out, _ = chain[i].Deobfuscate(out)
			//out = bytes.TrimRight(out, "\x00")
		}
	}
	return out, err
}
