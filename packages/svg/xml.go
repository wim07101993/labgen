package svg

import "encoding/xml"

type XmlEncodable interface {
	Encode(w *xml.Encoder) error
}

func WriteObject(w *xml.Encoder, name string, attrs []xml.Attr, children []XmlEncodable) (err error) {
	xmlName := xml.Name{Local: name}
	start := xml.StartElement{Name: xmlName, Attr: attrs}

	err = w.EncodeToken(start)
	if err != nil {
		return
	}

	if len(children) > 0 {
		for _, f := range children {

			if err = f.Encode(w); err != nil {
				return
			}
		}
	}

	return w.EncodeToken(xml.EndElement{Name: xmlName})
}

func WriteCharData(w *xml.Encoder, name string, attrs []xml.Attr, value string) (err error) {
	xmlName := xml.Name{Local: name}
	start := xml.StartElement{Name: xmlName, Attr: attrs}

	if err = w.EncodeToken(start); err != nil {
		return
	}

	if err = w.EncodeToken(xml.CharData(value)); err != nil {
		return
	}

	return w.EncodeToken(xml.EndElement{Name: xmlName})
}
