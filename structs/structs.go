package structs

import "fmt"

type Point struct{X, Y int}
type Wheel struct {
    p Point
    r int
}

func Scale(p *Point, factor int) *Point {
    p.X = p.X * factor
    p.Y = p.Y * factor
    //return Point{p.X * factor, p.Y * factor}
    return p
}

func MakePoint(point *Point) *Point {
    point.X = 10
    point.Y = 20
    fmt.Println("Make point function")
    return point
}

func MakeWheel(w *Wheel) *Wheel {
    w.p.X = 10
    w.p.Y = 20
    w.r   = 30
    return w
}