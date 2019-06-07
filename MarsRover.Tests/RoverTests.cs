using NUnit.Framework;

namespace MarsRover.Tests
{
    [TestFixture]
    public class RoverTests
    {
        [Test]
        public void TestConstructor_CallingDefault_ShouldCreateRoverAtZeroZero()
        {
            // Arrange & Act
            var rover = new Rover();

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
        }

        [Test]
        public void TestConstructor_GivenXPosition_ShouldCreateRoverAtXPositionAndZeroYPosition()
        {
            // Arrange
            var xInitPos = 70;

            // Act
            var rover = new Rover(xInitPos, 0);

            // Assert
            Assert.AreEqual(xInitPos, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
        }

        [Test]
        public void TestConstructor_GivenYPosition_ShouldCreateRoverAtZeroYPositionAndYPosition()
        {
            // Arrange
            var yInitPos = 53;

            // Act
            var rover = new Rover(0, yInitPos);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(yInitPos, rover.Position.Y);
        }

        [Test]
        public void TestConstructor_GivenSomePositionFacingEast_ShouldCreateRoverFacingEast()
        {
            // Arrange
            var pos = new Position(70, 53);
            var direction = Direction.East;

            // Act
            var rover = new Rover(pos, direction);

            // Assert
            Assert.AreEqual(70, rover.Position.X);
            Assert.AreEqual(53, rover.Position.Y);
            Assert.AreEqual(Direction.East, rover.Direction);
        }

        [Test]
        public void TestConstructor_GivenSomePointFacingWest_ShouldCreateRoverFacingWest()
        {
            // Arrange
            var pos = new Position(70, 53);
            var direction = Direction.West;

            // Act
            var rover = new Rover(pos, direction);

            // Assert
            Assert.AreEqual(70, rover.Position.X);
            Assert.AreEqual(53, rover.Position.Y);
            Assert.AreEqual(Direction.West, rover.Direction);
        }

        [Test]
        public void TestConstructor_GivenSomePointFacingNorth_ShouldCreateRoverFacingNorth()
        {
            // Arrange
            var pos = new Position(70, 53);
            var direction = Direction.North;

            // Act
            var rover = new Rover(pos, direction);

            // Assert
            Assert.AreEqual(70, rover.Position.X);
            Assert.AreEqual(53, rover.Position.Y);
            Assert.AreEqual(Direction.North, rover.Direction);
        }

        [Test]
        public void TestConstructor_GivenSomePointFacingSouth_ShouldCreateRoverFacingSouth()
        {
            // Arrange
            var pos = new Position(70, 53);
            var direction = Direction.South;

            // Act
            var rover = new Rover(pos, direction);

            // Assert
            Assert.AreEqual(70, rover.Position.X);
            Assert.AreEqual(53, rover.Position.Y);
            Assert.AreEqual(Direction.South, rover.Direction);
        }

        [Test]
        public void TestConstructor_GivenWorldDimensions_ShouldCreateRoverInWorld()
        {
            // Arrange
            var pos = new Position(0, 0);
            var direction = Direction.East;
            var world = new World(0, 1, 0, 1);

            // Act
            var rover = new Rover(pos, direction, world);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
            Assert.AreEqual(Direction.East, rover.Direction);
            Assert.AreEqual(0, rover.World.XMin);
            Assert.AreEqual(1, rover.World.XMax);
            Assert.AreEqual(0, rover.World.YMin);
            Assert.AreEqual(1, rover.World.YMax);
        }

        [Test]
        public void TestMove_GivenMoveForwardOneUnitFacingEast_ShouldMoveToOneZero()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.East;
            var world = new World(-1, 1, -1, 1);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(1, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
            Assert.AreEqual(Direction.East, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardOneUnitFacingWest_ShouldMoveToNegativeOneZero()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.West;
            var world = new World(-1, 1, -1, 1);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(-1, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
            Assert.AreEqual(Direction.West, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardOneUnitFacingNorth_ShouldMoveToZeroOne()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.North;
            var world = new World(-1, 1, -1, 1);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(1, rover.Position.Y);
            Assert.AreEqual(Direction.North, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardOneUnitFacingSouth_ShouldMoveToZeroNegativeOne()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.South;
            var world = new World(-1, 1, -1, 1);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(-1, rover.Position.Y);
            Assert.AreEqual(Direction.South, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoreForwardNManyUnitsFacingEast_ShouldMoveToNZero()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.East;
            var world = new World(-20, 20, -20, 20);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(11);

            // Assert
            Assert.AreEqual(11, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
            Assert.AreEqual(Direction.East, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardNManyUnitsFacingWest_ShouldMoveToNegativeNZero()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.West;
            var world = new World(-20, 20, -20, 20);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(11);

            // Assert
            Assert.AreEqual(-11, rover.Position.X);
            Assert.AreEqual(0, rover.Position.Y);
            Assert.AreEqual(Direction.West, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardNManyUnitsFacingNorth_ShouldMoveToZeroN()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.North;
            var world = new World(-2, 20, -20, 20);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(14);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(14, rover.Position.Y);
            Assert.AreEqual(Direction.North, rover.Direction);
        }

        [Test]
        public void TestMove_GivenMoveForwardNManyUnitsFacingSouth_ShouldMoveToZeroNegativeN()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.South;
            var world = new World(-20, 20, -20, 20);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(14);

            // Assert
            Assert.AreEqual(0, rover.Position.X);
            Assert.AreEqual(-14, rover.Position.Y);
            Assert.AreEqual(Direction.South, rover.Direction);
        }

        [Test]
        public void TestMove_GiveOnEastEdgeFacingEastAndMoveOneStep_ShouldWrapToWestSide()
        {
            // Arrange
            var initPosition = new Position(10, 0);
            var direction = Direction.East;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(-9, rover.Position.X);
        }

        [Test]
        public void TestMove_GiveOnEastEdgeFacingEastAndMoveThreeSteps_ShouldWrapToWestSide()
        {
            // Arrange
            var initPosition = new Position(10, 0);
            var direction = Direction.East;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(3);

            // Assert
            Assert.AreEqual(-7, rover.Position.X);
        }

        [Test]
        public void TestMove_GiveOnWestEdgeFacingWestAndMoveOneStep_ShouldWrapToEastSide()
        {
            // Arrange
            var initPosition = new Position(-10, 0);
            var direction = Direction.West;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(9, rover.Position.X);
        }

        [Test]
        public void TestMove_GiveOnWestEdgeFacingWestAndMoveThreeSteps_ShouldWrapToEastSide()
        {
            // Arrange
            var initPosition = new Position(-10, 0);
            var direction = Direction.West;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(3);

            // Assert
            Assert.AreEqual(7, rover.Position.X);
        }

        [Test]
        public void TestMove_GiveOnNorthEdgeFacingNorthAndMoveOneStep_ShouldWrapToSouthSide()
        {
            // Arrange
            var initPosition = new Position(0, 10);
            var direction = Direction.North;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(-9, rover.Position.Y);
        }

        [Test]
        public void TestMove_GiveOnNorthEdgeFacingWestAndMoveThreeSteps_ShouldWrapToSouthSide()
        {
            // Arrange
            var initPosition = new Position(0, 10);
            var direction = Direction.North;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(3);

            // Assert
            Assert.AreEqual(-7, rover.Position.Y);
        }

        [Test]
        public void TestMove_GiveOnSouthhEdgeFacingSouthAndMoveOneStep_ShouldWrapToNorthSide()
        {
            // Arrange
            var initPosition = new Position(0, -10);
            var direction = Direction.South;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(1);

            // Assert
            Assert.AreEqual(9, rover.Position.Y);
        }

        [Test]
        public void TestMove_GiveOnSouthEdgeFacingWestAndMoveThreeSteps_ShouldWrapToNorthSide()
        {
            // Arrange
            var initPosition = new Position(0, -10);
            var direction = Direction.South;
            var world = new World(-10, 10, -10, 10);
            var rover = new Rover(initPosition, direction, world);

            // Act
            rover.Move(3);

            // Assert
            Assert.AreEqual(7, rover.Position.Y);
        }

        [Test]
        public void TestTurnLeft_GivenFacingEast_ShouldFaceNorth()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.East;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnLeft();

            // Assert
            Assert.AreEqual(Direction.North, rover.Direction);
        }

        [Test]
        public void TestTurnLeft_GivenFacingNorth_ShouldFaceWest()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.North;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnLeft();

            // Assert
            Assert.AreEqual(Direction.West, rover.Direction);
        }

        [Test]
        public void TestTurnLeft_GivenFacingWest_ShouldFaceSouth()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.West;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnLeft();

            // Assert
            Assert.AreEqual(Direction.South, rover.Direction);
        }

        [Test]
        public void TestTurnLeft_GivenFacingSouth_ShouldFaceEast()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.South;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnLeft();

            // Assert
            Assert.AreEqual(Direction.East, rover.Direction);
        }

        [Test]
        public void TestTurnRight_GivenFacingEast_ShouldFaceSouth()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.East;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnRight();

            // Assert
            Assert.AreEqual(Direction.South, rover.Direction);
        }

        [Test]
        public void TestTurnRight_GivenFacingSouth_ShouldFaceWest()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.South;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnRight();

            // Assert
            Assert.AreEqual(Direction.West, rover.Direction);
        }

        [Test]
        public void TestTurnRight_GivenFacingWest_ShouldFaceNorth()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.West;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnRight();

            // Assert
            Assert.AreEqual(Direction.North, rover.Direction);
        }

        [Test]
        public void TestTurnRight_GivenFacingNorth_ShouldFaceEast()
        {
            // Arrange
            var initPosition = new Position(0, 0);
            var direction = Direction.North;
            var rover = new Rover(initPosition, direction);

            // Act
            rover.TurnRight();

            // Assert
            Assert.AreEqual(Direction.East, rover.Direction);
        }
    }
}
