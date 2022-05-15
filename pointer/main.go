package main

import (
	"reflect"
)

func createPointerType(t reflect.Type) reflect.Type {
	return reflect.PtrTo(t)
}

func followPointerType(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
func main() {
	name := "Alice"
	t := reflect.TypeOf(name)
	Printfln("Original Type: %v", t)
	pt := createPointerType(t)
	Printfln("Pointer Type: %v", pt)
	Printfln("Follow pointer type: %v", followPointerType(pt))
}
