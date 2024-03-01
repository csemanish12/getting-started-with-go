package main

import "fmt"

func main(){
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

// accept a string, loop over it a byte at a time and return
// reversed string at the end
func Reverse(s string) string {
	byteArray := []byte(s)
	for i,j := 0, len(byteArray)-1; i< len(byteArray)/2; i,j = i+1, j-1 {
		byteArray[i], byteArray[j] =  byteArray[j], byteArray[i]
	}
	return string(byteArray)
}
