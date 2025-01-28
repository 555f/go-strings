package gostrings

type UUIDParser[T any] func(s string) (T, error)

func must[T any](t T, _ error) T {
	return t
}
