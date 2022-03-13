package main

import "reflect"

func main() {
	// t := reflect.TypeOf(3)  // a reflect.Type
	// fmt.Println(t.String()) // "int"
	// fmt.Println(t)          // "int"

	// var w io.Writer = os.Stdout
	// fmt.Println(reflect.TypeOf(w)) // "*os.File"

	// v := reflect.ValueOf(3) // a reflect.Value
	// fmt.Println(v)          // "3"
	// fmt.Printf("%v\n", v)   // "3"
	// fmt.Println(v.String()) // NOTE: "<int Value>"

	// t = v.Type()            // a reflect.Type
	// fmt.Println(t.String()) // "int"

	// v = reflect.ValueOf(3) // a reflect.Value
	// x := v.Interface()     // an interface{}
	// i := x.(int)           // an int
	// fmt.Printf("%d\n", i)  // "3"

	// var a int64 = 1
	// var d time.Duration = 1 * time.Nanosecond
	// fmt.Println(Any(a))                  // "1"
	// fmt.Println(Any(d))                  // "1"
	// fmt.Println(Any([]int64{a}))         // "[]int64 0x8202b87b0"
	// fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"

	// strangelove := Movie{
	// 	Title:    "Dr. Strangelove",
	// 	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	// 	Year:     1964,
	// 	Actor: map[string]string{
	// 		"Dr. Strangelove":            "Peter Sellers",
	// 		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
	// 		"Pres. Merkin Muffley":       "Peter Sellers",
	// 		"Gen. Buck Turgidson":        "George C. Scott",
	// 		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
	// 		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	// 	},

	// 	Oscars: []string{
	// 		"Best Actor (Nomin.)",
	// 		"Best Adapted Screenplay (Nomin.)",
	// 		"Best Director (Nomin.)",
	// 		"Best Picture (Nomin.)",
	// 	},
	// }

	// Display("strangelove", strangelove)
	// Display("os.Stderr", os.Stderr)
	// var b interface{} = 3
	// Display("b", b)
	// Display("&b", &b)

	// content, err := Marshal(strangelove)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(content))

	// xx := 2                    // value   type   variable?
	// aa := reflect.ValueOf(2)   // 2       int    no
	// bb := reflect.ValueOf(xx)  // 2       int    no
	// cc := reflect.ValueOf(&xx) // &x      *int   no
	// dd := cc.Elem()            // 2       int    yes (x)

	// fmt.Println(aa.CanAddr()) // "false"
	// fmt.Println(bb.CanAddr()) // "false"
	// fmt.Println(cc.CanAddr()) // "false"
	// fmt.Println(dd.CanAddr()) // "true"

	// x := 2
	// d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	// px := d.Addr().Interface().(*int) // px := &x
	// *px = 3                           // x = 3
	// fmt.Println(x)                    // "3"
	// d.Set(reflect.ValueOf(4))
	// fmt.Println(x) // "4"

	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)                     // OK, x = 2
	rx.Set(reflect.ValueOf(3))       // OK, x = 3
	rx.SetString("hello")            // panic: string is not assignable to int
	rx.Set(reflect.ValueOf("hello")) // panic: string is not assignable to int

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	ry.SetInt(2)                     // panic: SetInt called on interface Value
	ry.Set(reflect.ValueOf(3))       // OK, y = int(3)
	ry.SetString("hello")            // panic: SetString called on interface Value
	ry.Set(reflect.ValueOf("hello")) // OK, y = "hello"
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}
