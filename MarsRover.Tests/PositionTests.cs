using NUnit.Framework;

namespace MarsRover.Tests
{
    [TestFixture]
    public class PositionTests
    {
        [Test]
        public void TestConstructor_CallingDefault_ShouldReturnExpectedPosition()
        {
            // Arrange & Act
            var position = new Position();

            // Assert
            Assert.AreEqual(0, position.X);
            Assert.AreEqual(0, position.Y);
        }

        [Test]
        public void TestConstructor_GivenXValue_ShouldCreatePositionAtXAndZero()
        {
            // Arrange
            var x = 15;

            // Act
            var position = new Position(x, 0);

            // Assert
            Assert.AreEqual(x, position.X);
            Assert.AreEqual(0, position.Y);
        }

        [Test]
        public void TestConstructor_GivenYValue_ShouldCreatePositionAtZeroAndY()
        {
            // Arrange
            var y = 20;

            // Act
            var position = new Position(0, y);

            // Assert
            Assert.AreEqual(0, position.X);
            Assert.AreEqual(y, position.Y);
        }

        [Test]
        public void TestConstructor_GivenXAndYValue_ShouldCreatePositionAtXAndY()
        {
            // Arrange
            var x = 15;
            var y = 20;

            // Act
            var position = new Position(x, y);

            // Assert
            Assert.AreEqual(x, position.X);
            Assert.AreEqual(y, position.Y);
        }
    }
}
