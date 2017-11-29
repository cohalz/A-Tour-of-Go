package main

import (
	"fmt"
)

//1. 追加したい型に対してエラー型を作る
type ErrNegativeSqrt float64

//2. 作ったエラー型をレシーバにしてErrorメソッドを実装する
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		//3. エラーを返したいタイミングでエラー型に包んで返す
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 1; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
		fmt.Println(z)
	}

	//4. エラーが起きなかった場合はnilを返す
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
