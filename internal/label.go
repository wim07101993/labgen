package internal

import (
	"fmt"

	"github.com/wim07101993/labgen/internal/parts"
	svg2 "github.com/wim07101993/labgen/packages/svg"
	"github.com/wim07101993/labgen/packages/svg_components"
)

type LabelBuilder struct {
	LabelHeight      float64       `json:"labelHeight,omitempty"`
	LabelWidth       float64       `json:"labelWidth,omitempty"`
	LogoWidth        float64       `json:"logoWidth,omitempty"`
	TitleFontSize    int           `json:"titleFontSize,omitempty"`
	SubtitleFontSize int           `json:"subtitleFontSize,omitempty"`
	LengthFontSize   int           `json:"lengthFontSize,omitempty"`
	PropsWidth       float64       `json:"propsWidth,omitempty"`
	DiagramWidth     float64       `json:"diagramWidth,omitempty"`
	Connector        *ConnectorCfg `json:"connector,omitempty"`
	Cable            *CableCfg     `json:"cable,omitempty"`
	Padding          float64       `json:"padding,omitempty"`
}

type ConnectorCfg struct {
	Width   float64   `json:"width,omitempty"`
	Height  float64   `json:"height,omitempty"`
	Padding float64   `json:"padding,omitempty"`
	Cable   *CableCfg `json:"cable,omitempty"`
}

type CableCfg struct {
	Thickness float64 `json:"thickness,omitempty"`
	Length    float64 `json:"length,omitempty"`
}

func (builder *LabelBuilder) BuildLabel(cable parts.Cable) (*svg2.Svg, error) {
	background, err := cable.BackgroundColor()
	if err != nil {
		return nil, err
	}
	foreground := cable.ForegroundColor()

	logoHeight := builder.LogoWidth * svg_components.Logo.ViewBox.Width / svg_components.Logo.ViewBox.Height

	defs := &svg2.Defs{Defs: []any{svg_components.Logo}}

	props, diagramDefs := builder.BuildProps(cable)
	defs.Defs = append(defs.Defs, diagramDefs...)

	return &svg2.Svg{
		NameSpace: svg2.Namespace,
		ViewBox:   &svg2.ViewBox{MinY: -builder.LabelHeight / 2, Width: builder.LabelWidth, Height: builder.LabelHeight},
		Width:     fmt.Sprintf("%fmm", builder.LabelWidth),
		Height:    fmt.Sprintf("%fmm", builder.LabelHeight),
		Children: []any{
			svg2.Style{Value: "text { font-family: arial; }"},
			&svg2.Rect{ // BACKGROUND
				Width:  builder.LabelWidth,
				Height: builder.LabelHeight,
				Y:      -builder.LabelHeight / 2,
				Fill:   background.Ptr(),
			},
			props,
			&svg2.Text{ // LENGTH
				Id:         "length",
				Y:          float64(builder.LengthFontSize)/2 - (float64(builder.LengthFontSize) / 6),
				X:          builder.LabelWidth - builder.Padding,
				Fill:       foreground.Ptr(),
				TextAnchor: svg2.TextAnchorEnd,
				FontWeight: 700,
				FontSize:   builder.LengthFontSize,
				Text:       fmt.Sprintf("%gm", cable.Length),
			},
			&svg2.G{ // LOGO
				Transform: &svg2.Translate{X: builder.LogoWidth/2 + builder.Padding},
				Children: []any{
					&svg2.Use{
						Href:      fmt.Sprintf("#%s", svg_components.Logo.Id),
						Width:     logoHeight,
						Height:    builder.LogoWidth,
						X:         -logoHeight / 2,
						Y:         -builder.LogoWidth / 2,
						Transform: &svg2.Rotate{Degrees: 90},
						Fill:      foreground.Ptr(),
					},
				},
			},
			defs,
		},
	}, nil
}

func (builder *LabelBuilder) BuildProps(cable parts.Cable) (g *svg2.G, defs []any) {
	subTitle := cable.SubTitle()

	y := builder.Padding/2 + float64(builder.TitleFontSize)
	titleTxt := &svg2.Text{
		Id:       "title",
		FontSize: builder.TitleFontSize,
		Y:        y,
		Text:     cable.Title,
	}
	props := []any{titleTxt}

	diagram := builder.BuildCableDiagram(cable.ConnectorsSideA, cable.ConnectorsSideB)
	diagramHeight := diagram.ViewBox.Height / diagram.ViewBox.Width * builder.DiagramWidth

	if subTitle != "" {
		y = y + float64(builder.SubtitleFontSize)
		txt := &svg2.Text{
			Id:        "subtitle",
			Text:      subTitle,
			FontSize:  builder.SubtitleFontSize,
			Transform: &svg2.Translate{Y: y},
		}
		props = append(props, txt)
	}

	y = y + builder.Padding/2
	props = append(props, &svg2.Use{
		Height:    diagramHeight,
		Width:     builder.DiagramWidth,
		Transform: &svg2.Translate{Y: y},
		Href:      diagram.Id.Href(),
	})

	y = y + diagramHeight + builder.Padding

	return &svg2.G{
		Children: []any{
			&svg2.Rect{ // PROPERTIES BACKGROUND
				X:      builder.Padding + builder.LogoWidth + builder.Padding,
				Y:      -y / 2,
				Width:  builder.PropsWidth,
				Height: y,
				Fill:   svg2.Silver.Ptr(),
			},
			&svg2.G{
				Transform: &svg2.Translate{
					X: builder.Padding + builder.LogoWidth + builder.Padding + builder.Padding,
					Y: -y / 2,
				},
				Children: props,
			},
		},
	}, []any{diagram}
}

func (builder *LabelBuilder) BuildCableDiagram(a []parts.ConnectorCount, b []parts.ConnectorCount) *svg2.Svg {
	diagramWidth := (builder.Connector.Width+builder.Connector.Cable.Length)*2 + builder.Cable.Length

	cable := &svg2.Line{
		X1:          -builder.Cable.Length / 2,
		X2:          builder.Cable.Length / 2,
		Stroke:      svg2.Black.Ptr(),
		StrokeWidth: 4,
	}

	s := &svg2.Svg{
		Id:        "diagram",
		NameSpace: svg2.Namespace,
		Children: []any{
			cable,
		},
	}

	wiresA := builder.buildWires(a, true)
	wiresB := builder.buildWires(b, false)

	connectorsA, defsA, heightA := builder.buildConnectors(a)
	connectorsB, defsB, heightB := builder.buildConnectors(b)

	countsA := builder.buildConnectorCounts(a, true)
	countsB := builder.buildConnectorCounts(b, false)

	connectorsA.Transform = &svg2.Translate{
		X: -diagramWidth/2 + builder.Connector.Width,
		Y: -heightA / 2,
	}
	connectorsB.Transform = &svg2.Translate{
		X: diagramWidth / 2,
		Y: -heightB / 2,
	}

	wiresA.Transform = &svg2.Translate{
		X: -diagramWidth/2 + builder.Connector.Width,
		Y: -(heightA/2 - builder.Connector.Height/2),
	}
	wiresB.Transform = &svg2.Translate{
		X: builder.Cable.Length / 2,
		Y: -(heightB/2 - builder.Connector.Height/2),
	}

	countsA.Transform = &svg2.Translate{
		X: -diagramWidth/2 + builder.Connector.Width,
	}
	countsB.Transform = &svg2.Translate{
		X: diagramWidth/2 - builder.Connector.Width,
	}

	height := max(heightA, heightB)

	const countHeight float64 = 10
	// left side is longer and last element has count
	if len(a) > len(b) {
		if a[len(a)-1].Count > 1 {
			height += countHeight
		}
		// right side is longer and last element has count
	} else if len(b) > len(a) {
		if b[len(b)-1].Count > 1 {
			height += countHeight
		}
		// left and right equally long and last element has count
	} else if len(a) > 0 && (a[len(a)-1].Count > 1 || b[len(b)-1].Count > 1) {
		height += countHeight
	}

	s.Height = fmt.Sprintf("%v", height)
	s.ViewBox = &svg2.ViewBox{
		MinX:   -diagramWidth / 2,
		MinY:   -height / 2,
		Width:  diagramWidth,
		Height: height,
	}

	defs := make([]any, 0, len(defsA)+len(defsB))
	defs = append(defs, defsA...)
	defs = append(defs, defsB...)

	splitterA := builder.buildSplitter(cable.X1, heightA)
	splitterB := builder.buildSplitter(cable.X2, heightB)

	s.Children = append(s.Children,
		splitterA, splitterB,
		wiresA, wiresB,
		connectorsA, connectorsB,
		countsA, countsB,
		&svg2.Defs{Defs: defs},
	)

	return s
}

func (builder *LabelBuilder) buildWires(cs []parts.ConnectorCount, isLeft bool) *svg2.G {
	g := &svg2.G{Children: make([]any, 0, len(cs)*2)}

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
		g.Children = append(g.Children, &svg2.Line{
			X1:          x1,
			X2:          x2,
			Y1:          y,
			Y2:          y,
			Stroke:      svg2.Black.Ptr(),
			StrokeWidth: builder.Cable.Thickness,
		})
	}

	return g
}

func (builder *LabelBuilder) buildSplitter(x float64, connectorsHeight float64) *svg2.Line {
	return &svg2.Line{
		X1:          x,
		X2:          x,
		Y1:          -(connectorsHeight/2 - builder.Connector.Height/2 + builder.Connector.Cable.Thickness/2),
		Y2:          connectorsHeight/2 - builder.Connector.Height/2 + builder.Connector.Cable.Thickness/2,
		Stroke:      svg2.Black.Ptr(),
		StrokeWidth: 4,
	}
}

func (builder *LabelBuilder) buildConnectors(cs []parts.ConnectorCount) (g *svg2.G, defs []any, height float64) {
	defs = make([]any, len(cs))
	g = &svg2.G{
		Children: make([]any, len(cs)),
	}

	height = builder.Connector.Height*float64(len(cs)) + builder.Connector.Padding*float64(len(cs)-1)

	for i, e := range cs {
		defs[i] = e.Connector.Svg()
		g.Children[i] = &svg2.Use{
			Href:      e.Connector.Svg().Id.Href(),
			Width:     builder.Connector.Width,
			Height:    builder.Connector.Height,
			Y:         float64(i) * (builder.Connector.Height + builder.Connector.Padding),
			Transform: svg2.MirrorX,
		}
	}

	return g, defs, height
}

func (builder *LabelBuilder) buildConnectorCounts(cs []parts.ConnectorCount, isLeft bool) (g *svg2.G) {
	g = &svg2.G{}

	for i, e := range cs {
		if e.Count < 2 {
			continue
		}
		y := float64(i) * (builder.Connector.Height + builder.Connector.Padding)
		txt := &svg2.Text{
			Y:    y,
			Text: fmt.Sprintf("%dx", e.Count),
		}
		if !isLeft {
			txt.TextAnchor = svg2.TextAnchorEnd
		}
		g.Children = append(g.Children, txt)
	}

	return g
}
