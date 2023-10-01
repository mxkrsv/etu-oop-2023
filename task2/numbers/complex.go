package numbers

import (
	"fmt"
)

type Complex[T StdlibNumeric] struct {
	real      T
	imaginary T
}

func (c *Complex[T]) New() *Complex[T] {
	return &Complex[T]{}
}

func (c *Complex[T]) FromString(s string) {
	//var a, b T
	//_, err := fmt.Sscanf(s, "%d%d", &a, &b)
	//c.real, c.imaginary = a, b
	_, err := fmt.Sscanf(s, "%d%d", &c.real, &c.imaginary)
	if err != nil {
		panic(err)
	}
}

func (c *Complex[T]) FromNumber(n T) {
	c.real = T(n)
	c.imaginary = 0
}

func (c *Complex[T]) String() string {
	if c == nil {
		panic("nullpo")
	}

	var sep byte
	if c.imaginary >= 0 {
		sep = '+'
	}

	return fmt.Sprintf("%v%c%vi", c.real, sep, c.imaginary)
}

func (c *Complex[T]) Add(other *Complex[T]) *Complex[T] {
	return &Complex[T]{real: c.real + other.real, imaginary: c.imaginary + other.imaginary}
}

func (c *Complex[T]) Sub(other *Complex[T]) *Complex[T] {
	return &Complex[T]{real: c.real - other.real, imaginary: c.imaginary - other.imaginary}
}

func (c *Complex[T]) Mul(other *Complex[T]) *Complex[T] {
	return &Complex[T]{real: c.real * other.real, imaginary: c.imaginary * other.imaginary}
}

func (c *Complex[T]) Div(other *Complex[T]) *Complex[T] {
	return &Complex[T]{real: c.real / other.real, imaginary: c.imaginary / other.imaginary}
}
