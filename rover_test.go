package main

import (
	"github.com/MarsRover/facing"
	"github.com/MarsRover/move"
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

func TestMove_GivenForwardAndFacingNorthAtZeroPosition_ShouldIncrementYPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 0,
		y: 1,
	}

	// Act
	rover.Move(move.Forward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenForwardAndFacingSouthAtZeroPosition_ShouldDecrementYPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 0,
		y: -1,
	}

	// Act
	rover.Move(move.Forward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenForwardAndFacingEastAtZeroPosition_ShouldIncrementXPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 1,
		y: 0,
	}

	// Act
	rover.Move(move.Forward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenForwardAndFacingWestAtZeroPosition_ShouldDecrementXPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: -1,
		y: 0,
	}

	// Act
	rover.Move(move.Forward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenBackwardAndFacingNorthAtZeroPosition_ShouldDecrementYPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 0,
		y: -1,
	}

	// Act
	rover.Move(move.Backward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenBackwardAndFacingSouthAtZeroPosition_ShouldIncrementYPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 0,
		y: 1,
	}

	// Act
	rover.Move(move.Backward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenBackwardAndFacingEastAtZeroPosition_ShouldDecrementXPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: -1,
		y: 0,
	}

	// Act
	rover.Move(move.Backward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}

func TestMove_GivenBackwardAndFacingWestAtZeroPosition_ShouldIncrementXPositionByOne(t *testing.T) {
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
	rover := NewMarsRover(initialPosition, initialFacingDirection, world)

	expectedPosition := position{
		x: 1,
		y: 0,
	}

	// Act
	rover.Move(move.Backward)

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
}
