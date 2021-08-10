package main

import "github.com/martinsrso/service-price/cmd"

var rootCommand = cmd.RootCommand

func main() {
	rootCommand().Execute()
}
