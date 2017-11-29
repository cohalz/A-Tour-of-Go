package main

import (
	"io"
	"os"
	"strings"
)

//変換して返すためにio.Readerをラップ
type rot13Reader struct {
	r io.Reader
}

//ここのレシーバは必須ではない
func (r rot13Reader) Rot13(c byte) byte {
	switch {
	case 'A' <= c && c <= 'Z':
		return 'A' + ((c - 'A' + 13) % 26)
	case 'a' <= c && c <= 'z':
		return 'a' + ((c - 'a' + 13) % 26)
	default:
		return c
	}
}

//rot13ReaderのRead実装
func (r rot13Reader) Read(b []byte) (int, error) {

	//1. まず普通にio.Readerから読み出す
	n, err := r.r.Read(b)

	//2. 読み出したものを変換
	for i := 0; i < n; i++ {
		b[i] = r.Rot13(b[i])
	}

	//3. 換字のみなのでnとerrはそのまま伝播
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	//io.Copyの中でバイト列が作られてるのであんま考えなくていい
	io.Copy(os.Stdout, &r)
}
