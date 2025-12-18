package main

import (
	"github.com/wim07101993/labgen/svg"
)

var Logo = buildLogo()

func buildLogo() *svg.Svg {
	parallelogram := &svg.Path{Id: "parallelogram", D: "M 0,0 h 100 l 50,300 h -100 Z"}
	triangle := &svg.Path{Id: "triangle", D: "M 0,0 h 100 l -50,300 Z"}

	return &svg.Svg{
		Id:        "logo",
		NameSpace: svg.Namespace,
		ViewBox:   &svg.ViewBox{Width: 850, Height: 300},
		Children: []any{
			&svg.G{
				Id: "w",
				Children: []any{
					&svg.Use{Href: parallelogram.Id.Href()},
					&svg.Use{
						Href:      triangle.Id.Href(),
						Transform: &svg.Translate{X: 125},
					},
					&svg.Use{
						Href: parallelogram.Id.Href(),
						Transform: svg.Combiner{
							&svg.Translate{X: 350},
							svg.MirrorX,
						},
					},
				},
			},
			&svg.G{
				Id:        "i",
				Transform: &svg.Translate{X: 500},
				Children: []any{
					&svg.Use{
						Href:      parallelogram.Id.Href(),
						Transform: svg.MirrorX,
					},
				},
			},
			&svg.G{
				Id:        "m",
				Transform: &svg.Translate{X: 500},
				Children: []any{
					&svg.Use{
						Href: "#w",
						Transform: svg.Combiner{
							svg.MirrorY,
							&svg.Translate{Y: -300},
						},
					},
				},
			},
			&svg.Defs{Defs: []any{parallelogram, triangle}},
		},
	}
}
