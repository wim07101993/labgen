package parts

import (
	"strconv"
	"strings"

	"github.com/wim07101993/labgen/packages/svg"
	"github.com/wim07101993/labgen/packages/svg_components"
)

var connectors = map[string]*svg.Svg{
	string(IecC13) + "-3-female": svg_components.IecC13,
	string(Shuko) + "-3-female":  svg_components.ShukoF,
	string(Shuko) + "-3-male":    svg_components.ShukoM,
	string(Jack) + "-2-male":     svg_components.JackTs,
	string(Jack) + "-3-male":     svg_components.JackTrs,
	string(Xlr) + "-3-female":    svg_components.Xlr3PF,
	string(Xlr) + "-3-male":      svg_components.Xlr3PM,
	//string(Xlr) + "-5-female":    xlr5Female,
	//string(Xlr) + "-5-male":      xlr5Male,
}

type ConnectorName string

const (
	Xlr    ConnectorName = "xlr"
	Jack   ConnectorName = "jack"
	IecC13 ConnectorName = "iec_c13"
	Shuko  ConnectorName = "shuko"
)

type Connector struct {
	Name      ConnectorName `json:"name,omitempty"`
	PoleCount int           `json:"poleCount,omitempty"`
	IsMale    bool          `json:"isMale"`
}

func (c *Connector) Svg() *svg.Svg {
	href := strings.Builder{}
	href.WriteString(string(c.Name))
	href.WriteRune('-')
	href.WriteString(strconv.Itoa(c.PoleCount))
	if c.IsMale {
		href.WriteString("-male")
	} else {
		href.WriteString("-female")
	}
	return connectors[href.String()]
}
