package matrix

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/mxkrsv/etu-oop-2023/task3/numbers"
)

type Row[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]] []N

type Matrix[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]] struct {
	size int
	rows []Row[n, N]
}

func scanRow[n numbers.StdlibNumeric, N numbers.CustomNumeric[n, N]](s *bufio.Scanner) (Row[n, N], int, error) {
	if !s.Scan() {
		return nil, 0, s.Err()
	}

	words := strings.Split(string(s.Text()), " ")
	values := make(Row[n, N], len(words))
	for i, word := range words {
		var tmp N
		tmp = tmp.New()
		//fmt.Fprintf(os.Stderr, "%T\n", reflect.Indirect(reflect.ValueOf(tmp)).Type())
		//_, err := fmt.Sscanf(word, "%v", &tmp)
		//if err != nil {
		//	return nil, 0, err
		//}

		//tmp.FromString(word)

		tmp.FromString(word)

		values[i] = tmp

		//values[i] = numbers.CustomNumericFromString[n, N](word)
	}

	return values, len(values), nil
}

func (m *Matrix[n, N]) Read(reader io.Reader) error {
	s := bufio.NewScanner(reader)

	r, cnt, err := scanRow[n, N](s)
	if err != nil {
		return err
	}

	m.rows = make([]Row[n, N], cnt)
	m.rows[0] = r
	m.size = cnt

	for i := 1; i < m.size; i++ {
		r, cnt, err = scanRow[n, N](s)
		if err != nil {
			return err
		}
		if cnt != m.size {
			return errors.New("invalid row size")
		}

		m.rows[i] = r
	}

	return nil
}

func (m *Matrix[n, N]) Print() error {
	for _, r := range m.rows {
		var sep byte
		for _, n := range r {
			_, err := fmt.Printf("%c%v", sep, n)
			//_, err := fmt.Printf("%c%T", sep, n)
			if err != nil {
				return err
			}
			sep = ' '
		}
		_, err := fmt.Printf("\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func isElement[E comparable](e E, s []E) bool {
	for _, ee := range s {
		if e == ee {
			return true
		}
	}

	return false
}

func (m *Matrix[n, N]) copyRowsCols(rows, cols []int) Matrix[n, N] {
	copied := Matrix[n, N]{size: len(rows)}
	copied.rows = make([]Row[n, N], copied.size)

	var ii int
	for i, r := range m.rows {
		if isElement(i, rows) {
			copied.rows[ii] = make(Row[n, N], copied.size)
			var jj int

			for j, c := range r {
				if isElement(j, cols) {
					copied.rows[ii][jj] = c

					jj++
				}
			}

			ii++
		}
	}

	return copied
}

func (m *Matrix[n, N]) minor(rows, cols []int) (N, error) {
	if len(rows) != len(cols) {
		var zero N
		zero = zero.New()
		return zero, errors.New("minor: number of rows and columns don't match")
	}
	errorOutOfRange := errors.New("minor: index out of range")
	for _, row := range rows {
		if row >= m.size {
			var zero N
			zero = zero.New()
			return zero, errorOutOfRange
		}
	}
	for _, col := range cols {
		if col >= m.size {
			var zero N
			zero = zero.New()
			return zero, errorOutOfRange
		}
	}

	minor := m.copyRowsCols(rows, cols)

	d, err := minor.det()
	if err != nil {
		var zero N
		zero = zero.New()
		return zero, err
	}

	return d, nil
}

func (m *Matrix[n, N]) firstMinor(row, col int) (N, error) {
	rows := make([]int, m.size-1)
	cols := make([]int, m.size-1)

	var ii int
	for i := 0; i < m.size; i++ {
		if i != row {
			rows[ii] = i
			ii++
		}
	}
	ii = 0
	for i := 0; i < m.size; i++ {
		if i != col {
			cols[ii] = i
			ii++
		}
	}

	return m.minor(rows, cols)
}

func (m *Matrix[n, N]) det() (N, error) {
	if m.size == 0 {
		var zero N
		zero = zero.New()
		return zero, nil
	}

	if m.size == 1 {
		return m.rows[0][0], nil
	}

	if m.size == 2 {
		return m.rows[0][0].Mul(m.rows[1][1]).
			Sub(m.rows[0][1].Mul(m.rows[1][0])), nil
	}

	var d N
	d = d.New()
	for i := 0; i < m.size; i++ {
		minor, err := m.firstMinor(0, i)
		if err != nil {
			var zero N
			zero = zero.New()
			return zero, err
		}

		var tmp N
		tmp = tmp.New()
		tmp.FromNumber(n(math.Pow(-1, float64(1+i+1))))
		d.Add(tmp.Mul(m.rows[0][i].Mul(minor)))

		//d += N(math.Pow(-1, float64(1+i+1))) *
		//	m.rows[0][i] *
		//	minor
	}

	return d, nil
}

func (m *Matrix[n, N]) Det() error {
	d, err := m.det()
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%v\n", d)
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix[n, N]) rank(rows, cols []int) (int, error) {
	if len(rows) == m.size {
		return len(rows), nil
	}

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			if !isElement(i, rows) && !isElement(j, cols) {
				var rr, cc []int
				rr = append(rr, rows...)
				rr = append(rr, i)
				cc = append(cc, rows...)
				cc = append(cc, j)

				d, err := m.minor(rr, cc)
				if err != nil {
					return 0, err
				}

				var zero N
				zero = zero.New()
				if d != zero {
					return m.rank(rr, cc)
				}
			}
		}
	}

	return len(rows), nil
}

func (m *Matrix[n, N]) Rank() error {
	r, err := m.rank([]int{}, []int{})
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%v\n", r)
	if err != nil {
		return err
	}

	return nil
}

func (m *Matrix[n, N]) copy() Matrix[n, N] {
	c := Matrix[n, N]{size: m.size}

	c.rows = make([]Row[n, N], m.size)
	for i, r := range m.rows {
		c.rows[i] = make(Row[n, N], m.size)
		copy(c.rows[i], r)
	}

	return c
}

func (m *Matrix[n, N]) Transpose() error {
	t := m.copy()

	for i := 0; i < m.size; i++ {
		for j := 0; j < m.size; j++ {
			if i > j {
				t.rows[i][j], t.rows[j][i] = t.rows[j][i], t.rows[i][j]
			}
		}
	}

	t.Print()

	return nil
}
