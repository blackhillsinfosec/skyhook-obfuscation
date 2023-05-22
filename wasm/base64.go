package main

import (
    obfs "github.com/blackhillsinfosec/skyhook-obfuscation"
    "syscall/js"
)

func b64Wrapper() js.Func {
    return js.FuncOf(func(this js.Value, args []js.Value) any {

        //==================
        // INPUT VALIDATIONS
        //==================

        var rounds int
        act, value, args, err := BaseArgs(args)

        if err == nil {
            _, err = MinArgLen(args, 1)
        }
        if err == nil {
            rounds, err = IntArg(args[0])
        }

        //================
        // PROMISE HANDLER
        //================

        handler := js.FuncOf(func(this js.Value, args []js.Value) any {

            resolve, reject := args[0], args[1]

            go func() {

                //=============================
                // DECODE AND RETURN THE OUTPUT
                //=============================

                if err != nil {

                    //===========
                    // BAD INPUTS
                    //===========

                    reject.Invoke(js.Global().Get("Error").New(err.Error()))

                } else {

                    //=================
                    // INPUTS VALIDATED
                    //=================

                    var output []uint8
                    b64 := obfs.Base64{Rounds: uint(rounds)}

                    var f func([]byte) ([]byte, error)

                    switch act {
                    case "obf":
                        f = b64.Obfuscate
                    case "deobf":
                        f = b64.Deobfuscate
                    }
                    output, err = f(value)

                    if err == nil {
                        // Get a new Uint8Array
                        out := js.Global().Get("Uint8Array").New(len(output))

                        // Copy decoded bytes to the new array
                        js.CopyBytesToJS(out, output)

                        // Not really sure what this is for
                        // https://github.com/golang/go/issues/32402
                        //runtime.KeepAlive(out)

                        resolve.Invoke(out)
                    } else {
                        reject.Invoke(js.Global().Get("Error").New(err.Error()))
                    }
                }
            }()

            return nil

        })

        promiseConstructor := js.Global().Get("Promise")
        return promiseConstructor.New(handler)

    })

}
