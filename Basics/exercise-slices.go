package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	/*
	hoge [i]intはあまり使わない?
	上の形式は数字または定数しかできない
	i := 3(var i = 3)
	var hoge [i]int
	=> 不可

	const j = 3
	var fuga [j]int
	=> 可
	*/
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		pic[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pic[y][x] = uint8(x * y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
