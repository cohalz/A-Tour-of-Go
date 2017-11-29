package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

/*
	全ての型はinterface{}を実装しているのでどんな型でも受け入れるメソッドが作れる
	感覚としてはAnyに近い?
	def println(args: Any*)
	fmt.Printも可変長
*/
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
