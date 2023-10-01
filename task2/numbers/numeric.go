package numbers

import "fmt"

type StdlibNumeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type CustomNumeric[N StdlibNumeric, T any] interface {
	comparable
	fmt.Stringer
	FromString(string)
	FromNumber(N)
	New() T
	Add(T) T
	Sub(T) T
	Mul(T) T
	Div(T) T
}
/*
func NewCustomNumeric[N StdlibNumeric, T CustomNumeric[N, T]]() T {
	var ret T
	return ret
}

func CustomNumericFromString[N StdlibNumeric, T CustomNumeric[N, T]](s string) T {
	var ret T
	ret.FromString(s)
	return ret
}

func CustomNumericFromNumber[N StdlibNumeric, T CustomNumeric[N, T]](n N) T {
	var ret T
	ret.FromNumber(n)
	return ret
}
*/
