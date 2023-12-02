package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

type finderMap map[string]int

type numFinder struct {
	m       finderMap
	longest int
}

func newFinderMap() finderMap {
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for i := 0; i <= 9; i++ {
		m[strconv.Itoa(i)] = i
	}

	//slog.Info("newFinderMap", "m", fmt.Sprintf("%#v", m))

	return m
}

func newNumFinder() *numFinder {
	m := newFinderMap()

	longest := 0

	for k := range m {
		if l := len(k); l > longest {
			longest = l
		}
	}

	return &numFinder{
		m:       m,
		longest: longest,
	}
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func (f *numFinder) isNumber(s string) (bool, int) {
	//slog.Info("isNumber", "s", s)
	l := min(len(s), f.longest)

	for i := 1; i <= l; i++ {
		ss := s[0:i]
		if v, ok := f.m[ss]; ok {
			//slog.Info("isNumber", "ss", ss, "v", v)
			return true, v
		}
		//slog.Info("isNumber", "ss", ss)
	}

	return false, 0
}

func (f *numFinder) lastNumber(l string) (int, error) {
	//slog.Info("last")
	for i := len(l) - 1; i >= 0; i-- {
		if ok, v := f.isNumber(l[i:]); ok {
			return v, nil
		}
	}
	return 0, fmt.Errorf("lastNumber: number not found")
}

func (f *numFinder) firstNumber(l string) (int, error) {
	//slog.Info("first")
	for i := 0; i < len(l); i++ {
		if ok, v := f.isNumber(l[i:]); ok {
			return v, nil
		}
	}
	return 0, fmt.Errorf("firstNumber: number not found")
}

func (f *numFinder) processLine(line string) (int, error) {
	fmt.Println(line)
	first, err := f.firstNumber(line)
	if err != nil {
		return 0, fmt.Errorf("processLine firstNumber: %v", err)
	}

	last, err := f.lastNumber(line)
	if err != nil {
		return 0, fmt.Errorf("processLine lastNumber: %v", err)
	}

	//slog.Info("processLine", "line", line, "first", first, "last", last)
	return first*10 + last, nil
}

func (f *numFinder) processContents(s *bufio.Scanner) (int, error) {
	total := 0
	for s.Scan() {
		v, err := f.processLine(s.Text())

		if err != nil {
			return 0, fmt.Errorf("processContents processLine: %v", err)
		}
		total += v
		//slog.Info("processContents", "v", v, "total", total)
	}

	return total, nil
}

func readInput(r io.Reader) (*bufio.Scanner, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("process: %v", err)
	}

	s := bufio.NewScanner(bytes.NewReader(b))
	s.Buffer(make([]byte, len(b)), len(b))

	return s, nil
}

func process(args []string) error {
	s, err := readInput(os.Stdin)
	if err != nil {
		return fmt.Errorf("process readInput: %v", err)
	}

	f := newNumFinder()
	v, err := f.processContents(s)
	if err != nil {
		return fmt.Errorf("process processContents: %v", err)
	}
	fmt.Println(v)

	return nil
}
