package obfuscate

import (
    "testing"
)

func TestObfuscate(t *testing.T) {
    chain := []Obfuscator{
        &XOR{Key: "secret"},
        &XOR{Key: "another secret"},
        &XOR{Key: "and another secret"},
    }

    input := "it's a secret"

    out, err := Obfuscate([]byte(input), chain)
    if err != nil {
        t.Log("Failed to obfuscate input")
        t.Fatal(err)
    }

    out, err = Deobfuscate(out, chain)
    if err != nil {
        t.Log("Failed to deobfuscate output")
        t.Fatal(err)
    }

    if string(out) != input {
        t.Fatalf("Deobfuscated output failed to match string input: '%s' != '%s'\n", out, input)
    }
}
