package main

import (
	_ "embed"
	"encoding/xml"
	"os"
)

func main() {
	builder := LabelBuilder{
		TitleFontSize:    4,
		SubtitleFontSize: 3,
		LengthFontSize:   8,
		LabelHeight:      70,
		LabelWidth:       70,
		PropsWidth:       36,
		LogoWidth:        6,
		Padding:          2,
		Connector: ConnectorCfg{
			Width:   30,
			Height:  30,
			Padding: 10,
			Cable: CableCfg{
				Thickness: 4,
				Length:    15,
			},
		},
		Cable: CableCfg{
			Thickness: 4,
			Length:    60,
		},
	}

	xlrCable := Cable{
		Title:   "Balanced Combi",
		MaxAmps: 16,
		Volt:    230,
		Length:  1.5,
		ConnectorsSideA: []ConnectorCount{
			{
				Connector: Connector{
					Name:      Xlr,
					PoleCount: 3,
					IsMale:    false,
				},
			},
			{
				Connector: Connector{
					Name:      Shuko,
					PoleCount: 3,
					IsMale:    false,
				},
			},
		},
		ConnectorsSideB: []ConnectorCount{
			{
				Connector: Connector{
					Name:      Xlr,
					PoleCount: 3,
					IsMale:    true,
				},
			},
			{
				Connector: Connector{
					Name:      Shuko,
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
