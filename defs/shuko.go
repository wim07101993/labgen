package defs

import "github.com/wim07101993/labgen/svg"

var ShukoF = &svg.Svg{
	NameSpace: svg.Namespace,
	Id:        "shuko-f",
	ViewBox: &svg.ViewBox{
		MinX:   -21,
		MinY:   -21,
		Width:  42,
		Height: 42,
	},
	Children: []any{
		&svg.Circle{
			Radius:      20,
			Stroke:      svg.Black.Ptr(),
			StrokeWidth: 2,
			Fill:        svg.White.Ptr(),
		},
		&svg.Rect{X: -2.5, Y: -20, Width: 5, Height: 5},
		&svg.Rect{X: -2.5, Y: 15, Width: 5, Height: 5},
		svg.Circle{Radius: 5.5, CenterX: -9.5},
		svg.Circle{Radius: 5.5, CenterX: 9.5},
	},
}

var ShukoM = &svg.Svg{
	NameSpace: svg.Namespace,
	Id:        "shuko-m",
	ViewBox: &svg.ViewBox{
		MinX:   -19,
		MinY:   -16,
		Width:  38,
		Height: 32,
	},
	Children: []any{
		&svg.Rect{
			X:           0,
			Y:           -16,
			Width:       19,
			Height:      32,
			Fill:        svg.Black.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg.Black.Ptr(),
		},
		&svg.Rect{
			X:           -19,
			Y:           -12.3,
			Width:       19,
			Height:      4.8,
			Fill:        svg.White.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg.Black.Ptr(),
		},
		&svg.Rect{
			X:           -19,
			Y:           7.5,
			Width:       19,
			Height:      4.8,
			Fill:        svg.White.Ptr(),
			StrokeWidth: 1,
			Stroke:      svg.Black.Ptr(),
		},
		&svg.Path{
			D:      "m19,-16 l4,9 v14 l-4,9 z",
			Fill:   svg.Black.Ptr(),
			Stroke: svg.Black.Ptr(),
		},
	},
}
