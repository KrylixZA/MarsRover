package main

import (
	"github.com/MarsRover/facing"
	"github.com/MarsRover/move"
	"github.com/MarsRover/turn"
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

func TestMove_GivenForwardAndFacingNorthAtZeroPosition_ShouldIncrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndFacingSouthAtZeroPosition_ShouldDecrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndFacingEastAtZeroPosition_ShouldIncrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndFacingWestAtZeroPosition_ShouldDecrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenBackwardAndFacingNorthAtZeroPosition_ShouldDecrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenBackwardAndFacingSouthAtZeroPosition_ShouldIncrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenBackwardAndFacingEastAtZeroPosition_ShouldDecrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenBackwardAndFacingWestAtZeroPosition_ShouldIncrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenLeftAndFacingNorthAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.West

	// Act
	rover.Turn(turn.Left)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenLeftAndFacingSouthAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.East

	// Act
	rover.Turn(turn.Left)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenLeftAndFacingEastAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.North

	// Act
	rover.Turn(turn.Left)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenLeftAndFacingWestAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.South

	// Act
	rover.Turn(turn.Left)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenRightAndFacingNorthAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.East

	// Act
	rover.Turn(turn.Right)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenRightAndFacingSouthAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.West

	// Act
	rover.Turn(turn.Right)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenRightAndFacingEastAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.South

	// Act
	rover.Turn(turn.Right)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenRightAndFacingWestAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.North

	// Act
	rover.Turn(turn.Right)

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}
