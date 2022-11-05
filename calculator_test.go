package calculator

import (
	"reflect"
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	type fields struct {
		Result int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calculator
	}{
		{
			name:   "Test addition",
			fields: fields{Result: 0},
			args:   args{10},
			want:   &Calculator{Result: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Result: tt.fields.Result,
			}
			if got := c.Add(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_Calculate(t *testing.T) {
	type fields struct {
		Result int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Test calculate",
			fields: fields{5},
			want:   5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Result: tt.fields.Result,
			}
			if got := c.Calculate(); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_Div(t *testing.T) {
	type fields struct {
		Result int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calculator
	}{
		{
			name:   "Test division",
			fields: fields{4},
			args:   args{2},
			want:   &Calculator{Result: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Result: tt.fields.Result,
			}
			if got := c.Div(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_Mul(t *testing.T) {
	type fields struct {
		Result int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calculator
	}{
		{
			name:   "Test multiplication",
			fields: fields{4},
			args:   args{2},
			want:   &Calculator{Result: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Result: tt.fields.Result,
			}
			if got := c.Mul(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_Sub(t *testing.T) {
	type fields struct {
		Result int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Calculator
	}{
		{
			name:   "Test subtraction",
			fields: fields{4},
			args:   args{3},
			want:   &Calculator{Result: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				Result: tt.fields.Result,
			}
			if got := c.Sub(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Calculator
	}{
		{
			name: "Test new calculator",
			want: &Calculator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
