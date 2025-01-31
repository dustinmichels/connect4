package main

import "connect4/server"

const useNgrok = true

func main() {
	server.Run(useNgrok)
}
