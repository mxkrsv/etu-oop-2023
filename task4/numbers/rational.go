package numbers

import (
	"fmt"
)

type Rational[T StdlibIntegers] struct {
	numerator   T
	denominator T
}

func (r *Rational[T]) New() *Rational[T] {
	return &Rational[T]{}
}

func (r *Rational[T]) reduce() {
	if r.numerator%r.denominator == 0 {
		r.numerator /= r.denominator
	}
}

func (r *Rational[T]) FromString(s string) {
	_, err := fmt.Sscanf(s, "%d/%d", &r.numerator, &r.denominator)
	if err != nil {
		panic(err)
	}
	if r.denominator == 0 {
		panic("zero denominator")
	}
	r.reduce()
}

func (r *Rational[T]) FromNumber(n T) {
	r.numerator = T(n)
	r.denominator = 1
}

func (r *Rational[T]) String() string {
	if r == nil {
		panic("nullpo")
	}

	return fmt.Sprintf("%v/%v", r.numerator, r.denominator)
}

func gcd[T StdlibIntegers](a, b T) T {
	if a == 1 || a == 0 {
		return b
	} else if b == 1 || b == 0 {
		return a
	} else if a >= b {
		return gcd(a%b, b)
	} else {
		return gcd(a, b%a)
	}
}

func abs[T StdlibIntegers](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func lcm[T StdlibIntegers](a, b T) T {
	return (abs(a) * abs(b)) / gcd(a, b)
}

func (r *Rational[T]) Add(other *Rational[T]) *Rational[T] {
	cm := lcm(r.denominator, other.denominator)
	rr := Rational[T]{
		numerator:   r.numerator*(cm/r.denominator) + other.numerator*(cm/other.denominator),
		denominator: cm,
	}
	rr.reduce()
	return &rr
}

func (r *Rational[T]) inverse() *Rational[T] {
	return &Rational[T]{
		numerator:   -r.numerator,
		denominator: r.denominator,
	}
}

func (r *Rational[T]) Sub(other *Rational[T]) *Rational[T] {
	return r.Add(other.inverse())
}

func (r *Rational[T]) Mul(other *Rational[T]) *Rational[T] {
	rr := Rational[T]{
		numerator:   r.numerator * other.numerator,
		denominator: r.denominator * other.denominator,
	}
	rr.reduce()
	return &rr
}
