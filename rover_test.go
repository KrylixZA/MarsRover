package main

import (
	"github.com/MarsRover/facing"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRover_GivenZeroPositionFacingNorthWithWorldDimensions_ShouldCreateNewRover(t *testing.T) {
	// Arrange
	initialPosition := position{
		x: 0,
		y: 0,
	}
	initialFacingDirection := facing.North
	world := world{
		length:  50,
		breadth: 50,
	}

	// Act
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestNewRover_GivenZeroPositionFacingSouthWithWorldDimensions_ShouldCreateNewRover(t *testing.T) {
	// Arrange
	initialPosition := position{
		x: 0,
		y: 0,
	}
	initialFacingDirection := facing.South
	world := world{
		length:  50,
		breadth: 50,
	}

	// Act
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestNewRover_GivenZeroPositionFacingEastWithWorldDimensions_ShouldCreateNewRover(t *testing.T) {
	// Arrange
	initialPosition := position{
		x: 0,
		y: 0,
	}
	initialFacingDirection := facing.East
	world := world{
		length:  50,
		breadth: 50,
	}

	// Act
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestNewRover_GivenZeroPositionFacingWestWithWorldDimensions_ShouldCreateNewRover(t *testing.T) {
	// Arrange
	initialPosition := position{
		x: 0,
		y: 0,
	}
	initialFacingDirection := facing.West
	world := world{
		length:  50,
		breadth: 50,
	}

	// Act
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}
