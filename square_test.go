package square

import (
	"reflect"
	"testing"
)

func TestSquare_Area(t *testing.T) {
	type fields struct {
		start Point
		a     uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{"side10", fields{Point{0, 0}, 10}, 100},
		{"side10", fields{Point{20, 20}, 10}, 100},
		{"side10", fields{Point{10, 50}, 10}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				start: tt.fields.start,
				a:     tt.fields.a,
			}
			if got := s.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_End(t *testing.T) {
	type fields struct {
		start Point
		a     uint
	}
	tests := []struct {
		name   string
		fields fields
		want   Point
	}{
		{"side10", fields{Point{0, 0}, 10}, Point{10, 10}},
		{"side10", fields{Point{20, 20}, 10}, Point{30, 30}},
		{"side10", fields{Point{10, 50}, 10}, Point{20, 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				start: tt.fields.start,
				a:     tt.fields.a,
			}
			if got := s.End(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("End() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare_Perimeter(t *testing.T) {
	type fields struct {
		start Point
		a     uint
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{"side10", fields{Point{0, 0}, 10}, 40},
		{"side10", fields{Point{20, 20}, 10}, 40},
		{"side10", fields{Point{10, 50}, 5}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Square{
				start: tt.fields.start,
				a:     tt.fields.a,
			}
			if got := s.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
