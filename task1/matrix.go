package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

//type Number[T numeric] T

type Row[N Numeric] []N

type Matrix[N Numeric] struct {
	size int
	rows []Row[N]
}

func scanRow[N Numeric](s *bufio.Scanner) (Row[N], int, error) {
	if !s.Scan() {
		return nil, 0, s.Err()
	}

	words := strings.Split(string(s.Text()), " ")
	values := make(Row[N], len(words))
	for i, word := range words {
		var tmp N
		_, err := fmt.Sscanf(word, "%v", &tmp)
		if err != nil {
			return nil, 0, err
		}

		values[i] = tmp
	}

	return values, len(values), nil
}

func (m *Matrix[N]) Read() error {
	s := bufio.NewScanner(os.Stdin)

	r, n, err := scanRow[N](s)
	if err != nil {
		return err
	}

	m.rows = make([]Row[N], n)
	m.rows[0] = r
	m.size = n

	for i := 1; i < m.size; i++ {
		r, n, err = scanRow[N](s)
		if err != nil {
			return err
		}
		if n != m.size {
			return errors.New("invalid row size")
		}

		m.rows[i] = r
	}

	return nil
}

func (m *Matrix[N]) Print() error {
	for _, r := range m.rows {
		var sep byte
		for _, n := range r {
			_, err := fmt.Printf("%c%v", sep, n)
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

func (m *Matrix[N]) copyRowsCols(rows, cols []int) Matrix[N] {
	copied := Matrix[N]{size: len(rows)}
	copied.rows = make([]Row[N], copied.size)

	var ii int
	for i, r := range m.rows {
		if isElement(i, rows) {
			copied.rows[ii] = make(Row[N], copied.size)
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

func (m *Matrix[N]) minor(rows, cols []int) (N, error) {
	if len(rows) != len(cols) {
		return 0, errors.New("minor: number of rows and columns don't match")
	}
	errorOutOfRange := errors.New("minor: index out of range")
	for _, row := range rows {
		if row >= m.size {
			return 0, errorOutOfRange
		}
	}
	for _, col := range cols {
		if col >= m.size {
			return 0, errorOutOfRange
		}
	}

	minor := m.copyRowsCols(rows, cols)

	d, err := minor.det()
	if err != nil {
		return 0, err
	}

	return d, nil
}

func (m *Matrix[N]) firstMinor(row, col int) (N, error) {
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

func (m *Matrix[N]) det() (N, error) {
	if m.size == 0 {
		return 0, nil
	}

	if m.size == 1 {
		return m.rows[0][0], nil
	}

	if m.size == 2 {
		return m.rows[0][0]*m.rows[1][1] - m.rows[0][1]*m.rows[1][0], nil
	}

	var d N
	for i := 0; i < m.size; i++ {
		minor, err := m.firstMinor(0, i)
		if err != nil {
			return 0, err
		}

		d += N(math.Pow(-1, float64(1+i+1))) *
			m.rows[0][i] *
			minor
	}

	return d, nil
}

func (m *Matrix[N]) Det() error {
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

func (m *Matrix[N]) rank(rows, cols []int) (int, error) {
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

				if d != 0 {
					return m.rank(rr, cc)
				}
			}
		}
	}

	return len(rows), nil
}

func (m *Matrix[N]) Rank() error {
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

func (m *Matrix[N]) copy() Matrix[N] {
	c := Matrix[N]{size: m.size}

	c.rows = make([]Row[N], m.size)
	for i, r := range m.rows {
		c.rows[i] = make(Row[N], m.size)
		copy(c.rows[i], r)
	}

	return c
}

func (m *Matrix[N]) Transpose() error {
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
