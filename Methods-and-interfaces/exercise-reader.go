package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

//[復習]MyReaderに生やすのでレシーバを前に書く必要がある
func (r MyReader) Read(b []byte) (int, error) {
	size := len(b)

	//Read は、データを与えられたバイトスライスへ入れ、
	for i := 0; i < size; i++ {
		b[i] = 'A'
	}

	//入れたバイトのサイズとエラーの値を返します。
	return size, nil

	//ストリームの終端は、 io.EOF のエラーで返します。
	// => 今回は無限ストリームなので無し
}

func main() {
	reader.Validate(MyReader{})
}
