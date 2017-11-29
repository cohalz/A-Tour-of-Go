package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

//インタフェースの実装ではnilか確認し適切に処理するべき
func (t *T) M() {

	//実装しないとpanicを起こす
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t

	//ここは大丈夫
	describe(i)

	//ぬるぽ
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
