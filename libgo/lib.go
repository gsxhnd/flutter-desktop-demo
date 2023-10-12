// filename: lib.go
package main

import "C"

//export GetKey
func GetKey() *C.char {
	theKey := "123-456-789"
	return C.CString(theKey)
}

func main() {}
