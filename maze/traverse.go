package maze

import (
    "fmt"
    "strings"
)
type void struct {}
var setMarker void

type traverser struct {
    toVisit []*Point
    //set implementation of point key markers
    visited map[string]void
}

func (tr *traverser) enqueueNodes(nodes []*Point) {
    tr.toVisit = append(tr.toVisit, nodes...)
}

func (tr *traverser) popLastNode() *Point{
    vlen := len(tr.toVisit)
    last := tr.toVisit[vlen -1]
    tr.toVisit = tr.toVisit[0:vlen-1]
    return last
}
func (tr *traverser) popFirstNode() *Point{
    first := tr.toVisit[0]
    tr.toVisit = tr.toVisit[1:]
    return first
}

func (tr *traverser) visitNode(p *Point) {
    tr.visited[p.getKey()] = setMarker
}

func (tr *traverser) isNodeVisited(p *Point) bool{
    _, ok := tr.visited[p.getKey()]
    return ok
}

func (tr *traverser) isVisitComplete() bool {
    return len(tr.toVisit) == 0
}

func PrintPath(plist []*Point) string{
    b := &strings.Builder{}
    for _, p := range plist {
        fmt.Fprintf(b, " [%d,%d] ", p.X, p.Y)
    }
    return b.String()
}
