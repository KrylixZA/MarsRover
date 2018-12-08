package main

import (
	"github.com/MarsRover/facing"
	"github.com/MarsRover/move"
)

type marsRover struct {
	position position
	facing   facing.Direction
	world    world
}

func NewMarsRover(initialPosition position, facing facing.Direction, world world) *marsRover {
	return &marsRover{
		position: initialPosition,
		facing:   facing,
		world:    world,
	}
}

func (rover *marsRover) Move(movement move.Move) {
	switch movement {
	case move.Forward:
		if rover.facing == facing.North {
			rover.position.y++
		} else if rover.facing == facing.South {
			rover.position.y--
		} else if rover.facing == facing.East {
			rover.position.x++
		} else if rover.facing == facing.West {
			rover.position.x--
		}
	case move.Backward:
		if rover.facing == facing.North {
			rover.position.y--
		} else if rover.facing == facing.South {
			rover.position.y++
		} else if rover.facing == facing.East {
			rover.position.x--
		} else if rover.facing == facing.West {
			rover.position.x++
		}
	}
}
