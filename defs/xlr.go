package defs

import "github.com/wim07101993/labgen/svg"

var Xlr3PM = &svg.Svg{
	NameSpace: svg.Namespace,
	Id:        "xlr-3p-m",
	ViewBox: &svg.ViewBox{
		MinX:   -120,
		MinY:   -120,
		Width:  240,
		Height: 240,
	},
	Children: []any{
		&svg.Circle{
			Radius:      114,
			Stroke:      svg.Black.Ptr(),
			StrokeWidth: 12,
			Fill:        svg.White.Ptr(),
		},
		&svg.Circle{Radius: 24, CenterX: -60},
		&svg.Circle{Radius: 24, CenterX: 60},
		&svg.Circle{Radius: 24, CenterY: 60},
	},
}

var Xlr3PF = &svg.Svg{
	NameSpace: svg.Namespace,
	Id:        "xlr-3p-f",
	ViewBox: &svg.ViewBox{
		MinX:   -120,
		MinY:   -120,
		Width:  240,
		Height: 240,
	},
	Children: []any{
		&svg.Circle{
			Radius:      114,
			Stroke:      svg.Black.Ptr(),
			StrokeWidth: 12,
			Fill:        svg.Black.Ptr(),
		},
		&svg.Circle{Radius: 24, CenterX: -60, Fill: svg.White.Ptr()},
		&svg.Circle{Radius: 24, CenterX: 60, Fill: svg.White.Ptr()},
		&svg.Circle{Radius: 24, CenterY: 60, Fill: svg.White.Ptr()},
	},
}
