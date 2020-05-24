package main

import (
    "fmt"
    "github.com/maze-walker/maze"
)

func main() {
    inputPath := "input.txt"
    m, err := maze.New(inputPath, &maze.Config{
        StartMarker: START_POINT_MARKER,
        EndMarker:   END_POINT_MARKER,
        WallMarker:  WALL_MARKER,
        PointDelim: POINT_DELIM,
    })
    if err != nil {
        fmt.Printf("Maze from %s was not created successfully: %v \n", inputPath, err)
        return
    }
    fmt.Printf("Maze from %s was created successfully \n", inputPath)
    bfsPath, err := m.BFS()
    if err != nil {
        fmt.Println("Your skeleton will lie in the Maze forever.")
        return
    }
    fmt.Printf("BFS(%d) path: %s \n",
        len(bfsPath), maze.PrintPath(bfsPath))

    dfsPath, err := m.DFS()
    if err != nil {
        fmt.Println("Your skeleton will lie in the Maze forever.")
        return
    }
    fmt.Printf("DFS(%d) path: %s \n",
        len(dfsPath), maze.PrintPath(dfsPath))

    djikstraPath, err := m.Djikstra()
    if err != nil {
        fmt.Println("Your skeleton will lie in the Maze forever.")
        return
    }
    fmt.Printf("Djikstra(%d) path: %s \n", len(djikstraPath), maze.PrintPath(djikstraPath))
}
