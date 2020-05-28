package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 循环，将数字转换成二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	if result == "" {
		result = "0"
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	forever()
	printFile("abc.txt")
	fmt.Println(
		convertToBin(13),
		convertToBin(0),
	)
}
