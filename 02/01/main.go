package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hoge.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	str := fmt.Sprintln("hogehoge")
	file.Write([]byte(str))

	num := fmt.Sprintf("num:%d\n", 100)
	file.Write([]byte(num))

	floatString := fmt.Sprintf("float:%2f\n", 1.2345)
	file.Write([]byte(floatString))
}
