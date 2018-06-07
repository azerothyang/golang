package main

import (
	"io"
	"strings"
	"fmt"
)

func main() {
	str := strings.NewReader("hello")
	//bt, _ := readFrom(str, 3)
	//fmt.Println(string(bt))

	//bt, _ = readFrom(os.Stdin, 100)
	//fmt.Println(string(bt))

	p := make([]byte, 2)
	str.ReadAt(p, 2)
	fmt.Printf("%s\n", p)
}

func readFrom(reader io.Reader, num int) ([]byte, error) {
	content := make([]byte, num)
	n, err := reader.Read(content)
	//如果读到内容, 就返回
	if n > 0 {
		return content[:num], err
	}
	return content, err
}
