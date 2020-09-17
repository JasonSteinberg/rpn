package main

import (
	"reflect"
	"testing"
)

func TestStackCalc_Push(t *testing.T) {
	dEmpty := []numOp{}
	d1 := []numOp{1, 2}

	type fields struct {
		data []numOp
	}
	type args struct {
		a numOp
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    error
		wantResult int
	}{
		{
			name:       "Can store a value",
			fields:     fields{dEmpty},
			args:       args{1},
			wantErr:    nil,
			wantResult: 1,
		},
		{
			name:       "Can store another value",
			fields:     fields{d1},
			args:       args{44},
			wantErr:    nil,
			wantResult: 44,
		},
		{
			name:       "Can reduce value on operation",
			fields:     fields{d1},
			args:       args{"+"},
			wantErr:    nil,
			wantResult: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StackCalc{
				data: tt.fields.data,
			}
			err := r.Push(tt.args.a)
			if err != tt.wantErr || r.Value() != tt.wantResult {
				t.Errorf("Push() error = %v, wantErr %v, wantResult %v/%v", err, tt.wantErr, tt.wantResult, r.Value())
			}
		})
	}
}

func TestStackCalc_Value(t *testing.T) {
	d1 := []numOp{3, 4}
	d2 := []numOp{5, 6, 7, 8}

	type fields struct {
		data []numOp
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"Normal", fields{d1}, 4},
		{"Normal", fields{d2}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StackCalc{
				data: tt.fields.data,
			}
			if got := r.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackCalc_Length(t *testing.T) {
	d1 := []numOp{3, 4}
	d2 := []numOp{5, 6, 7, 8}
	type fields struct {
		data []numOp
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"2 Values", fields{d1}, 2},
		{"4 Values", fields{d2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StackCalc{
				data: tt.fields.data,
			}
			if got := r.Length(); got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackCalc_reduce(t *testing.T) {
	tests := []struct {
		name   string
		data   StackCalc
		op     string
		result int
	}{
		{"Addition", StackCalc{[]numOp{3, 5}}, "+", 8},
		{"Division", StackCalc{[]numOp{16, 4}}, "/", 4},
		{"Multiplication", StackCalc{[]numOp{4, 4}}, "*", 16},
		{"Subtraction", StackCalc{[]numOp{8, 9}}, "-", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &StackCalc{
				data: tt.data.data,
			}
			r.reduce(tt.op)
			if got := r.Value(); got != tt.result {
				t.Errorf("reduce for %s = %v, want %v", tt.name, got, tt.result)
			}
		})
	}

}

func Test_doOperation(t *testing.T) {
	type args struct {
		x  int
		y  int
		op string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"Add values", args{1, 2, "+"}, 3, true},
		{"Divide values", args{16, 4, "/"}, 4, true},
		{"Subtract values", args{4, 3, "-"}, 1, true},
		{"Multiple values", args{5, 2, "*"}, 10, true},
		{"Bad operation", args{5, 2, "i"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := doOperation(tt.args.x, tt.args.y, tt.args.op)
			if got != tt.want {
				t.Errorf("doOperation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("doOperation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_isMathOp(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Let's make sure math works
		{"Addition", args{"+"}, true},
		{"Minus", args{"-"}, true},
		{"Division", args{"/"}, true},
		{"Multiplication", args{"*"}, true},

		// Bad data
		{"Bad math combo", args{"+-"}, false},
		{"Bad number", args{"1"}, false},
		{"Bad letter", args{"a"}, false},
		{"Bad letters", args{"adf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMathOp(tt.args.a); got != tt.want {
				t.Errorf("isMathOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeStackCalc(t *testing.T) {
	type args struct {
		tmpStc          *StackCalc
		computationList string
	}
	tests := []struct {
		name    string
		args    args
		want    *StackCalc
		wantErr bool
	}{
		{name: "Provided Test Case", args: args{&StackCalc{}, "1 2 + 3 3 + +"}, want: &StackCalc{data: []numOp{9}}, wantErr: false},
		{name: "Double digits", args: args{&StackCalc{}, "14 2 / 1 + 3 3 + +"}, want: &StackCalc{data: []numOp{14}}, wantErr: false},
		{name: "All operations", args: args{&StackCalc{}, "4 2 + 3 / 4 * 1 -"}, want: &StackCalc{data: []numOp{7}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ComputeStackCalc(tt.args.tmpStc, tt.args.computationList)
			if (err != nil) != tt.wantErr {
				t.Errorf("ComputeStackCalc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComputeStackCalc() got = %v, want %v", got, tt.want)
			}
		})

	}
}
