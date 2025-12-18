package svg

import (
	"encoding/xml"
	"fmt"
)

type ViewBox struct {
	MinX   float64
	MinY   float64
	Width  float64
	Height float64
}

func (vb ViewBox) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: vb.String(),
	}, nil
}

func (vb ViewBox) String() string {
	return fmt.Sprintf("%f %f %f %f", vb.MinX, vb.MinY, vb.Width, vb.Height)
}
