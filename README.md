# maze-walker
You gimme a maze and I'll show you how to get out of it.

If you cannot get to the destination, I will tell you `Your skeleton will lie in the Maze forever.` 

The maze is actually an unweighted graph with loops.
The maze walker finds a solution to the maze using: 
- `BFS`- put all nodes in a queue and pop first node 
- `DFS`- put all nodes in a stack and pop last node
- `Djikstra` - put all nodes in a `PriorityQueue` and pop the cheapest node; 
will yield the same as `BFS` since each maze step is same cost

## Sample input
The project has a sample `input.txt`.
1. Start point is defined with the letter `A`. 
1. End point is defined with the letter `B`. 
1. Maze walls are defined with the letter `X`.
1. Maze points are split with a space.
1. Maze rows are split with `\n`.

```text
A _ _ _ _ _ _ X _ _
_ _ _ _ _ X _ X _ _
_ _ X X X X _ X _ _
_ _ X _ _ X _ X _ _
_ _ X _ _ X _ _ _ _
_ _ X _ _ X _ _ X _
_ _ _ _ _ X _ _ X B
```

## Constraints 
1. The maze walker cannot move on the diagonal. 
1. Some constants are expected in the maze input. See the sample input above. 
New constants can be configured in `config.go` 
1. The solution path will contain coordinates of how to get out of the maze 
indexed at `[0,0]` in the upper left corner.
