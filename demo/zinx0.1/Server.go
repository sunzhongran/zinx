package main

import "zinx/znet"

func main() {
	s := znet.NewServer("0.1")

	s.Server()
}
