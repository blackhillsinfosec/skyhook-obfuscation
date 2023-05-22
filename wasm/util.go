package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"syscall/js"
)

// sToBwrapper makes a function available to JS that will convert a
// string to an Uint8Array values.
func sToBwrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		var err error
		var b []uint8

		_, err = MinArgLen(args, 1)
		if err == nil {
			b, err = StoB(args[0])
		}

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {

			resolve, reject := args[0], args[1]

			go func() {

				if err != nil {
					reject.Invoke(js.Global().Get("Error").New(err.Error()))
				} else {
					out := js.Global().Get("Uint8Array").New(len(b))
					js.CopyBytesToJS(out, b)
					runtime.KeepAlive(out)
					resolve.Invoke(out)
				}

			}()

			return nil

		})

		pC := js.Global().Get("Promise")
		return pC.New(handler)

	})
}

// bToSwrapper makes a function available to JS that will convert
// an Uint88Array object to string.
func bToSwrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		var err error
		var s string

		_, err = MinArgLen(args, 1)
		if err == nil {
			s, err = BtoS(args[0])
		}

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {

			resolve, reject := args[0], args[1]

			go func() {

				if err != nil {
					reject.Invoke(js.Global().Get("Error").New(err.Error()))
				} else {
					resolve.Invoke(s)
				}

			}()

			return nil

		})

		pC := js.Global().Get("Promise")
		return pC.New(handler)

	})
}

// BaseArgs performs base input checking for WASM functions. It assumes that args
// will have at least two arguments, act and value, where act is the action
// being taken by the call -- "obf" or "deobf" --, and value is the binary data
// to be acted on. Any remaining arguments will be returned in rem, enabling the
// calling function to receive additional arguments.
//
// - Expected preliminary arguments: args{act string, value []uint8}
// - Valid act values: obf, deobf
//
// BaseArgs returns the act string and binary value, along with a slice  with
// any remaining js.Value arguments.
//
// An error is returned when args<2 or invalid values are supplied for act or
// value.
func BaseArgs(args []js.Value) (act string, value []uint8, rem []js.Value, err error) {

	_, err = MinArgLen(args, 2)

	if err == nil {
		act, err = StrArg(args[0], true)
		if !strings.HasPrefix(act, "deobf") && !strings.HasPrefix(act, "obf") {
			err = errors.New("action must be either \"obf\" or \"deobf\"")
		}
		rem = args[1:]
	}

	if err == nil {
		value, err = Uint8ArrayArg(args[1], true)
		rem = args[2:]
	}

	return act, value, rem, err
}

// MinArgLen gets the length of args and checks against req to
// determine if the expected number is present.
func MinArgLen(args []js.Value, req int) (int, error) {
	l := len(args)
	var err error
	if l < req {
		err = errors.New(fmt.Sprintf("%v args required", req))
	}
	return l, err
}

// Uint8ArrayArg returns an error should arg not be an Uint8Array.
//
// valReq determines if the argument should also have a length greater
// than 0, resulting in an error when byteLength is 0 and valReq is true.
func Uint8ArrayArg(arg js.Value, valReq bool) ([]uint8, error) {
	if arg.Type() != js.TypeObject {
		return nil, errors.New("Uint8Array object required")
	}

	// TODO this seems like weak checking
	//  could lead to crashes!
	//  couldn't find an obvious method of determining if the
	//  source arg is exactly a Uint8Array

	bLen := arg.Get("byteLength").Int()
	if valReq && bLen == 0 {
		return nil, errors.New("Uint8Array must have a length > 0")
	}

	input := make([]uint8, bLen)
	js.CopyBytesToGo(input, arg)

	return input, nil
}

// StoB converts a JS string to bytes.
func StoB(s js.Value) ([]uint8, error) {
	if st, err := StrArg(s, false); err != nil {
		return nil, err
	} else {
		return []byte(st), err
	}
}

// BtoS converts a Uint8Array to a string.
func BtoS(b js.Value) (string, error) {
	if b, err := Uint8ArrayArg(b, false); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

// IntArg returns an error should arg not be an integer.
func IntArg(arg js.Value) (int, error) {
	if arg.Type() != js.TypeNumber {
		return 0, errors.New("int value required")
	}
	return arg.Int(), nil
}

// StrArg returns an error should arg not be a JS string.
func StrArg(arg js.Value, valReq bool) (string, error) {

	if arg.Type() != js.TypeString {
		return "", errors.New("string value required")
	}

	val := arg.String()
	if valReq && len(val) == 0 {
		return "", errors.New("string length > 0 required")
	}

	return arg.String(), nil
}

func md5wrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {

		var err error
		var b []uint8

		_, err = MinArgLen(args, 1)
		if err == nil {
			b, err = Uint8ArrayArg(args[0], true)
		}

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {
			resolve, reject := args[0], args[1]

			go func() {
				if err != nil {
					reject.Invoke(js.Global().Get("Error").New(err.Error()))
				} else {
					h := md5.New()
					h.Write(b)
					b = h.Sum(nil)
					s := fmt.Sprintf("%x", b)
					b = []byte(s)
					out := js.Global().Get("Uint8Array").New(len(b))
					js.CopyBytesToJS(out, b)
					runtime.KeepAlive(out)
					resolve.Invoke(out)
				}
			}()
			return nil
		})

		pC := js.Global().Get("Promise")
		return pC.New(handler)

	})
}
