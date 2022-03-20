package main

import (
	v1 "github.com/donvito/hellomod"
	v2 "github.com/donvito/hellomod/v2"
)

func main() {
	v1.SayHello()
	v2.SayHello("Sadith")
}
