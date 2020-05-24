package maze

import (
    "errors"
)

func (m *Maze) BFS() ([]*Point, error){
    tr := &traverser{
        toVisit: make([]*Point, 0),
        visited:  make(map[string]void, 0),
    }
    cp := m.startPoint
    tr.enqueueNodes(m.getLegalNextMoves(cp))
    tr.visitNode(cp)
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
    if tr.isVisitComplete() {
        return []*Point{}, errors.New("destination unreachable")
    }
    // change the current point - BFS pops the first node like a queue
    cp = tr.popFirstNode()
    // next point has already been visited
    if tr.isNodeVisited(cp) {
        return m.bfs_traverse(cp, tr, path)
    }
    tr.enqueueNodes(m.getLegalNextMoves(cp))
    tr.visitNode(cp)
    newPath := append(path, cp)
    return m.bfs_traverse(cp, tr, newPath)
}

