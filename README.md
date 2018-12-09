# Mars Rover
A simple example of a Mars Rover that has to do a number of things...with a twist. This entire project was solved using Test-Driven Development following the Red/Green/Refactor principles.

## The Kata
Develop an API that moves a rover around a grid on Mars.
* You are given the initial starting point(x,y) of a rover and the direction (N,S,E,W) it is facing.
* The rover receives an array of commands (implementation is at your discretion)
    * Implement commands that:
    * Move the rover forward(f)
    * Move the rover backward(b)
    * Turn the rover left(l)
    * Turn the rover right(r)

* Implement wrapping from one edge or the grid to another (planets are spheres after all)

### Bonus
Implement obstacle detection before each move to a new square. If a given sequence of commands encounters an
obstacle the rover moves up to the last possible point and reports the obstacle. You will need to amend your code to
take in an array of obstacles.

##### Credit
All credit goes to [Chillisoft](http://www.chillisoft.co.za/) who provided me with the Katas and the training to solve these problems.