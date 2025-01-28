package gostrings

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestParseBool(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[T constraints.Signed] struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "parse bool success",
			args: args{
				s: "true",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "parse bool error",
			args: args{
				s: "true2",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool
			err := ParseBool(tt.args.s, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt(t *testing.T) {
	type args struct {
		s       string
		base    int
		bitSize int
	}
	type testCase[T constraints.Signed] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{
			name: "test int",
			args: args{
				s:       "10",
				base:    10,
				bitSize: 64,
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int
			err := ParseInt(tt.args.s, tt.args.base, tt.args.bitSize, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseInt64(t *testing.T) {
	type args struct {
		s       string
		base    int
		bitSize int
	}
	type testCase[T constraints.Signed] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int64]{
		{
			name: "test int64",
			args: args{
				s:       "10",
				base:    10,
				bitSize: 64,
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int64
			err := ParseInt[int64](tt.args.s, tt.args.base, tt.args.bitSize, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFloat(t *testing.T) {
	type args struct {
		s       string
		bitSize int
	}
	type testCase[T constraints.Float] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[float32]{
		{
			name: "parse float32 success",
			args: args{
				s:       "10.1",
				bitSize: 64,
			},
			want:    10.1,
			wantErr: false,
		},
		{
			name: "parse float32 error",
			args: args{
				s:       "qw",
				bitSize: 64,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got float32
			err := ParseFloat[float32](tt.args.s, tt.args.bitSize, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseComplex(t *testing.T) {
	type args struct {
		s       string
		bitSize int
	}
	type testCase[T constraints.Complex] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[complex64]{
		{
			name: "parse complex64 success",
			args: args{
				s:       "10.1",
				bitSize: 64,
			},
			want:    10.1,
			wantErr: false,
		},
		{
			name: "parse complex64 error",
			args: args{
				s:       "qw",
				bitSize: 64,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got complex64
			err := ParseComplex[complex64](tt.args.s, tt.args.bitSize, &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
