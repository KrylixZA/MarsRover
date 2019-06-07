namespace MarsRover
{
    public class Rover
    {
        public Position Position { get; set; }
        public Direction Direction { get; set; }
        public World World { get; set; }

        public Rover() : this(0, 0)
        {
        }

        public Rover(int xInitPos, int yInitPos) : this(new Position(xInitPos, yInitPos))
        {
        }

        public Rover(Position initPosition, Direction direction = Direction.East) : this(initPosition, direction, new World(0, 1, 0 , 1))
        {
        }

        public Rover(Position initPosition, Direction direction, World world)
        {
            this.Position = initPosition;
            this.Direction = direction;
            this.World = world;
        }

        public void Move(int nrToMove)
        {
            switch (this.Direction)
            {
                case Direction.East:
                    this.Position.X += nrToMove;
                    if (this.Position.X > this.World.XMax) {
                        this.Position.X = (this.Position.X % this.World.XMax) + this.World.XMin;
                    }
                    break;
                case Direction.West:
                    this.Position.X -= nrToMove;
                    if (this.Position.X < this.World.XMin)
                    {
                        this.Position.X = (this.Position.X % this.World.XMax) + this.World.XMax;
                    }
                    break;
                case Direction.North:
                    this.Position.Y += nrToMove;
                    if (this.Position.Y > this.World.YMax)
                    {
                        this.Position.Y = (this.Position.Y % this.World.YMax) + this.World.YMin;
                    }
                    break;
                case Direction.South:
                    this.Position.Y -= nrToMove;
                    if (this.Position.Y < this.World.YMin)
                    {
                        this.Position.Y = (this.Position.Y % this.World.YMax) + this.World.YMax;
                    }
                    break;
            }
        }

        public void TurnLeft()
        {
            switch (this.Direction)
            {
                case Direction.East:
                    this.Direction = Direction.North;
                    break;
                case Direction.North:
                    this.Direction = Direction.West;
                    break;
                case Direction.West:
                    this.Direction = Direction.South;
                    break;
                case Direction.South:
                    this.Direction = Direction.East;
                    break;
            }
        }

        public void TurnRight()
        {
            switch (this.Direction)
            {
                case Direction.East:
                    this.Direction = Direction.South;
                    break;
                case Direction.South:
                    this.Direction = Direction.West;
                    break;
                case Direction.West:
                    this.Direction = Direction.North;
                    break;
                case Direction.North:
                    this.Direction = Direction.East;
                    break;
            }
        }
    }
}
