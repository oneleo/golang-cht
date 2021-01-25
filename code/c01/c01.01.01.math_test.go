package c01

import (
	"fmt"
	"testing"
)

func TestAddMinusMultiply(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name         string
		args         args
		wantAdd      int
		wantMinus    int
		wantMultiply int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1: a = 1; b = 1",
			args: args{
				a: 1,
				b: 1,
			},
			wantAdd:      2,
			wantMinus:    0,
			wantMultiply: 1,
		},
		{
			name: "Test 2: a = 2; b = 1",
			args: args{
				a: 2,
				b: 1,
			},
			wantAdd:      3,
			wantMinus:    1,
			wantMultiply: 2,
		},
		{
			name: "Test 3: a = 1; b = 2",
			args: args{
				a: 1,
				b: 2,
			},
			wantAdd:      3,
			wantMinus:    -1,
			wantMultiply: 2,
		},
		{
			name: "Test 4: a = 2; b = 2",
			args: args{
				a: 2,
				b: 2,
			},
			wantAdd:      4,
			wantMinus:    0,
			wantMultiply: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdd, gotMinus, gotMultiply := AddMinusMultiply(tt.args.a, tt.args.b)
			if gotAdd != tt.wantAdd {
				t.Errorf("AddMinusMultiply() gotAdd = %v, want %v", gotAdd, tt.wantAdd)
			}
			if gotMinus != tt.wantMinus {
				t.Errorf("AddMinusMultiply() gotMinus = %v, want %v", gotMinus, tt.wantMinus)
			}
			if gotMultiply != tt.wantMultiply {
				t.Errorf("AddMinusMultiply() gotMultiply = %v, want %v", gotMultiply, tt.wantMultiply)
			}
		})
	}
}

func ExampleAddMinusMultiply() {
	a := 1
	b := 1
	out1, out2, out3 := AddMinusMultiply(a, b)
	fmt.Println(out1, out2, out3)
	// Output:
	// 2 0 1
}

func TestSum(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		// TODO: Add test cases.
		{
			name: "Test 1: 1 + 2",
			args: args{
				in: []int{1, 2},
			},
			wantSum: 3,
		},
		{
			name: "Test 2: 1 + 2 + 3",
			args: args{
				in: []int{1, 2, 3},
			},
			wantSum: 6,
		},
		{
			name: "Test 3: 1 + 2 + 3 + 4",
			args: args{
				in: []int{1, 2, 3, 4},
			},
			wantSum: 10,
		},
		{
			name: "Test 4: 1 + 2 + 3 + 4 + 5",
			args: args{
				in: []int{1, 2, 3, 4, 5},
			},
			wantSum: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := Sum(tt.args.in...); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func ExampleSum() {
	in := []int{5, 4, 3, 2, 1}
	out := Sum(in...)
	fmt.Println(out)
	// Output:
	// 15
}
