package calculator

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateCacheKey(t *testing.T) {
	type args struct {
		operation Operation
		x         float64
		y         float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it should return add_2_1 for the key in case Y is bigger than X",
			args: args{
				operation: add,
				x:         1.1,
				y:         2.2,
			},
			want: "add_2.2_1.1",
		},
		{
			name: "it should return add_2_1 for the key in case X is bigger than Y",
			args: args{
				operation: add,
				x:         2,
				y:         1,
			},
			want: "add_2_1",
		},
		{
			name: "it should return multiply_2_1 for the key in case Y is bigger than X",
			args: args{
				operation: multiply,
				x:         1.1,
				y:         2.2,
			},
			want: "multiply_2.2_1.1",
		},
		{
			name: "it should return multiply_2_1 for the key in case X is bigger than Y",
			args: args{
				operation: multiply,
				x:         2,
				y:         1,
			},
			want: "multiply_2_1",
		},
		{
			name: "it should return divide_2_1 for divide",
			args: args{
				operation: divide,
				x:         1.1,
				y:         2.2,
			},
			want: "divide_1.1_2.2",
		},
		{
			name: "it should return subtract_2_1 for subtract",
			args: args{
				operation: subtract,
				x:         2,
				y:         1,
			},
			want: "subtract_2_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateCacheKey(tt.args.operation, tt.args.x, tt.args.y)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_calculateResult(t *testing.T) {
	type args struct {
		x      float64
		y      float64
		action Operation
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "it should return expected result for add operation",
			args: args{
				x:      1.2,
				y:      3.333,
				action: add,
			},
			want: 4.533,
		},
		{
			name: "it should return expected result for add operation",
			args: args{
				x:      1,
				y:      3,
				action: add,
			},
			want: 4,
		},
		{
			name: "it should return expected result for add operation",
			args: args{
				x:      1.2,
				y:      3.333,
				action: add,
			},
			want: 4.533,
		},
		{
			name: "it should return expected result for multiply operation",
			args: args{
				x:      1,
				y:      3,
				action: multiply,
			},
			want: 3,
		},
		{
			name: "it should return expected result for divide divide",
			args: args{
				x:      5,
				y:      2,
				action: divide,
			},
			want: 2.5,
		},
		{
			name: "it should return expected result for subtract operation",
			args: args{
				x:      1,
				y:      3,
				action: subtract,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateResult(tt.args.x, tt.args.y, tt.args.action)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_readQueryParams(t *testing.T) {
	type args struct {
		v url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		want1   float64
		wantErr bool
		errMsg  string
	}{
		{
			name: "it should return read and return query params",
			args: args{
				v: url.Values{
					"x": []string{"1"},
					"y": []string{"2"},
				},
			},
			want:  1,
			want1: 2,
		},
		{
			name: "it should return read and return query params with float params",
			args: args{
				v: url.Values{
					"x": []string{"1.2345"},
					"y": []string{"2.3456"},
				},
			},
			want:  1.2345,
			want1: 2.3456,
		},
		{
			name: "it should return read and return an error if X query is not valid",
			args: args{
				v: url.Values{
					"x": []string{"AAA"},
					"y": []string{"2.3456"},
				},
			},
			wantErr: true,
			errMsg:  "strconv.ParseFloat: parsing \"AAA\": invalid syntax",
		},
		{
			name: "it should return read and return an error if Y query is not valid",
			args: args{
				v: url.Values{
					"x": []string{"1.2345"},
					"y": []string{"AAA"},
				},
			},
			wantErr: true,
			errMsg:  "strconv.ParseFloat: parsing \"AAA\": invalid syntax",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := readQueryParams(tt.args.v)
			if tt.wantErr {
				require.EqualError(t, err, tt.errMsg)
			}
			require.Equal(t, tt.want, got)
			require.Equal(t, tt.want1, got1)
		})
	}
}
