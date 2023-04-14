package main

import (
	"reflect"
	"testing"
)

func Test_cell_age(t *testing.T) {
	type fields struct {
		alive bool
	}
	type args struct {
		aliveNeighbors uint8
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   cell
	}{
		// TODO: Add test cases.
		{
			"given alive cell, age should return dead cell when no neighbors are alive",
			fields{true},
			args{0},
			cell{false},
		},
		{
			"given alive cell, age should return dead cell when one neighbor is alive",
			fields{true},
			args{1},
			cell{false},
		},
		{
			"given alive cell, age should return dead cell when four neighbors are alive",
			fields{true},
			args{4},
			cell{false},
		},
		{
			"given alive cell, age should return dead cell when more than four neighbors are alive",
			fields{true},
			args{5},
			cell{false},
		},
		{
			"given alive cell, age should return alive cell when two neighbors are alive",
			fields{true},
			args{2},
			cell{true},
		},
		{
			"given alive cell, age should return alive cell when three neighbors are alive",
			fields{true},
			args{3},
			cell{true},
		},
		{
			"given dead cell, age should return alive cell when three neighbors are alive",
			fields{false},
			args{3},
			cell{true},
		},
		{
			"given dead cell, age should return dead cell when one neighbor is alive",
			fields{false},
			args{1},
			cell{false},
		},
		{
			"given dead cell, age should return dead cell when two neighbors are alive",
			fields{false},
			args{2},
			cell{false},
		},
		{
			"given dead cell, age should return dead cell when four neighbors are alive",
			fields{false},
			args{4},
			cell{false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cell{
				alive: tt.fields.alive,
			}
			if got := c.age(tt.args.aliveNeighbors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("age() = %v, want %v", got, tt.want)
			}
		})
	}
}
