package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_point_neighbors(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	tests := []struct {
		name      string
		fields    fields
		contained point
	}{
		{
			"neighbors should return array containing upper left neighbor",
			fields{1, 1},
			point{0, 0},
		},
		{
			"neighbors should return array containing upper neighbor",
			fields{1, 1},
			point{1, 0},
		},
		{
			"neighbors should return array containing upper right neighbor",
			fields{1, 1},
			point{2, 0},
		},
		{
			"neighbors should return array containing left neighbor",
			fields{1, 1},
			point{0, 1},
		},
		{
			"neighbors should return array containing right neighbor",
			fields{1, 1},
			point{2, 1},
		},
		{
			"neighbors should return array containing lower left neighbor",
			fields{1, 1},
			point{0, 2},
		},
		{
			"neighbors should return array containing lower neighbor",
			fields{1, 1},
			point{1, 2},
		},
		{
			"neighbors should return array containing lower right neighbor",
			fields{1, 1},
			point{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := point{
				x: tt.fields.x,
				y: tt.fields.y,
			}

			got := p.neighbors()
			assert.Contains(t, got, tt.contained)
		})
	}
}
