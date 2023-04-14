package main

import "strings"

type world struct {
	cells [][]cell
}

func (w world) aliveNeighborsOf(point point) uint8 {
	var sum uint8
	for _, neighbor := range point.neighbors() {
		if w.pointInBounds(neighbor) && w.cells[neighbor.y][neighbor.x].alive {
			sum++
		}
	}
	return sum
}

func (w world) pointInBounds(point point) bool {
	return !(point.y < 0 || point.y >= len(w.cells) || point.x < 0 || point.x >= len(w.cells[point.y]))
}

func (w world) age() world {
	slice := make([][]cell, len(w.cells))
	for i := range slice {
		slice[i] = make([]cell, len(w.cells[0]))
	}

	for y, row := range w.cells {
		for x, cell := range row {
			slice[y][x] = cell.age(w.aliveNeighborsOf(point{x, y}))
		}
	}
	return world{slice}
}

func (w world) String() string {
	var sb strings.Builder
	for y := range w.cells {
		for x := range w.cells[y] {
			sb.WriteString(w.cells[y][x].String())
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
