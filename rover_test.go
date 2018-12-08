package main

import (
	"github.com/MarsRover/facing"
	"github.com/MarsRover/move"
	"github.com/MarsRover/turn"
	"github.com/stretchr/testify/assert"
	"testing"
)

// constructor

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

// move forward 1 position

func TestMove_GivenSingleForwardMovementAndFacingNorthAtZeroPosition_ShouldIncrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleForwardMovementAndFacingSouthAtZeroPosition_ShouldDecrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleForwardMovementAndFacingEastAtZeroPosition_ShouldIncrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleForwardMovementAndFacingWestAtZeroPosition_ShouldDecrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// move backward 1 position

func TestMove_GivenTwoBackwardMovementsAndFacingNorthAtZeroPosition_ShouldDecrementYPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		y: -2,
	}

	// Act
	rover.Move([]move.Move{move.Backward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoBackwardMovementsAndFacingSouthAtZeroPosition_ShouldIncrementYPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		y: 2,
	}

	// Act
	rover.Move([]move.Move{move.Backward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoBackwardMovementsAndFacingEastAtZeroPosition_ShouldDecrementXPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		x: -2,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Backward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoBackwardMovementsAndFacingWestAtZeroPosition_ShouldIncrementXPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		x: 2,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Backward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// move forward multiple positions

func TestMove_GivenTwoForwardMovementsAndFacingNorthAtZeroPosition_ShouldIncrementYPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		y: 2,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoForwardMovementsAndFacingSouthAtZeroPosition_ShouldDecrementYPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		y: -2,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoForwardMovementsAndFacingEastAtZeroPosition_ShouldIncrementXPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		x: 2,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenTwoForwardMovementAndFacingWestAtZeroPosition_ShouldDecrementXPositionByTwoAndChangeNothingElse(t *testing.T) {
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
		x: -2,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Forward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// move backward multiple positions

func TestMove_GivenSingleBackwardMovementAndFacingNorthAtZeroPosition_ShouldDecrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleBackwardMovementAndFacingSouthAtZeroPosition_ShouldIncrementYPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleBackwardMovementAndFacingEastAtZeroPosition_ShouldDecrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenSingleBackwardMovementAndFacingWestAtZeroPosition_ShouldIncrementXPositionByOneAndChangeNothingElse(t *testing.T) {
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
	rover.Move([]move.Move{move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// move forward and backward

func TestMove_GivenForwardAndBackwardMovementsAndFacingNorthAtZeroPosition_ShouldEndUpAtZeroPositionAndChangeNothingElse(t *testing.T) {
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
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndBackwardMovementsAndFacingSouthAtZeroPosition_ShouldEndUpAtZeroPositionAndChangeNothingElse(t *testing.T) {
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
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndBackwardMovementsAndFacingEastAtZeroPosition_ShouldEndUpAtZeroPositionAndChangeNothingElse(t *testing.T) {
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
		x: 0,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestMove_GivenForwardAndBackwardMovementsAndFacingWestAtZeroPosition_ShouldEndUpAtZeroPositionAndChangeNothingElse(t *testing.T) {
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
		x: 0,
		y: 0,
	}

	// Act
	rover.Move([]move.Move{move.Forward, move.Backward})

	// Assert
	assert.Equal(t, expectedPosition, rover.position)
	assert.Equal(t, initialFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// turn left once

func TestTurn_GivenSingleLeftTurnAndFacingNorthAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleLeftTurnAndFacingSouthAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleLeftTurnAndFacingEastAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleLeftTurnAndFacingWestAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// turn left multiple times

func TestTurn_GivenTwoLeftTurnsAndFacingNorthAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.South

	// Act
	rover.Turn([]turn.Turn{turn.Left, turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoLeftTurnsAndFacingSouthAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.North

	// Act
	rover.Turn([]turn.Turn{turn.Left, turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoLeftTurnsAndFacingEastAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.West

	// Act
	rover.Turn([]turn.Turn{turn.Left, turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoLeftTurnsAndFacingWestAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.East

	// Act
	rover.Turn([]turn.Turn{turn.Left, turn.Left})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// turn right once

func TestTurn_GivenSingleRightTurnAndFacingNorthAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleRightTurnAndFacingSouthAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleRightTurnAndFacingEastAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenSingleRightTurnAndFacingWestAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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
	rover.Turn([]turn.Turn{turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

// turn right multiple times

func TestTurn_GivenTwoRightTurnsAndFacingNorthAtZeroPosition_ShouldFaceSouthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.South

	// Act
	rover.Turn([]turn.Turn{turn.Right, turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoRightTurnsAndFacingSouthAtZeroPosition_ShouldFaceNorthAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.North

	// Act
	rover.Turn([]turn.Turn{turn.Right, turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoRightTurnsAndFacingEastAtZeroPosition_ShouldFaceWestAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.West

	// Act
	rover.Turn([]turn.Turn{turn.Right, turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}

func TestTurn_GivenTwoRightTurnsAndFacingWestAtZeroPosition_ShouldFaceEastAndChangeNothingElse(t *testing.T) {
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

	expectedFacingDirection := facing.East

	// Act
	rover.Turn([]turn.Turn{turn.Right, turn.Right})

	// Assert
	assert.Equal(t, initialPosition, rover.position)
	assert.Equal(t, expectedFacingDirection, rover.facing)
	assert.Equal(t, world, rover.world)
}
