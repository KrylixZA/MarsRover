using NUnit.Framework;

namespace MarsRover.Tests
{
    [TestFixture]
    public class WorldTests
    {
        [Test]
        public void TestConstructor_CallingDefault_ShouldCreateWorldWith1x1Size()
        {
            // Arrange & Act
            var world = new World();

            // Assert
            Assert.AreEqual(0, world.XMin);
            Assert.AreEqual(1, world.XMax);
            Assert.AreEqual(0, world.YMin);
            Assert.AreEqual(1, world.YMax);
        }

        [Test]
        public void TestConstructor_GivenMinAndMaxXValues_ShouldCreateWorld()
        {
            // Arrange
            var xMin = -50;
            var xMax = 50;

            // Act
            var world = new World(xMin, xMax, 0, 1);

            // Assert
            Assert.AreEqual(xMin, world.XMin);
            Assert.AreEqual(xMax, world.XMax);
        }

        [Test]
        public void TestConstructor_GivenMinAndMaxYValues_ShouldCreateWorld()
        {
            // Arrange
            var yMin = -50;
            var yMax = 50;

            // Act
            var world = new World(0, 1, yMin, yMax);

            // Assert
            Assert.AreEqual(yMin, world.YMin);
            Assert.AreEqual(yMax, world.YMax);
        }

        [Test]
        public void TestConstructor_GivenValuesForXAndY_ShouldCreateWorld()
        {
            // Arrange
            var xMin = -50;
            var xMax = 50;
            var yMin = -50;
            var yMax = 50;

            // Act
            var world = new World(xMin, xMax, yMin, yMax);

            // Assert
            Assert.AreEqual(xMin, world.XMin);
            Assert.AreEqual(xMax, world.XMax);
            Assert.AreEqual(yMin, world.YMin);
            Assert.AreEqual(yMax, world.YMax);
        }
    }
}
