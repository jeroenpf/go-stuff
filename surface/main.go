package main

import (
	"fmt"
	"learn/hex"
	"math"
)

const (
	width, height = 600, 320
	cells         = 150
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // Sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill: white; stroke-width: 0.5' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z, aok := corner(i+1, j)
			bx, by, _, bok := corner(i, j)
			cx, cy, _, cok := corner(i, j+1)
			dx, dy, _, dok := corner(i+1, j+1)

			if aok && bok && cok && dok {
				step := (((z + 1) * 255) / 2)
				color := hex.GetGradient(0x139100, 0x001196, step)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:#%x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, 0, false
	}

	// Project (x, y, z)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Distance from (0, 0)
	return math.Sin(r) / r
}
