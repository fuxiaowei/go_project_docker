package main

import "os"

func main() {
	println("go 程序开始执行................")
	name := os.Getenv("name")
	println("================" + name)
}
