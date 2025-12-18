package main

import (
	_ "embed"
	"encoding/xml"
	"os"
)

func main() {
	xlrCable := Cable{
		Title:  "Balanced",
		Length: 10,
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
	svg := xlrCable.DiagramSvg()

	indented, err := xml.MarshalIndent(svg, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("xlrcable.svg", indented, 777)
	if err != nil {
		panic(err)
	}
}
