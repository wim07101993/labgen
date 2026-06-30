package main

import (
	_ "embed"
	"encoding/xml"
	"os"

	"github.com/wim07101993/labgen/internal"
	builder2 "github.com/wim07101993/labgen/internal/parts"
)

func main() {
	builder := internal.LabelBuilder{
		TitleFontSize:    4,
		SubtitleFontSize: 3,
		LengthFontSize:   8,
		LabelHeight:      70,
		LabelWidth:       70,
		PropsWidth:       36,
		DiagramWidth:     25,
		LogoWidth:        6,
		Padding:          2,
		Connector: &internal.ConnectorCfg{
			Width:   30,
			Height:  30,
			Padding: 10,
			Cable: &internal.CableCfg{
				Thickness: 4,
				Length:    20,
			},
		},
		Cable: &internal.CableCfg{
			Thickness: 4,
			Length:    60,
		},
	}

	xlrCable := builder2.Cable{
		Title:   "Balanced Combi",
		MaxAmps: 16,
		Volt:    230,
		Length:  1.5,
		ConnectorsSideA: []builder2.ConnectorCount{
			{
				Connector: builder2.Connector{
					Name:      builder2.Xlr,
					PoleCount: 3,
					IsMale:    false,
				},
			},
			{
				Connector: builder2.Connector{
					Name:      builder2.Shuko,
					PoleCount: 3,
					IsMale:    false,
				},
				Count: 3,
			},
		},
		ConnectorsSideB: []builder2.ConnectorCount{
			{
				Connector: builder2.Connector{
					Name:      builder2.Xlr,
					PoleCount: 3,
					IsMale:    true,
				},
			},
			{
				Connector: builder2.Connector{
					Name:      builder2.Shuko,
					PoleCount: 3,
					IsMale:    true,
				},
			},
		},
	}
	svg, err := builder.BuildLabel(xlrCable)
	if err != nil {
		panic(err)
	}

	indented, err := xml.MarshalIndent(svg, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("xlrcable.svg", indented, 777)
	if err != nil {
		panic(err)
	}
}
