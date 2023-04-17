package helpers

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func ParseComplex[T constraints.Complex](s string, bitSize int) (T, error) {
	i, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func ParseInt[T constraints.Signed](s string, base, bitSize int) (T, error) {
	i, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func ParseUint[T constraints.Unsigned](s string, base, bitSize int) (T, error) {
	i, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}

func ParseFloat[T constraints.Float](s string, bitSize int) (T, error) {
	i, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return 0, err
	}
	return T(i), nil
}
