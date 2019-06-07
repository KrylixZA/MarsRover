using System;

namespace MarsRover
{
    public class World
    {
        public int XMin { get; set; }
        public int XMax { get; set; }
        public int YMin { get; set; }
        public int YMax { get; set; }

        public World() : this(0, 1, 0, 1)
        {
        }

        public World(int xMin, int xMax, int yMin, int yMax)
        {
            this.XMin = xMin;
            this.XMax = xMax;

            this.YMin = yMin;
            this.YMax = yMax;
        }
    }
}
