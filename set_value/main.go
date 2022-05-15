package main

import (
	"reflect"
)

func setAll(src interface{}, targets ...interface{}) {
	srcVal := reflect.ValueOf(src)
	for _, target := range targets {
		targetVal := reflect.ValueOf(target)
		if targetVal.Kind() == reflect.Ptr &&
			targetVal.Elem().Type() == srcVal.Type() &&
			targetVal.Elem().CanSet() {
			targetVal.Elem().Set(srcVal)
		}
	}
}
func main() {
	name := "Alice"
	price := 279
	city := "London"
	setAll("New String", &name, &price, &city)
	setAll(10, &name, &price, &city)
	for _, val := range []interface{}{name, price, city} {
		Printfln("Value: %v", val)
	}
}
