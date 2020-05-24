package maze

import "fmt"

type Point struct {
    X,Y int
    IsWall bool
    IsStart bool
    IsDestination bool
}

// getNeighbors returns all the 4 points that we can walk to from this point
// [x,y] -> [x-1,y], [x+1,y], [x, y-1], [x, y+1]
func (p Point) getNeighbors() []string{
    n := make([]string, 4)
    n[0] = generateKey(p.X - 1, p.Y)
    n[1] = generateKey(p.X + 1, p.Y)
    n[2] = generateKey(p.X, p.Y - 1)
    n[3] = generateKey(p.X, p.Y + 1)
    return n
}

func (p Point) getKey() string {
    return generateKey(p.X, p.Y)
}

func generateKey(x,y int) string {
    return fmt.Sprintf("[%d-%d]", x, y)
}