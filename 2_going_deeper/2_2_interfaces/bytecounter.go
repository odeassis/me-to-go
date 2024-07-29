package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))

	return len(p), nil
}

func main() {
	var msg ByteCounter

	msg.Write([]byte("Hello"))

	fmt.Println(msg)

	msg_2 := ByteCounter(0)

	var name = "odeassis"
	fmt.Fprint(&msg_2, "hello %s", name)

	fmt.Println(msg_2)
}
