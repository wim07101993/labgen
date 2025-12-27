package main

import (
	"fmt"

	"github.com/wim07101993/labgen/svg"
)

type LabelBuilder struct {
	TitleFontSize    int
	SubtitleFontSize int
	LengthFontSize   int
	LabelHeight      float64
	LabelWidth       float64
	Connector        ConnectorCfg
	Cable            CableCfg
	PropsWidth       float64
	PropsHeight      float64
}

type ConnectorCfg struct {
	Width   float64
	Height  float64
	Padding float64
	Cable   CableCfg
}

type CableCfg struct {
	Thickness float64
	Length    float64
}

func (builder *LabelBuilder) BuildLabel(cable Cable) (*svg.Svg, error) {
	background, err := cable.BackgroundColor()
	if err != nil {
		return nil, err
	}
	foreground := cable.ForegroundColor()

	subTitle := cable.subTitle()

	titleTxt := &svg.Text{
		Id:       "title",
		FontSize: builder.TitleFontSize,
		Y:        float64(builder.TitleFontSize),
		Text:     cable.Title,
	}

	diagram := builder.BuildCableDiagram(cable.ConnectorsSideA, cable.ConnectorsSideB)

	const propsLeftPadding = 11

	props := []any{
		titleTxt,
		&svg.Use{
			Width:  builder.PropsWidth,
			Height: builder.PropsHeight - float64(builder.TitleFontSize+builder.SubtitleFontSize) - 2,
			Transform: &svg.Translate{
				Y: float64(builder.TitleFontSize + builder.SubtitleFontSize),
			},
			Href: diagram.Id.Href(),
		},
	}

	if subTitle != "" {
		txt := &svg.Text{
			Id:        "subtitle",
			FontSize:  builder.SubtitleFontSize,
			Transform: &svg.Translate{Y: titleTxt.Y + 24},
		}
		props = append(props, txt)
	}

	return &svg.Svg{
		NameSpace: svg.Namespace,
		ViewBox:   &svg.ViewBox{MinY: -builder.LabelHeight / 2, Width: builder.LabelWidth, Height: builder.LabelHeight},
		Width:     fmt.Sprintf("%fmm", builder.LabelWidth),
		Height:    fmt.Sprintf("%fmm", builder.LabelHeight),
		Children: []any{
			svg.Style{Value: "text { font-family: arial; }"},
			&svg.Rect{ // BACKGROUND
				Width:  builder.LabelWidth,
				Height: builder.LabelHeight,
				Y:      -builder.LabelHeight / 2,
				Fill:   background.Ptr(),
			},
			&svg.Rect{ // PROPERTIES BACKGROUND
				X:      builder.LabelWidth * 3 / 20,
				Y:      -builder.PropsHeight / 2,
				Width:  builder.PropsWidth,
				Height: builder.PropsHeight,
				Fill:   svg.Silver.Ptr(),
			},
			&svg.G{
				Transform: &svg.Translate{X: propsLeftPadding, Y: -builder.PropsHeight/2 + 1},
				Children:  props,
			},
			&svg.Text{ // LENGTH
				Id:         "length",
				Y:          float64(builder.LengthFontSize)/2 - (float64(builder.LengthFontSize) / 6),
				X:          58,
				Fill:       foreground.Ptr(),
				TextAnchor: svg.TextAnchorEnd,
				FontWeight: 70,
				FontSize:   builder.LengthFontSize,
				Text:       fmt.Sprintf("%gm", cable.Length),
			},
			&svg.G{ // LOGO
				Transform: &svg.Translate{X: 1},
				Children: []any{
					&svg.Use{
						Href:      fmt.Sprintf("#%s", Logo.Id),
						Width:     builder.LabelWidth,
						Height:    builder.LabelHeight / 10,
						X:         -builder.LabelWidth / 2,
						Y:         -builder.LabelHeight / 10,
						Transform: &svg.Rotate{Degrees: 90},
						Fill:      foreground.Ptr(),
					},
				},
			},
			&svg.Defs{Defs: []any{
				Logo,
				diagram,
			}},
		},
	}, nil
}

func (builder *LabelBuilder) BuildCableDiagram(a []ConnectorCount, b []ConnectorCount) *svg.Svg {
	diagramWidth := (builder.Connector.Width+builder.Connector.Cable.Length)*2 + builder.Cable.Length

	cable := &svg.Line{
		X1:          -builder.Cable.Length / 2,
		X2:          builder.Cable.Length / 2,
		Stroke:      svg.Black.Ptr(),
		StrokeWidth: 4,
	}

	s := &svg.Svg{
		Id:        "diagram",
		NameSpace: svg.Namespace,
		Children: []any{
			cable,
		},
	}

	wiresA := builder.buildWires(a, true)
	wiresB := builder.buildWires(b, false)

	gA, defsA, heightA := builder.buildConnectors(a)
	gB, defsB, heightB := builder.buildConnectors(b)

	gA.Transform = &svg.Translate{
		X: -diagramWidth/2 + builder.Connector.Width,
		Y: -heightA / 2,
	}
	gB.Transform = &svg.Translate{
		X: diagramWidth / 2,
		Y: -heightB / 2,
	}

	wiresA.Transform = &svg.Translate{
		X: -diagramWidth/2 + builder.Connector.Width,
		Y: -(heightA/2 - builder.Connector.Height/2),
	}
	wiresB.Transform = &svg.Translate{
		X: builder.Cable.Length / 2,
		Y: -(heightB/2 - builder.Connector.Height/2),
	}

	height := max(heightA, heightB)
	s.Height = fmt.Sprintf("%v", height)
	s.ViewBox = &svg.ViewBox{
		MinX:   -diagramWidth / 2,
		MinY:   -height / 2,
		Width:  diagramWidth,
		Height: height,
	}

	defs := make([]any, 0, len(defsA)+len(defsB))
	for _, def := range defsA {
		defs = append(defs, def)
	}
	for _, def := range defsB {
		defs = append(defs, def)
	}

	splitterA := builder.buildSplitter(cable.X1, heightA)
	splitterB := builder.buildSplitter(cable.X2, heightB)

	s.Children = append(s.Children,
		splitterA, splitterB,
		wiresA, wiresB,
		gA, gB,
		&svg.Defs{Defs: defs},
	)

	return s
}

func (builder *LabelBuilder) buildWires(cs []ConnectorCount, isLeft bool) *svg.G {
	g := &svg.G{
		Children: make([]any, 0, len(cs)*2),
	}

	var x1 float64
	var x2 float64
	if isLeft {
		x1 = -builder.Connector.Width / 2
		x2 = builder.Connector.Cable.Length
	} else {
		x1 = 0
		x2 = builder.Connector.Width/2 + builder.Connector.Cable.Length
	}
	for i := range cs {
		y := float64(i) * (builder.Connector.Height + builder.Connector.Padding)
		g.Children = append(g.Children, &svg.Line{
			X1:          x1,
			X2:          x2,
			Y1:          y,
			Y2:          y,
			Stroke:      svg.Black.Ptr(),
			StrokeWidth: builder.Cable.Thickness,
		})
	}

	return g
}

func (builder *LabelBuilder) buildSplitter(x float64, connectorsHeight float64) *svg.Line {
	return &svg.Line{
		X1:          x,
		X2:          x,
		Y1:          -(connectorsHeight/2 - builder.Connector.Height/2 + builder.Connector.Cable.Thickness/2),
		Y2:          connectorsHeight/2 - builder.Connector.Height/2 + builder.Connector.Cable.Thickness/2,
		Stroke:      svg.Black.Ptr(),
		StrokeWidth: 4,
	}
}

func (builder *LabelBuilder) buildConnectors(cs []ConnectorCount) (g *svg.G, defs []any, height float64) {
	defs = make([]any, len(cs))
	g = &svg.G{
		Children: make([]any, len(cs)),
	}

	for i, e := range cs {
		defs[i] = e.Connector.Svg()
		if i == 0 {
			height += builder.Connector.Height + builder.Connector.Padding
		} else {
			height += builder.Connector.Height
		}
		g.Children[i] = &svg.Use{
			Href:      e.Connector.Svg().Id.Href(),
			Width:     builder.Connector.Width,
			Height:    builder.Connector.Height,
			Y:         float64(i) * (builder.Connector.Height + builder.Connector.Padding),
			Transform: svg.MirrorX,
		}
	}

	return g, defs, height
}
