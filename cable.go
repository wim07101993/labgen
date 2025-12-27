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
