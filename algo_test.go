package obfuscate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var (
	testLens = []int{1, 4, 100, 301, 500, 1000, 2000}
	data     []byte
)

func init() {
	if f, err := os.Open("10M.data"); err == nil {
		data, _ = io.ReadAll(f)
		fmt.Printf("%v Bytes read from test file\n", len(data))
	} else {
		fmt.Printf("Failed to load test data: %v\n", err)
	}
}

func testAlgo(o Obfuscator, t *testing.T) {
	var val string
	for _, l := range testLens {
		for i := 0; i < l; i++ {
			val += "A"
		}
		var err error
		var obf, deobf []byte

		obf, err = o.Obfuscate([]byte(val))
		if err != nil {
			t.Fatal(err)
		}

		deobf, err = o.Deobfuscate(obf)
		if err != nil {
			t.Fatal(err)
		}

		if string(deobf) != val {
			t.Fatal("Deobfuscated value did mot match original value")
		}
	}

	if len(data) > 0 {

		var err error
		var obf, deobf []byte

		obf, err = o.Obfuscate(data)
		if err != nil {
			t.Fatal(err)
		}

		deobf, err = o.Deobfuscate(obf)
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(data, deobf) != 0 {
			t.Fatal("Deobfuscated file data did not match original input")
		}
	}
}
