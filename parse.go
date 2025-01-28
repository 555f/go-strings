package gostrings

import (
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

func ParseBool(s string, out *bool) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}

	*out = b

	return nil
}

func ParseComplex[T constraints.Complex](s string, bitSize int, out *T) error {
	i, err := strconv.ParseComplex(s, bitSize)
	if err != nil {
		return err
	}

	*out = T(i)

	return nil
}

func ParseInt[T constraints.Signed](s string, base, bitSize int, out *T) error {
	i, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		return err
	}

	*out = T(i)

	return nil
}

func ParseUint[T constraints.Unsigned](s string, base, bitSize int, out *T) error {
	i, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		return err
	}

	*out = T(i)

	return nil
}

func ParseFloat[T constraints.Float](s string, bitSize int, out *T) error {
	i, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		return err
	}

	*out = T(i)

	return nil
}

func ParseTime(layout string, s string, out *time.Time) error {
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*out = t

	return nil
}

func ParseDuration(s string, out *time.Duration) error {
	t, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	*out = t

	return nil
}
