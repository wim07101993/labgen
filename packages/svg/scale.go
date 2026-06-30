package svg

import (
	"encoding/xml"
	"fmt"
)

type Scale struct {
	X float64
	Y float64
}

func (s *Scale) String() string {
	if s.Y == 0 {
		return fmt.Sprintf("scale(%f)", s.X)
	}
	return fmt.Sprintf("scale(%f %f)", s.X, s.Y)
}

func (s *Scale) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: s.String(),
	}, nil
}
