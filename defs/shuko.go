package defs

import "github.com/wim07101993/labgen/svg"

var ShukoF = &svg.Svg{
	NameSpace: svg.Namespace,
	Id:        "shuko-f",
	ViewBox: &svg.ViewBox{
		MinX:   -60,
		MinY:   -60,
		Width:  120,
		Height: 120,
	},
	Children: []any{
		&svg.Circle{
			Radius:      114,
			Stroke:      svg.Black.Ptr(),
			StrokeWidth: 12,
			Fill:        svg.White.Ptr(),
		},
		&svg.Rect{X: -15, Y: -115, Width: 30, Height: 30},
		&svg.Rect{X: -15, Y: 85, Width: 30, Height: 30},
		svg.Circle{Radius: 30, CenterX: -55},
		svg.Circle{Radius: 30, CenterX: 55},
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
