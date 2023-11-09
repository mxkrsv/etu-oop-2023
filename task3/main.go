package main

import (
	"github.com/mxkrsv/etu-oop-2023/task3/application/cli"
	//"github.com/mxkrsv/etu-oop-2023/task3/application/gui"
	"github.com/mxkrsv/etu-oop-2023/task3/numbers"
)

type application interface {
	Run()
}

func main() {
	var a application = cli.NewApplication[int32, *numbers.Rational[int32]]()
	a.Run()
}
