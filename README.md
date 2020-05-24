# maze-walker
You gimme a maze and I'll show you how to get out of it.

If you cannot get to the destination, I will tell you `Your skeleton will lie in the Maze forever.` 

## Constraints 
1. The maze walker cannot move on the diagonal. 
1. Some constants are expected in the maze input: 
    1. Start point is  defined with the letter `A`. 
    1. End point is defined with the letter `B`. 
    1. Maze walls are defined with the letter `X`.
    1. Maze points are split with a space.
    1. New constants can be configured in `config.go`
1. The maze walker finds a solution to the maze using `BFS` and `DFS`. 
1. The solution path will contain coordinates of how to get out of the maze 
indexed at `[0,0]` in the upper left corner.
1. We only support rectangular mazes. Please don't confuse us. 
