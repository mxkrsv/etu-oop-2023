package main

import (
	//"github.com/mxkrsv/etu-oop-2023/task4/application/cli"
	"github.com/mxkrsv/etu-oop-2023/task4/application/gui"
	"github.com/mxkrsv/etu-oop-2023/task4/numbers"
)

type application interface {
	Run()
}

func main() {
	var a application = gui.NewApplication[int32, *numbers.Rational[int32]]()
	a.Run()
}
