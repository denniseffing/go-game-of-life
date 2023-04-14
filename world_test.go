package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_world_aliveNeighborsOf(t *testing.T) {
	type fields struct {
		cells [][]cell
	}
	type args struct {
		point point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		{
			"aliveNeighborsOf should return 0 when all neighbors are dead",
			fields{
				[][]cell{
					{cell{false}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{false}},
				},
			},
			args{point{1, 1}},
			0,
		},
		{
			"aliveNeighborsOf should return 1 when upper left neighbor is alive",
			fields{
				[][]cell{
					{cell{true}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{false}},
				},
			},
			args{point{1, 1}},
			1,
		},
		{
			"aliveNeighborsOf should return 1 when lower right neighbor is alive",
			fields{
				[][]cell{
					{cell{false}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{true}},
				},
			},
			args{point{1, 1}},
			1,
		},
		{
			"aliveNeighborsOf should return 8 when all neighbors are alive",
			fields{
				[][]cell{
					{cell{true}, cell{true}, cell{true}},
					{cell{true}, cell{false}, cell{true}},
					{cell{true}, cell{true}, cell{true}},
				},
			},
			args{point{1, 1}},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := world{
				cells: tt.fields.cells,
			}
			if got := w.aliveNeighborsOf(tt.args.point); got != tt.want {
				t.Errorf("aliveNeighborsOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_world_age(t *testing.T) {
	type fields struct {
		cells [][]cell
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"given horizontal blinker, age should return vertical blinker",
			fields{
				[][]cell{
					{cell{false}, cell{true}, cell{false}},
					{cell{false}, cell{true}, cell{false}},
					{cell{false}, cell{true}, cell{false}},
				},
			},
			world{
				[][]cell{
					{cell{false}, cell{false}, cell{false}},
					{cell{true}, cell{true}, cell{true}},
					{cell{false}, cell{false}, cell{false}},
				},
			}.String(),
		},
		{
			"given first period of glider, age should return second period of glider",
			fields{
				[][]cell{
					{cell{false}, cell{true}, cell{false}, cell{false}},
					{cell{false}, cell{false}, cell{true}, cell{false}},
					{cell{true}, cell{true}, cell{true}, cell{false}},
					{cell{false}, cell{false}, cell{false}, cell{false}},
				},
			},
			world{
				[][]cell{
					{cell{false}, cell{false}, cell{false}, cell{false}},
					{cell{true}, cell{false}, cell{true}, cell{false}},
					{cell{false}, cell{true}, cell{true}, cell{false}},
					{cell{false}, cell{true}, cell{false}, cell{false}},
				},
			}.String(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := world{
				cells: tt.fields.cells,
			}

			assert.Equalf(t, tt.want, w.age().String(), "age()")
		})
	}
}
