package aoc

import (
	"fmt"
	"io"
)

var _ Parser[int] = ParserFunc[int](nil)

type ParserFunc[T any] func(io.Reader) (T, error)

func NewParser[T any](fn ParserFunc[T]) Parser[T] {
	return ParserFunc[T](fn)
}

func (f ParserFunc[T]) Parse(r io.Reader) (T, error) {
	return f(r)
}

type Parser[Input any] interface {
	Parse(io.Reader) (Input, error)
}

var _ Solver[int] = SolverFunc[int](nil)

type SolverFunc[Input any] func(Input) string

func (f SolverFunc[Input]) Solve(i Input) string {
	return f(i)
}

func NewSolver[Input any](fn SolverFunc[Input]) Solver[Input] {
	return SolverFunc[Input](fn)
}

type Solver[Input any] interface {
	Solve(Input) string
}

type Day[Input any] struct {
	Day     int
	parser  Parser[Input]
	partOne Solver[Input]
	partTwo Solver[Input]
}

func NewDay[Input any](day int, parser Parser[Input], partOne Solver[Input], partTwo Solver[Input]) *Day[Input] {
	return &Day[Input]{Day: day, parser: parser, partOne: partOne, partTwo: partTwo}
}

func (d *Day[Input]) load() (Input, error) {
	reader, err := Get(d.Day)
	if err != nil {
		var zero Input
		return zero, fmt.Errorf("get: %w", err)
	}

	input, err := d.parser.Parse(reader)
	if err != nil {
		var zero Input
		return zero, fmt.Errorf("parse: %w", err)
	}

	return input, nil
}

func (d *Day[Input]) PartOne() (string, error) {
	input, err := d.load()
	if err != nil {
		return "", err
	}

	return d.partOne.Solve(input), nil
}

func (d *Day[Input]) PartTwo() (string, error) {
	input, err := d.load()
	if err != nil {
		return "", err
	}

	return d.partTwo.Solve(input), nil
}

func (d *Day[Input]) Run(w io.Writer) {
	fmt.Fprintf(w, "%s\n", d)

	if res, err := d.PartOne(); err != nil {
		fmt.Fprintf(w, "Part One: %v\n", err)
	} else {
		fmt.Fprintf(w, "Part One: %s\n", res)
	}

	if res, err := d.PartTwo(); err != nil {
		fmt.Fprintf(w, "Part Two: %v\n", err)
	} else {
		fmt.Fprintf(w, "Part Two: %s\n", res)
	}
}

func (d *Day[Input]) String() string {
	return fmt.Sprintf("Day %d", d.Day)
}

func Run[T any](day int, w io.Writer, parser Parser[T], partOne Solver[T], partTwo Solver[T]) {
	d := NewDay(day, parser, partOne, partTwo)
	d.Run(w)
}
