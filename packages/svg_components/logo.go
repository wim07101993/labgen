package svg_components

import (
	svg2 "github.com/wim07101993/labgen/packages/svg"
)

var Logo = buildLogo()

func buildLogo() *svg2.Svg {
	parallelogram := &svg2.Path{Id: "parallelogram", D: "M 0,0 h 100 l 50,300 h -100 Z"}
	triangle := &svg2.Path{Id: "triangle", D: "M 0,0 h 100 l -50,300 Z"}

	return &svg2.Svg{
		Id:        "logo",
		NameSpace: svg2.Namespace,
		ViewBox:   &svg2.ViewBox{Width: 850, Height: 300},
		Children: []any{
			&svg2.G{
				Id: "w",
				Children: []any{
					&svg2.Use{Href: parallelogram.Id.Href()},
					&svg2.Use{
						Href:      triangle.Id.Href(),
						Transform: &svg2.Translate{X: 125},
					},
					&svg2.Use{
						Href: parallelogram.Id.Href(),
						Transform: svg2.Combiner{
							&svg2.Translate{X: 350},
							svg2.MirrorX,
						},
					},
				},
			},
			&svg2.G{
				Id:        "i",
				Transform: &svg2.Translate{X: 500},
				Children: []any{
					&svg2.Use{
						Href:      parallelogram.Id.Href(),
						Transform: svg2.MirrorX,
					},
				},
			},
			&svg2.G{
				Id:        "m",
				Transform: &svg2.Translate{X: 500},
				Children: []any{
					&svg2.Use{
						Href: "#w",
						Transform: svg2.Combiner{
							svg2.MirrorY,
							&svg2.Translate{Y: -300},
						},
					},
				},
			},
			&svg2.Defs{Defs: []any{parallelogram, triangle}},
		},
	}
}
