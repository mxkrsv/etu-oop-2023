package main

import (
	"bufio"
	"github.com/mxkrsv/etu-oop-2023/task1/application"
	"os"
)

func main() {
	a := application.NewApplication[float64]()
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
