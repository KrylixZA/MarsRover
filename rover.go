package main

import (
	"github.com/MarsRover/facing"
	moveEnum "github.com/MarsRover/move"
	turnEnum "github.com/MarsRover/turn"
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

func (rover *marsRover) Move(move moveEnum.Move) {
	switch move {
	case moveEnum.Forward:
		if rover.facing == facing.North {
			rover.position.y++
		} else if rover.facing == facing.South {
			rover.position.y--
		} else if rover.facing == facing.East {
			rover.position.x++
		} else if rover.facing == facing.West {
			rover.position.x--
		}
	case moveEnum.Backward:
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

func (rover *marsRover) Turn(turn turnEnum.Turn) {
	switch turn {
	case turnEnum.Left:
		if rover.facing == facing.North {
			rover.facing = facing.West
		} else if rover.facing == facing.South {
			rover.facing = facing.East
		} else if rover.facing == facing.East {
			rover.facing = facing.North
		} else if rover.facing == facing.West {
			rover.facing = facing.South
		}
	case turnEnum.Right:
		if rover.facing == facing.North {
			rover.facing = facing.East
		} else if rover.facing == facing.South {
			rover.facing = facing.West
		} else if rover.facing == facing.East {
			rover.facing = facing.South
		} else if rover.facing == facing.West {
			rover.facing = facing.North
		}
	}
}
