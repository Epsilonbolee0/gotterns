package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

type Line struct {
	X1, Y1, X2, Y2 int
}

type VectorImage struct {
	Lines []Line
}

// The interface given
func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1

	return &VectorImage{[]Line{
		Line{0, 0, width, 0},
		Line{0, 0, 0, height},
		Line{width, 0, width, height},
		Line{0, height, width, height},
	}}
}

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

type runeMatrix struct {
	data [][]rune
}

func NewRuneMatrix(width, height int) *runeMatrix {
	data := make([][]rune, height)
	for i := 0; i < height; i++ {
		data[i] = make([]rune, width)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	return &runeMatrix{data}
}

func (r *runeMatrix) Draw(points []Point) *runeMatrix {
	for _, point := range points {
		r.data[point.Y][point.X] = '*'
	}

	return r
}

func (r *runeMatrix) String() string {
	b := strings.Builder{}
	for _, line := range r.data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

type borders struct {
	Point
}

func NewBorders(points []Point) *borders {
	maxX, maxY := 0, 0

	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}
		if point.Y > maxY {
			maxY = point.Y
		}
	}

	return &borders{Point{maxX + 1, maxY + 1}}
}

func DrawPoints(owner RasterImage) string {
	points := owner.GetPoints()

	b := NewBorders(points)

	return NewRuneMatrix(b.X, b.Y).
		Draw(points).
		String()

}

var pointCache = map[[16]byte][]Point{}

type vectorToRasterAdapter struct {
	points []Point
}

func (a vectorToRasterAdapter) GetPoints() []Point {
	return a.points
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("We currently have", len(a.points), "points")
}

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if points, ok := pointCache[h]; ok {
		for _, point := range points {
			a.points = append(a.points, point)
		}
	} else {
		a.addLine(line)
		pointCache[h] = a.points
	}
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}

	return adapter
}

func main() {
	rc := NewRectangle(60, 14)
	rc2 := NewRectangle(10, 16)
	a := VectorToRaster(rc)
	b := VectorToRaster(rc2)

	fmt.Print(DrawPoints(a))
	fmt.Print(DrawPoints(b))
}
