package main

import (
	"image/color"

	"github.com/mmcloughlin/globe"
)

func main() {
	g := globe.New()
	g.DrawGraticule(10.0)
	g.DrawLandBoundaries(globe.Color(color.NRGBA{255, 0, 0, 255})) // what is the "style" option here?
	g.CenterOn(51.453349, -2.588323)
	g.SavePNG("land.png", 400)
}
