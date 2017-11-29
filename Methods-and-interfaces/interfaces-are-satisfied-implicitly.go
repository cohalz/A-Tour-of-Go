package main

import "fmt"

/*
interfaceはabstractのようなもの
M()メソッドが生えている型に対して
var hoge I = ???が可能になる
*/

//共通化したいインタフェースを宣言
//型Iとして受け取り呼び出すにはその型でMの実装が必要
type I interface {
	M()
}

//Mを実装したい型
type T struct {
	S string
}

//MはメソッドなのでT型のレシーバを受け取り実装
//この実装がないとcannot use T literal (type T) as type I in assignment:
//	T does not implement I (missing M method)エラーが発生
func (t T) M() {
	fmt.Println(t.S)
}

//T型の変数を作っているがIのインタフェースを実装しているのでIとして受け取れる
func main() {
	var i I = T{"hello"}
	i.M()
	/*
	i := T{"hello"}
	var i T = T{"hello"}
	ももちろん可能
	*/
}
