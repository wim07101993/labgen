package svg

import (
	"encoding/xml"
	"strings"
)

var (
	MirrorY = &Scale{X: 1, Y: -1}
	MirrorX = &Scale{X: -1, Y: 1}
)

type Transformer interface {
	MarshalXMLAttr(name xml.Name) (xml.Attr, error)
	String() string
}

type Combiner []Transformer

func (c Combiner) String() string {
	b := strings.Builder{}
	for _, t := range c {
		b.WriteString(t.String())
	}
	return b.String()
}

func (c Combiner) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: c.String(),
	}, nil
}
