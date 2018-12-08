package main

import "github.com/MarsRover/facing"

type marsRover struct {
	position position
	facing   facing.Direction
	world    world
}

func NewMarsRover(initialPosition position, facing facing.Direction, world world) marsRover {
	return marsRover{
		position: initialPosition,
		facing:   facing,
		world:    world,
	}
}
