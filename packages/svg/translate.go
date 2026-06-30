package svg

import (
	"encoding/xml"
	"fmt"
)

type Translate struct {
	X float64
	Y float64
}

func (t *Translate) String() string {
	if t.Y == 0 {
		return fmt.Sprintf("translate(%f)", t.X)
	}
	return fmt.Sprintf("translate(%f %f)", t.X, t.Y)
}

func (t *Translate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: t.String(),
	}, nil
}
