package maze

import (
    "bufio"
    "github.com/pkg/errors"
    "os"
    "strings"
)

type Maze struct {
    input string
    startPoint *Point
    conf *Config
    points map[string]*Point
}

type Config struct {
    StartMarker, EndMarker, WallMarker, PointDelim string
}

func New(input string, config *Config) (*Maze, error) {
    m := &Maze{
        input: input,
        conf: config,
        points: make(map[string]*Point, 0),
    }
    if err := m.parse(input); err != nil {
        return nil, errors.Wrap(err, "unable to create maze")
    }
    return m, nil
}

func (m *Maze) parse(path string) error{
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var xcount int
    for scanner.Scan() {
        m.parseLine(xcount, scanner.Text())
        xcount++
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

// parseLine takes in an x coord and a line and parses it
func (m *Maze) parseLine(xcount int, text string) {
    pp := strings.Split(text, m.conf.PointDelim)
    for i, p := range pp {
        pnt := &Point{
            X:  xcount,
            Y: i,
            IsDestination: m.isDestinationMarker(p),
            IsStart: m.isStartMarker(p),
            IsWall:  m.isWallMarker(p),
        }
        m.points[pnt.getKey()] = pnt
        if pnt.IsStart {
            m.startPoint = pnt
        }
    }
}

func (m *Maze) isDestinationMarker(p string) bool {
    return p == m.conf.EndMarker
}

func (m *Maze) isStartMarker(p string) bool {
    return p == m.conf.StartMarker
}

func (m *Maze) isWallMarker(p string) bool {
    return p == m.conf.WallMarker
}

func (m *Maze) isLegalMove(key string) bool {
    pnt, ok := m.points[key]
    if !ok {
        return false
    }
    return !pnt.IsWall
}

func (m *Maze) getLegalNextMoves(cp *Point) []*Point {
    legalNeighbors := make([]*Point,0)
    for _, nkey := range cp.getNeighbors() {
        if m.isLegalMove(nkey) {
            legalNeighbors = append(legalNeighbors, m.points[nkey])
        }
    }
    return legalNeighbors
}