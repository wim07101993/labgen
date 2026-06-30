package svg

import (
	"encoding/xml"
	"fmt"
)

type Rotate struct {
	Degrees float64
	X       float64
	Y       float64
}

func (r *Rotate) String() string {
	if r.X == 0 && r.Y == 0 {
		return fmt.Sprintf("rotate(%f)", r.Degrees)
	}
	return fmt.Sprintf("rotate(%f %f %f)", r.Degrees, r.X, r.Y)
}

func (r *Rotate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: r.String(),
	}, nil
}
