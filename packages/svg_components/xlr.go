package svg_components

import (
	svg2 "github.com/wim07101993/labgen/packages/svg"
)

var Xlr3PM = &svg2.Svg{
	NameSpace: svg2.Namespace,
	Id:        "xlr-3p-m",
	ViewBox: &svg2.ViewBox{
		MinX:   -120,
		MinY:   -120,
		Width:  240,
		Height: 240,
	},
	Children: []any{
		&svg2.Circle{
			Radius:      114,
			Stroke:      svg2.Black.Ptr(),
			StrokeWidth: 12,
			Fill:        svg2.White.Ptr(),
		},
		&svg2.Circle{Radius: 24, CenterX: -60},
		&svg2.Circle{Radius: 24, CenterX: 60},
		&svg2.Circle{Radius: 24, CenterY: 60},
	},
}

var Xlr3PF = &svg2.Svg{
	NameSpace: svg2.Namespace,
	Id:        "xlr-3p-f",
	ViewBox: &svg2.ViewBox{
		MinX:   -120,
		MinY:   -120,
		Width:  240,
		Height: 240,
	},
	Children: []any{
		&svg2.Circle{
			Radius:      114,
			Stroke:      svg2.Black.Ptr(),
			StrokeWidth: 12,
			Fill:        svg2.Black.Ptr(),
		},
		&svg2.Circle{Radius: 24, CenterX: -60, Fill: svg2.White.Ptr()},
		&svg2.Circle{Radius: 24, CenterX: 60, Fill: svg2.White.Ptr()},
		&svg2.Circle{Radius: 24, CenterY: 60, Fill: svg2.White.Ptr()},
	},
}
