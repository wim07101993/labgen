package defs

import "github.com/wim07101993/labgen/svg"

var IecC13 = buildIecC13()

func buildIecC13() *svg.Svg {
	hole := &svg.Path{
		Id:          "hole",
		D:           "M0,-25 h12.5 v50 h-25 v-50 h12.5 z",
		Stroke:      svg.Black.Ptr(),
		StrokeWidth: 10,
		Fill:        svg.Black.Ptr(),
	}

	return &svg.Svg{
		NameSpace: svg.Namespace,
		Id:        "iec-c13",
		ViewBox: &svg.ViewBox{
			MinX:   -70,
			MinY:   -130,
			Width:  140,
			Height: 230,
		},
		Children: []any{
			&svg.Path{
				Fill:        svg.White.Ptr(),
				Stroke:      svg.Black.Ptr(),
				StrokeWidth: 10,
				D:           "M0,-79 h60 l55,55 v103 h-230 v-103 l55,-55 z",
			},
			&svg.Use{Href: hole.Id.Href(), Y: -12.5},
			&svg.Use{Href: hole.Id.Href(), Y: 20, X: -70},
			&svg.Use{Href: hole.Id.Href(), Y: 20, X: 70},
			&svg.Defs{
				Defs: []any{hole},
			},
		},
	}
}
