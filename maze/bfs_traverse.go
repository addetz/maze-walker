package maze

import (
    "errors"
)

type void struct {}
var setMarker void

func (m *Maze) BFS() ([]*Point, error){
    tr := &traverser{
        toVisit: make([]*Point, 0),
        visited:  make(map[string]void, 0),
    }
    cp := m.startPoint
    tr.toVisit = append(tr.toVisit, m.getLegalNextMoves(cp)...)
    tr.visited[cp.getKey()] = setMarker
    return m.bfs_traverse(cp, tr, []*Point{cp})
}

// bfs_traverse looks for the end point and returns a path
// or an error if we cannot get there
func (m *Maze) bfs_traverse(cp *Point, tr *traverser, path []*Point) ([]*Point, error) {
    // we made it to the destination! return the path and get out!
    if cp.IsDestination {
        return path, nil
    }
    // nothing more to visit and there was no destination
    if len(tr.toVisit) == 0 {
        return []*Point{}, errors.New("destination unreachable")
    }

    // change the current point
    cp = tr.toVisit[0]
    tr.toVisit = tr.toVisit[1:]
    // next point has already been visited
    _, ok := tr.visited[cp.getKey()]
    if ok {
        return m.bfs_traverse(cp, tr, path)
    }
    tr.toVisit = append(tr.toVisit, m.getLegalNextMoves(cp)...)
    tr.visited[cp.getKey()] = setMarker
    newPath := append(path, cp)
    return m.bfs_traverse(cp, tr, newPath)
}

