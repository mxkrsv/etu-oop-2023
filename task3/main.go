package main

import (
	"bufio"
	"os"

	"github.com/mxkrsv/etu-oop-2023/task2/application"
	"github.com/mxkrsv/etu-oop-2023/task2/numbers"
)

func main() {
	a := application.NewApplication[int32, *numbers.Rational[int32]]()
	a.PrintUsage()

	s := bufio.NewScanner(os.Stdin)
	os.Stdout.WriteString("> ")
	for s.Scan() {
		err := a.DispatchCommand(s.Text())
		if err != nil {
			panic(err)
		}

		os.Stdout.WriteString("> ")
	}

	if s.Err() != nil {
		panic(s.Err())
	}
}
