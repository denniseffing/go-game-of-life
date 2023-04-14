package main

type cell struct {
	alive bool
}

func (c cell) age(aliveNeighbors uint8) cell {
	if c.alive && aliveNeighbors == 2 {
		return cell{true}
	}
	if aliveNeighbors == 3 {
		return cell{true}
	}
	return cell{false}
}
