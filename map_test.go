package gostrings

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/exp/constraints"
)

func TestSplitKeyValInt16(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[int16]{
		{
			name: "test int16",
			args: args{
				s: "a=12312;b=13312",
			},
			wantResult: map[string]int16{
				"a": 12312,
				"b": 13312,
			},
			wantErr: false,
		},
		{
			name: "test int16 failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test int16 failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValInt[int16](tt.args.s, ";", "=", 10, 64)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKeyValInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitKeyValInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplitKeyValInt8(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[int8]{
		{
			name: "test int8",
			args: args{
				s: "a=127;b=126",
			},
			wantResult: map[string]int8{
				"a": 127,
				"b": 126,
			},
			wantErr: false,
		},
		{
			name: "test int8 failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test int8 failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValInt[int8](tt.args.s, ";", "=", 10, 64)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKeyValInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitKeyValInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplitKeyValInt(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[V constraints.Integer] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[int]{
		{
			name: "test int",
			args: args{
				s: "a=8888888;b=555555",
			},
			wantResult: map[string]int{
				"a": 8888888,
				"b": 555555,
			},
			wantErr: false,
		},
		{
			name: "test int failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test int failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValInt[int](tt.args.s, ";", "=", 10, 64)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKeyValInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitKeyValInt() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplitKeyValFloat32(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[V constraints.Float] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[float32]{
		{
			name: "test int",
			args: args{
				s: "a=888.888;b=555.555",
			},
			wantResult: map[string]float32{
				"a": 888.888,
				"b": 555.555,
			},
			wantErr: false,
		},
		{
			name: "test int failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test int failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValFloat[float32](tt.args.s, ";", "=", 64)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKeyValFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitKeyValFloat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplitKeyValTime(t *testing.T) {
	type args struct {
		s      string
		layout string
	}
	type testCase[V time.Time] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[time.Time]{
		{
			name: "test time",
			args: args{
				s:      "a=2023-03-03T13:30:00Z;b=2023-03-03T13:30:01Z",
				layout: time.RFC3339,
			},
			wantResult: map[string]time.Time{
				"a": time.Date(2023, time.March, 3, 13, 30, 0, 0, time.UTC),
				"b": time.Date(2023, time.March, 3, 13, 30, 1, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "test time failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test time failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValTime[time.Time](tt.args.s, ";", "=", tt.args.layout)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitKeyValTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("SplitKeyValTime() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestSplitKeyValDuration(t *testing.T) {
	type args struct {
		s string
	}
	type testCase[V time.Duration] struct {
		name       string
		args       args
		wantResult map[string]V
		wantErr    bool
	}
	tests := []testCase[time.Duration]{
		{
			name: "test time",
			args: args{
				s: "a=1h1m1s;b=2h2m2s",
			},
			wantResult: map[string]time.Duration{
				"a": mustParseDuration(time.ParseDuration("1h1m1s")),
				"b": mustParseDuration(time.ParseDuration("2h2m2s")),
			},
			wantErr: false,
		},
		{
			name: "test time failed kv sep",
			args: args{
				s: "a=12312;b-13312",
			},

			wantErr: true,
		},
		{
			name: "test time failed item sep",
			args: args{
				s: "a=12312,b=13312",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := SplitKeyValDuration[time.Duration](tt.args.s, ";", "=")
			if (err != nil) != tt.wantErr {
				t.Errorf("TestSplitKeyValDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("TestSplitKeyValDuration() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func mustParseDuration(d time.Duration, err error) time.Duration {
	if err != nil {
		panic(err)
	}
	return d
}
