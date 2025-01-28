package gostrings

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseBool(s string, recv *bool) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*recv = b
	return nil
}

func ParseComplex[T constraints.Complex](s string, bitSize int, recv *T) error {
	i, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		return err
	}
	*recv = T(i)
	return nil
}

func ParseInt[T constraints.Signed](s string, base, bitSize int, recv *T) error {
	i, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		return err
	}
	*recv = T(i)
	return nil
}

func ParseUint[T constraints.Unsigned](s string, base, bitSize int, recv *T) error {
	i, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		return err
	}
	*recv = T(i)
	return nil
}

func ParseFloat[T constraints.Float](s string, bitSize int, recv *T) error {
	i, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return err
	}
	*recv = T(i)
	return nil
}
