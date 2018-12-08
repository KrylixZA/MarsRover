# Mars Rover

This a solution to a Mars Rover Kata designed to teach the fundamentals of Test-Driven Development.

## The Kata
Develop an API that moves a rover around a grid on Mars.
* You are given the initial starting point(x,y) of a rover and the direction (N,S,E,W) it is facing.
* The rover receives a character array of commands
    * Implement commands that:
    * Move the rover forward(f)
    * Move the rover backward(b)
    * Turn the rover left(l)
    * Turn the rover right(r)

* Implement wrapping from one edge or the grid to another ( Planets are spheres after all )

### Bonus
Implement obstacle detection before each move to a new square. If a given sequence of commands encounters an
obstacle the rover moves up to the last possible point and reports the obstacle. You will need to amend your code to
take in an array of obstacles.