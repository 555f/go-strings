package gostrings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/constraints"
)

func SplitKeyValInt[V constraints.Signed](s, sep, sepKV string, base, bitSize int) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		i, err := ParseInt[V](kv[1], base, bitSize)
		if err != nil {
			return nil, err
		}
		result[kv[0]] = i
	}
	return
}

func SplitKeyValUint[V constraints.Unsigned](s, sep, sepKV string, base, bitSize int) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		i, err := ParseUint[V](kv[1], base, bitSize)
		if err != nil {
			return nil, err
		}
		result[kv[0]] = i
	}
	return
}

func SplitKeyValFloat[V constraints.Float](s, sep, sepKV string, bitSize int) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		i, err := strconv.ParseFloat(kv[1], bitSize)
		if err != nil {
			return nil, err
		}
		result[kv[0]] = V(i)
	}
	return
}

func SplitKeyValString[V ~string](s, sep, sepKV string) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		result[kv[0]] = V(kv[1])
	}
	return
}

func SplitKeyValTime[V time.Time](s, sep, sepKV, layout string) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := time.Parse(layout, kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid date format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}

func SplitKeyValDuration[V time.Duration](s, sep, sepKV string) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := time.ParseDuration(kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}

func SplitKeyValUUID[V uuid.UUID](s, sep, sepKV string) (result map[string]V, err error) {
	if s == "" {
		return map[string]V{}, nil
	}
	parts := strings.Split(s, sep)
	result = make(map[string]V, len(parts))
	for _, v := range parts {
		kv := strings.Split(v, sepKV)
		if len(kv) != 2 {
			return nil, errors.New("invalid string format, should be 'key" + sepKV + "val" + sep + "key" + sepKV + "val'")
		}
		t, err := uuid.Parse(kv[1])
		if err != nil {
			return nil, fmt.Errorf("invalid duration format for key %s: %w", kv[0], err)
		}
		result[kv[0]] = V(t)
	}
	return
}
