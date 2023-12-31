package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/mxkrsv/etu-oop-2023/task4/matrix"
	"github.com/mxkrsv/etu-oop-2023/task4/numbers"
)

type command struct {
	name   string
	desc   string
	action func() error
}

type Application[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]] struct {
	commands []command
	matrix   matrix.Matrix[n, N]
}

func NewApplication[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]]() Application[n, N] {
	a := Application[n, N]{}
	a.commands = []command{
		{
			name:   "read",
			desc:   "Read a matrix from the terminal",
			action: func() error { return a.matrix.Read(os.Stdin) },
		},
		{
			name: "det",
			desc: "Calculate and print the determinant of a matrix",
			action: func() error {
				det, err := a.matrix.Det()
				if err != nil {
					return err
				}
				_, err = fmt.Printf("%v\n", det)
				return err
			},
		},
		{
			name: "transpose",
			desc: "Calculate and print the transpose of a matrix",
			action: func() error {
				err := a.matrix.Transpose()
				if err != nil {
					return err
				}
				_, err = fmt.Printf("%v\n", a.matrix)
				return err
			},
		},
		{
			name: "rank",
			desc: "Calculate and print the rank of a matrix",
			action: func() error {
				rank, err := a.matrix.Rank()
				if err != nil {
					return err
				}
				_, err = fmt.Printf("%v\n", rank)
				return err
			},
		},
		{
			name: "print",
			desc: "Print a matrix to the terminal",
			action: func() error {
				_, err := fmt.Printf("%v\n", a.matrix)
				return err
			},
		},
		{
			name:   "exit",
			desc:   "Exit from the application",
			action: func() error {
				os.Exit(0)
				return nil
			},
		},
	}

	return a
}

func (a Application[n, N]) PrintUsage() error {
	_, err := fmt.Printf("Using type: %T\n", *new(N))
	fmt.Println("----")
	_, err = fmt.Println("Commands:")
	if err != nil {
		return err
	}

	for _, c := range a.commands {
		_, err = fmt.Printf("%s: %s\n", c.name, c.desc)
		if err != nil {
			return err
		}
	}

	fmt.Println("====")

	return nil
}

func (a Application[n, N]) DispatchCommand(c string) error {
	for _, cmd := range a.commands {
		if c == cmd.name {
			err := cmd.action()
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("unknown command")
}

func (a Application[n, N]) Run() {
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
}
