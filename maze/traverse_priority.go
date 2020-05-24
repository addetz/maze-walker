package maze

import "sort"

type priorityTraverser struct {
    // list with the cheapest point first
    toVisit []*priorityPoint
    //set implementation of point key markers
    visited map[string]void
}

type priorityPoint struct {
    point *Point
    cost int
}

func (pt *priorityTraverser) popFirstNode() (*Point, int) {
    first := pt.toVisit[0]
    pt.toVisit = pt.toVisit[1:]
    return first.point, first.cost
}

func (pt *priorityTraverser) enqueueNodes(points []*Point, cost int) {
    nodes := make([]*priorityPoint, len(points))
    for i, p := range points {
        nodes[i] = &priorityPoint{
            point: p,
            cost:  cost,
        }
    }
    allNodes := append(pt.toVisit, nodes...)
    sort.SliceStable(allNodes, func(i, j int) bool {
        return allNodes[i].cost < allNodes[j].cost
    })
    pt.toVisit = allNodes
}


func (pt *priorityTraverser) visitNode(p *Point) {
    pt.visited[p.getKey()] = setMarker
}

func (pt *priorityTraverser) isNodeVisited(p *Point) bool{
    _, ok := pt.visited[p.getKey()]
    return ok
}

func (pt *priorityTraverser) isVisitComplete() bool {
    return len(pt.toVisit) == 0
}