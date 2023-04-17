package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/exp/constraints"
)

func Split(s, sep string) (result []string, err error) {
	return strings.Split(s, sep), nil
}

func SplitInt[V constraints.Signed](s, sep string, base, bitSize int) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		i, err := ParseInt[V](v, base, bitSize)
		if err != nil {
			return nil, err
		}
		result[idx] = i
	}
	return
}

func SplitUint[V constraints.Unsigned](s, sep string, base, bitSize int) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		i, err := ParseUint[V](v, base, bitSize)
		if err != nil {
			return nil, err
		}
		result[idx] = i
	}
	return
}

func SplitFloat[V constraints.Float](s, sep string, bitSize int) (result []V, err error) {
	parts := strings.Split(s, sep)
	result = make([]V, len(parts))
	for idx, v := range parts {
		i, err := ParseFloat[V](v, bitSize)
		if err != nil {
			return nil, err
		}
		result[idx] = i
	}
	return
}

func SplitTime(s, sep, sepKV, layout string) (result []time.Time, err error) {
	parts := strings.Split(s, sep)
	result = make([]time.Time, len(parts))
	for idx, v := range parts {
		t, err := time.Parse(layout, v)
		if err != nil {
			return nil, fmt.Errorf("invalid date format for index %d: %w", idx, err)
		}
		result[idx] = t
	}
	return
}

func SplitDuration(s, sep string) (result []time.Duration, err error) {
	parts := strings.Split(s, sep)
	result = make([]time.Duration, len(parts))
	for idx, v := range parts {
		t, err := time.ParseDuration(v)
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for index %d: %w", idx, err)
		}
		result[idx] = t
	}
	return
}

func SplitUUID(s, sep string) (result []uuid.UUID, err error) {
	parts := strings.Split(s, sep)
	result = make([]uuid.UUID, len(parts))
	for idx, v := range parts {
		t, err := uuid.Parse(v)
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %d: %w", idx, err)
		}
		result[idx] = t
	}
	return
}

func JoinInt[V constraints.Integer](values []V, sep string, base int) (result string) {
	for i, v := range values {
		if i > 0 {
			result += sep
		}
		result += strconv.FormatInt(int64(v), base)
	}
	return
}

func JoinFloat[V constraints.Float](values []V, sep string, fmt byte, prec int, bitSize int) (result string) {
	for i, v := range values {
		if i > 0 {
			result += sep
		}
		result += strconv.FormatFloat(float64(v), fmt, prec, bitSize)
	}
	return
}
