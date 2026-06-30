package svg_components

import (
	svg2 "github.com/wim07101993/labgen/packages/svg"
)

var ShukoF = &svg2.Svg{
	NameSpace: svg2.Namespace,
	Id:        "shuko-f",
	ViewBox: &svg2.ViewBox{
		MinX:   -21,
		MinY:   -21,
		Width:  42,
		Height: 42,
	},
	Children: []any{
		&svg2.Circle{
			Radius:      20,
			Stroke:      svg2.Black.Ptr(),
			StrokeWidth: 2,
			Fill:        svg2.White.Ptr(),
		},
		&svg2.Rect{X: -2.5, Y: -20, Width: 5, Height: 5},
		&svg2.Rect{X: -2.5, Y: 15, Width: 5, Height: 5},
		svg2.Circle{Radius: 5.5, CenterX: -9.5},
		svg2.Circle{Radius: 5.5, CenterX: 9.5},
	},
}

var ShukoM = &svg2.Svg{
	NameSpace: svg2.Namespace,
	Id:        "shuko-m",
	ViewBox: &svg2.ViewBox{
		MinX:   -19,
		MinY:   -16,
		Width:  38,
		Height: 32,
	},
	Children: []any{
		&svg2.Rect{
			X:           0,
			Y:           -16,
			Width:       19,
			Height:      32,
			Fill:        svg2.Black.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg2.Black.Ptr(),
		},
		&svg2.Rect{
			X:           -19,
			Y:           -12.3,
			Width:       19,
			Height:      4.8,
			Fill:        svg2.White.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg2.Black.Ptr(),
		},
		&svg2.Rect{
			X:           -19,
			Y:           7.5,
			Width:       19,
			Height:      4.8,
			Fill:        svg2.White.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg2.Black.Ptr(),
		},
		&svg2.Path{
			D:      "m19,-16 l4,9 v14 l-4,9 z",
			Fill:   svg2.Black.Ptr(),
			Stroke: svg2.Black.Ptr(),
		},
	},
}
