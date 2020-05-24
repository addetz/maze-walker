package maze

import (
    "errors"
)

func (m *Maze) Djikstra() ([]*Point, error){
    pt := &priorityTraverser{
        toVisit: make([]*priorityPoint, 0),
        visited:  make(map[string]void, 0),
    }
    cp := m.startPoint
    pt.enqueueNodes(m.getLegalNextMoves(cp), 1)
    pt.visitNode(cp)
    return m.djikstra_traverse(cp, pt, []*Point{cp})
}

// bfs_traverse looks for the end point and returns a path
// or an error if we cannot get there
func (m *Maze) djikstra_traverse(cp *Point, pt *priorityTraverser, path []*Point) ([]*Point, error) {
    // we made it to the destination! return the path and get out!
    if cp.IsDestination {
        return path, nil
    }
    // nothing more to visit and there was no destination
    if pt.isVisitComplete() {
        return []*Point{}, errors.New("destination unreachable")
    }
    // change the current point - BFS pops the first node like a queue
    current, cost := pt.popFirstNode()
    // next point has already been visited
    if pt.isNodeVisited(current) {
        return m.djikstra_traverse(current, pt, path)
    }
    pt.enqueueNodes(m.getLegalNextMoves(current), cost+1)
    pt.visitNode(current)
    newPath := append(path, current)
    return m.djikstra_traverse(current, pt, newPath)
}


