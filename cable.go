package main

import (
	"fmt"

	"github.com/wim07101993/labgen/svg"
)

type Cable struct {
	Title           string           `json:"title"`
	MaxAmps         int              `json:"maxAmps,omitempty"`
	Volt            int              `json:"volt,omitempty"`
	Length          float64          `json:"length,omitempty"`
	ConnectorsSideA []ConnectorCount `json:"connectorsSideA,omitempty"`
	ConnectorsSideB []ConnectorCount `json:"connectorsSideB,omitempty"`
}

type ConnectorCount struct {
	Connector Connector `json:"connector,omitempty"`
	Count     int       `json:"count,omitempty"`
}

func (c *Cable) BackgroundColor() (svg.Color, error) {
	switch {
	case c.Length <= 1: //1
		return svg.White, nil
	case c.Length > 1 && c.Length <= 2: //1.5
		return svg.Green, nil
	case c.Length > 2 && c.Length <= 4: //3
		return svg.Red, nil
	case c.Length > 4 && c.Length <= 9: //5
		return svg.Black, nil
	case c.Length > 9 && c.Length <= 13: //10
		return svg.Blue, nil
	case c.Length > 13 && c.Length <= 18: //15
		return svg.Orange, nil
	case c.Length > 18 && c.Length <= 30: //20
		return svg.Cyan, nil
	case c.Length > 30: //50
		return svg.Magenta, nil
	default:
		return 0, fmt.Errorf("no color for length %v", c.Length)
	}
}

func (c *Cable) ForegroundColor() svg.Color {
	switch {
	case c.Length <= 1: //1
		return svg.Black
	default:
		return svg.White
	}
}

func (c *Cable) subTitle() string {
	if c.Volt != 0 && c.MaxAmps != 0 {
		return fmt.Sprintf("%dV %dA", c.Volt, c.MaxAmps)
	}
	if c.Volt != 0 {
		return fmt.Sprintf("%dV", c.Volt)
	}
	if c.MaxAmps != 0 {
		return fmt.Sprintf("%dA", c.MaxAmps)
	}
	return ""
}

func (c *Cable) LabelSvg() (*svg.Svg, error) {
	const height float64 = 60
	const width float64 = 60

	const titleFontSize = 4
	const lengthFontSize = 8

	background, err := c.BackgroundColor()
	if err != nil {
		return nil, err
	}
	foreground := c.ForegroundColor()

	subTitle := c.subTitle()

	titleTxt := &svg.Text{
		Id:       "title",
		FontSize: titleFontSize,
		Y:        float64(titleFontSize),
		Text:     c.Title,
	}

	diagram := BuildCableDiagram(c.ConnectorsSideA, c.ConnectorsSideB)

	props := []any{
		titleTxt,
		&svg.Use{
			Href: diagram.Id.Href(),
		},
	}

	if subTitle != "" {
		txt := &svg.Text{
			Id:        "subtitle",
			FontSize:  24,
			Transform: &svg.Translate{Y: titleTxt.Y + 24},
		}
		props = append(props, txt)
	}

	var propsHeight float64 = 15

	return &svg.Svg{
		NameSpace: svg.Namespace,
		ViewBox:   &svg.ViewBox{MinY: -height / 2, Width: width, Height: height},
		Width:     fmt.Sprintf("%fmm", width),
		Height:    fmt.Sprintf("%fmm", height),
		Children: []any{
			svg.Style{Value: "text { font-family: arial; }"},
			&svg.Rect{ // BACKGROUND
				Width:  width,
				Height: height,
				Y:      -height / 2,
				Fill:   background.Ptr(),
			},
			&svg.Rect{ // PROPERTIES BACKGROUND
				X:      width * 3 / 20,
				Y:      -propsHeight / 2,
				Width:  31,
				Height: propsHeight,
				Fill:   svg.Silver.Ptr(),
			},
			&svg.G{
				Transform: &svg.Translate{X: 11, Y: -propsHeight/2 + 1},
				Children:  props,
			},
			&svg.Text{ // LENGTH
				Id:         "length",
				Y:          float64(lengthFontSize)/2 - (float64(lengthFontSize) / 6),
				X:          58,
				Fill:       foreground.Ptr(),
				TextAnchor: svg.TextAnchorEnd,
				FontWeight: 70,
				FontSize:   lengthFontSize,
				Text:       fmt.Sprintf("%g", c.Length),
			},
			&svg.G{ // LOGO
				Transform: &svg.Translate{X: 1},
				Children: []any{
					&svg.Use{
						Href:      fmt.Sprintf("#%s", Logo.Id),
						Width:     width,
						Height:    height / 10,
						X:         -width / 2,
						Y:         -height / 10,
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

func (c *Cable) DiagramSvg() *svg.Svg {
	return BuildCableDiagram(c.ConnectorsSideA, c.ConnectorsSideB)
}

func BuildCableDiagram(a []ConnectorCount, b []ConnectorCount) *svg.Svg {
	const conWidth float64 = 30
	const conHeight float64 = 30
	const conWireLength float64 = 15
	const cableLength float64 = 90
	const cableWidth float64 = 4

	diagramHeight := conHeight * float64(max(len(a), len(b))) / 2
	diagramHeight = max(diagramHeight, conHeight)
	diagramWidth := (conWidth+conWireLength)*2 + cableLength

	cable := &svg.Line{
		X1:          -cableLength / 2,
		X2:          cableLength / 2,
		Stroke:      svg.Black.Ptr(),
		StrokeWidth: 4,
	}

	cableSplitAHeight := float64(len(a)-1) * conHeight
	cableSplitBHeight := float64(len(b)-1) * conHeight

	s := &svg.Svg{
		Id:        "diagram",
		NameSpace: svg.Namespace,
		ViewBox: &svg.ViewBox{
			MinX:   -diagramWidth / 2,
			MinY:   -diagramHeight / 2,
			Width:  diagramWidth,
			Height: diagramHeight,
		},
		Children: []any{
			cable,
			&svg.Line{
				X1:          cable.X1,
				X2:          cable.X1,
				Y1:          cableSplitAHeight / 2,
				Y2:          -cableSplitAHeight / 2,
				Stroke:      svg.Black.Ptr(),
				StrokeWidth: 4,
			},
			&svg.Line{
				X1:          cable.X2,
				X2:          cable.X2,
				Y1:          cableSplitBHeight / 2,
				Y2:          -cableSplitBHeight / 2,
				Stroke:      svg.Black.Ptr(),
				StrokeWidth: 4,
			},
		},
	}

	wiresA := buildWires(a, conWidth, conHeight, conWireLength, cableWidth, true)
	wiresA.Transform = &svg.Translate{X: -diagramWidth/2 + conWidth}
	wiresB := buildWires(b, conWidth, conHeight, conWireLength, cableWidth, false)
	wiresB.Transform = &svg.Translate{X: cableLength / 2}

	gA, defsA := buildConnectors(a, conWidth, conHeight)
	gA.Transform = &svg.Translate{
		X: -diagramWidth/2 + conWidth,
		Y: -diagramHeight / 2,
	}
	gB, defsB := buildConnectors(b, conWidth, conHeight)
	gB.Transform = &svg.Translate{
		X: diagramWidth / 2,
		Y: -diagramHeight / 2,
	}

	defs := make([]any, 0, len(defsA)+len(defsB))
	for _, def := range defsA {
		defs = append(defs, def)
	}
	for _, def := range defsB {
		defs = append(defs, def)
	}

	s.Children = append(s.Children,
		wiresA, wiresB,
		gA, gB,
		&svg.Defs{Defs: defs},
	)

	return s
}

func buildWires(cs []ConnectorCount, conWidth float64, conHeight float64, wireLength float64, wireWidth float64, isLeft bool) *svg.G {
	g := &svg.G{
		Children: make([]any, 0, len(cs)*2),
	}

	var x1 float64
	var x2 float64
	if isLeft {
		x1 = -conWidth / 2
		x2 = wireLength
	} else {
		x1 = 0
		x2 = conWidth/2 + wireLength
	}
	for i := range cs {
		y := float64(i) * conHeight
		g.Children = append(g.Children, &svg.Line{
			X1:          x1,
			X2:          x2,
			Y1:          y,
			Y2:          y,
			Stroke:      svg.Blue.Ptr(),
			StrokeWidth: wireWidth,
		})
	}

	return g
}
func buildConnectors(cs []ConnectorCount, conWidth float64, conHeight float64) (g *svg.G, defs []any) {
	defs = make([]any, len(cs))
	g = &svg.G{
		Children: make([]any, len(cs)),
	}

	for i, e := range cs {
		defs[i] = e.Connector.Svg()
		g.Children[i] = &svg.Use{
			Href:      e.Connector.Svg().Id.Href(),
			Width:     conWidth,
			Height:    conHeight,
			Y:         float64(i) * conHeight,
			Transform: svg.MirrorX,
		}
	}

	return g, defs
}
