package main

import (
	obfs "github.com/blackhillsinfosec/skyhook-obfuscation"
	"syscall/js"
)

func xorWrapper() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) any {

		//==================
		// INPUT VALIDATIONS
		//==================

		var key []uint8
		act, value, args, err := BaseArgs(args)

		// Ensure key argument is supplied
		if err == nil {
			_, err = MinArgLen(args, 1)
		}

		// Extract the key argument
		if err == nil {
			key, err = Uint8ArrayArg(args[0], true)
		}

		//===========================
		// DEFINE THE PROMISE HANDLER
		//===========================

		handler := js.FuncOf(func(this js.Value, args []js.Value) any {

			resolve, reject := args[0], args[1]

			go func() {

				if err != nil {

					//==================
					// RETURN ANY ERRORS
					//==================

					reject.Invoke(js.Global().Get("Error").New(err.Error()))

				} else {

					//=================
					// APPLY XOR ACTION
					//=================

					xor := obfs.XOR{Key: string(key)}
					switch act {
					case "obf":
						value, err = xor.Obfuscate(value)
					case "deobf":
						value, err = xor.Deobfuscate(value)
					}

				}

				//==============
				// RETURN OUTPUT
				//==============

				out := js.Global().Get("Uint8Array").New(len(value))
				js.CopyBytesToJS(out, value)
				value = nil
				//runtime.KeepAlive(out)
				resolve.Invoke(out)

			}()

			return nil

		})

		//===================
		// RETURN THE PROMISE
		//===================

		return js.Global().Get("Promise").New(handler)

	})

}
