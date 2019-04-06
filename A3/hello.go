package main

import "fmt"
import "github.com/ehteshamz/greetings"

func ConcatSlice(slicetoConcat []byte) string {
	s := string(slicetoConcat)
	fmt.Println(s)
	return s
}

func Encrypt(slicetoConcat []byte, ceaserCount int) []byte {
	for i := 0; i < len(slicetoConcat); i++ {
		slicetoConcat[i] = slicetoConcat[i] + 3
	}
	return slicetoConcat
}

func main() {
	temp := []byte{'a', 'r', 'h', 'a', 'm'}
	ConcatSlice(temp)
	letters := []byte{'a', 'r', 'h', 'a', 'm'}
	fmt.Println(string(Encrypt(letters, 3)))
	fmt.Println(greetings.PrintGreetings("Arham"))
}
