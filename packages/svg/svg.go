package svg

const Namespace = "http://www.w3.org/2000/svg"

type Svg struct {
	XMLName   struct{} `xml:"svg"`
	Id        Id       `xml:"id,attr,omitempty"`
	NameSpace string   `xml:"xmlns,attr,omitempty"`
	ViewBox   *ViewBox `xml:"viewBox,attr,omitempty"`
	Width     string   `xml:"width,attr,omitempty"`
	Height    string   `xml:"height,attr,omitempty"`
	Children  []any    `xml:",innerxml"`
}

type Use struct {
	XMLName   struct{}    `xml:"use"`
	Href      string      `xml:"href,attr,omitempty"`
	Width     float64     `xml:"width,attr,omitempty"`
	Height    float64     `xml:"height,attr,omitempty"`
	X         float64     `xml:"x,attr,omitempty"`
	Y         float64     `xml:"y,attr,omitempty"`
	Transform Transformer `xml:"transform,attr,omitempty"`
	Fill      *Color      `xml:"fill,attr,omitempty"`
}

type TextAnchor string

var (
	TextAnchorStart TextAnchor = "start"
	TextAnchorEnd   TextAnchor = "end"
)

type Text struct {
	XMLName    struct{}    `xml:"text"`
	Id         Id          `xml:"id,attr,omitempty"`
	Y          float64     `xml:"y,attr,omitempty"`
	X          float64     `xml:"x,attr,omitempty"`
	Fill       *Color      `xml:"fill,attr,omitempty"`
	TextAnchor TextAnchor  `xml:"text-anchor,attr,omitempty"`
	FontWeight int         `xml:"font-weight,attr,omitempty"`
	FontSize   int         `xml:"font-size,attr,omitempty"`
	Transform  Transformer `xml:"transform,attr,omitempty"`
	Text       string      `xml:",chardata"`
}
type Style struct {
	XMLName struct{} `xml:"style"`
	Value   string   `xml:",chardata"`
}

type Rect struct {
	XMLName     struct{}    `xml:"rect"`
	Id          Id          `xml:"id,attr,omitempty"`
	Y           float64     `xml:"y,attr,omitempty"`
	X           float64     `xml:"x,attr,omitempty"`
	Width       float64     `xml:"width,attr,omitempty"`
	Height      float64     `xml:"height,attr,omitempty"`
	Fill        *Color      `xml:"fill,attr,omitempty"`
	Stroke      *Color      `xml:"stroke,attr,omitempty"`
	StrokeWidth float64     `xml:"stroke-width,attr,omitempty"`
	Transform   Transformer `xml:"transform,attr,omitempty"`
}

type Path struct {
	XMLName     struct{} `xml:"path"`
	Id          Id       `xml:"id,attr,omitempty"`
	D           string   `xml:"d,attr,omitempty"`
	Stroke      *Color   `xml:"stroke,attr,omitempty"`
	StrokeWidth float64  `xml:"stroke-width,attr,omitempty"`
	Fill        *Color   `xml:"fill,attr,omitempty"`
}

type Line struct {
	XMLName     struct{} `xml:"line"`
	X1          float64  `xml:"x1,attr,omitempty"`
	X2          float64  `xml:"x2,attr,omitempty"`
	Y1          float64  `xml:"y1,attr,omitempty"`
	Y2          float64  `xml:"y2,attr,omitempty"`
	Stroke      *Color   `xml:"stroke,attr,omitempty"`
	StrokeWidth float64  `xml:"stroke-width,attr,omitempty"`
}

type G struct {
	XMLName   struct{}    `xml:"g"`
	Id        Id          `xml:"id,attr,omitempty"`
	Transform Transformer `xml:"transform,attr,omitempty"`
	Children  []any
}

type Defs struct {
	XMLName struct{} `xml:"defs"`
	Defs    []any    `xml:",innerxml"`
}

type Circle struct {
	XMLName     struct{}    `xml:"circle"`
	Id          Id          `xml:"id,attr,omitempty"`
	CenterY     float64     `xml:"cy,attr,omitempty"`
	CenterX     float64     `xml:"cx,attr,omitempty"`
	Radius      float64     `xml:"r,attr,omitempty"`
	Width       float64     `xml:"width,attr,omitempty"`
	Height      float64     `xml:"height,attr,omitempty"`
	Fill        *Color      `xml:"fill,attr,omitempty"`
	Stroke      *Color      `xml:"stroke,attr,omitempty"`
	StrokeWidth float64     `xml:"stroke-width,attr,omitempty"`
	Transform   Transformer `xml:"transform,attr,omitempty"`
}
