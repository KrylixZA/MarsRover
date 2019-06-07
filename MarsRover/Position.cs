namespace MarsRover
{
    public class Position
    {
        public int X { get; set; }
        public int Y { get; set; }

        public Position() : this(0, 0)
        {
        }

        public Position(int x, int y)
        {
            this.X = x;
            this.Y = y;
        }
    }
}
