package maze

import (
    "fmt"
    "strings"
)

type traverser struct {
    toVisit []*Point
    //set implementation of point key markers
    visited map[string]void
}

func PrintPath(plist []*Point) string{
    b := &strings.Builder{}
    for _, p := range plist {
        fmt.Fprintf(b, " [%d,%d] ", p.X, p.Y)
    }
    return b.String()
}
