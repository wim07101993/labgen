package svg

import (
	"encoding/xml"
	"fmt"
)

type Color int

const (
	White   Color = 0xFFFFFF
	Green   Color = 0x00FF00
	Red     Color = 0xFF0000
	Black   Color = 0x0000000
	Blue    Color = 0x0000FF
	Orange  Color = 0xFFA500
	Cyan    Color = 0x00FFFF
	Magenta Color = 0xFF00FF
	Silver  Color = 0xC0C0C0
)

func (c Color) String() string {
	return fmt.Sprintf("#%06x", int(c))
}

func (c Color) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{
		Name:  name,
		Value: c.String(),
	}, nil
}

func (c Color) Ptr() *Color {
	return &c
}
