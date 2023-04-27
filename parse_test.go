package gostrings

import (
	"testing"

	"golang.org/x/exp/constraints"
)

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
			got, err := ParseInt[int](tt.args.s, tt.args.base, tt.args.bitSize)
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
			got, err := ParseInt[int64](tt.args.s, tt.args.base, tt.args.bitSize)
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
			name: "test float32",
			args: args{
				s:       "10.1",
				bitSize: 64,
			},
			want:    10.1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFloat[float32](tt.args.s, tt.args.bitSize)
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
