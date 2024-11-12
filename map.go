package gostrings

import (
	"errors"
	"fmt"
	"sort"
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

func JoinKeyValInt[V constraints.Integer](values map[string]V, sep, sepKV string, base int) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + strconv.FormatInt(int64(values[k]), base)
	}
	return
}

func JoinKeyValUint[V constraints.Unsigned](values map[string]V, sep, sepKV string, base int) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + strconv.FormatUint(uint64(values[k]), base)
	}
	return
}

func JoinKeyValFloat[V constraints.Float](values map[string]V, sep, sepKV string, fmt byte, prec, bitSize int) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + strconv.FormatFloat(float64(values[k]), fmt, prec, bitSize)
	}
	return
}

func JoinKeyValString[V string](values map[string]V, sep, sepKV string) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + string(values[k])
	}
	return
}

func JoinKeyValTime[V time.Time](values map[string]V, sep, sepKV string, layout string) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + time.Time(values[k]).Format(layout)
	}
	return
}

func JoinKeyValDuration[V time.Duration](values map[string]V, sep, sepKV string) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + time.Duration(values[k]).String()
	}
	return
}

func JoinKeyValUUID[V uuid.UUID](values map[string]V, sep, sepKV string) (result string) {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		if i > 0 {
			result += sep
		}
		result += k + sepKV + uuid.UUID(values[k]).String()
	}
	return
}
