package main

import (
	"syscall/js"
)

func main() {
	//=========================
	// MAKE FUNCTIONS AVAILABLE
	//=========================

	js.Global().Set("skyBtos", bToSwrapper())
	js.Global().Set("skyStob", sToBwrapper())
	js.Global().Set("skyMd5Sum", md5wrapper())

	js.Global().Set("skyB64", b64Wrapper())
	js.Global().Set("skyBlowfish", blowfishWrapper())
	js.Global().Set("skyXor", xorWrapper())
	js.Global().Set("skyAes", aesWrapper())
	js.Global().Set("skyTwofish", twofishWrapper())

	//=====================
	// CHANNEL FOR BLOCKING
	//=====================

	<-make(chan bool)
}
