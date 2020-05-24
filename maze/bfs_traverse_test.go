package maze

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMaze_BFS(t *testing.T) {
    var (
        m *Maze
        startPoint *Point
    )

    /*
    A _ _
    _ X _
    _ _ B
     */
    setupRegular := func() {
        points := make(map[string]*Point, 0)
        startPoint = makePoint( points, 0, 0, false, true, false)
        makePoint(points, 0,1, false, false,false)
        makePoint(points, 0,2, false, false,false)
        makePoint(points, 1,0, false, false, false)
        makePoint(points, 1,1, true, false,false)
        makePoint(points, 1,2, false, false,false)
        makePoint(points, 2,0, false, false, false)
        makePoint(points, 2,1, true, false,false)
        makePoint(points, 2,2, false, false,true)
        m = &Maze {
            startPoint: startPoint,
            points: points,
        }
    }

    /*
    A X _
    _ X _
    _ X B
     */
    setupImpossible := func() {
        points := make(map[string]*Point, 0)
        startPoint = makePoint( points, 0, 0, false, true, false)
        makePoint(points, 0,1, true, false,false)
        makePoint(points, 0,2, false, false,false)
        makePoint(points, 1,0, false, false, false)
        makePoint(points, 1,1, true, false,false)
        makePoint(points, 1,2, false, false,false)
        makePoint(points, 2,0, false, false, false)
        makePoint(points, 2,1, true, false,false)
        makePoint(points, 2,2, false, false,true)
        m = &Maze {
            startPoint: startPoint,
            points: points,
        }
    }

    t.Run("bfs traverse", func(t *testing.T) {
        setupRegular()
        path, err := m.BFS()
        assert.Nil(t, err)
        assert.Equal(t, len(path),7)
    })
    t.Run("bfs traverse impossible", func(t *testing.T) {
        setupImpossible()
        path, err := m.BFS()
        assert.NotNil(t, err)
        assert.Equal(t, len(path),0)
    })
}

func makePoint(points map[string]*Point, x,y int, isWall, isStart, isDest bool) *Point{
   p := &Point{
        X:             x,
        Y:             y,
        IsWall:        isWall,
        IsStart:       isStart,
        IsDestination: isDest,
    }
    points[p.getKey()] = p
    return p
}
