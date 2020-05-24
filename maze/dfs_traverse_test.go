package maze

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestMaze_DFS(t *testing.T) {
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

    t.Run("dfs traverse", func(t *testing.T) {
        setupRegular()
        path, err := m.DFS()
        assert.Nil(t, err)
        assert.Equal(t, 5, len(path))
    })
    t.Run("dfs traverse impossible", func(t *testing.T) {
        setupImpossible()
        path, err := m.DFS()
        assert.NotNil(t, err)
        assert.Equal(t, 0, len(path))
    })
}