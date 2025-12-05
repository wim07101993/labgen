package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed label_template.svg
var d6template string

//go:embed types/balanced-xlr.svg
var balancedXlr string

//go:embed types/combi-xlr-shuko.svg
var combiXlrShuko string

//go:embed types/combi-xlr,shuko-xlr,iec-c13.svg
var combiXlrShukoXlrIecC13 string

//go:embed types/dmx-3-pin-xlr.svg
var dmx3PinXlr string

//go:embed types/shuko-230V-16A.svg
var shuko230V16A string

//go:embed types/shuko-230V-20A.svg
var shuko230V20A string

//go:embed types/shuko-iec-c13.svg
var shukoIecC13 string

//go:embed types/unbalanced-jack.svg
var unbalancedJack string

//go:embed defs/iec-c13.svg
var iecC13 string

//go:embed defs/jack.svg
var jack string

//go:embed defs/shuko-f.svg
var shukoF string

//go:embed defs/shuko-m.svg
var shukoM string

//go:embed defs/xlr-f.svg
var xlrF string

//go:embed defs/xlr-m.svg
var xlrM string

var Colors = map[string]struct {
	Background string
	Foreground string
}{
	"1":   {"white", "black"},
	"1.5": {"green", "white"},
	"3":   {"red", "white"},
	"5":   {"black", "white"},
	"10":  {"blue", "white"},
	"15":  {"orange", "white"},
	"20":  {"cyan", "white"},
	"50":  {"magenta", "white"},
}

var labels = []Label{
	NewLabel("balanced-xlr", "1"),
	NewLabel("balanced-xlr", "1.5"),
	NewLabel("balanced-xlr", "3"),
	NewLabel("balanced-xlr", "5"),
	NewLabel("balanced-xlr", "10"),
	NewLabel("balanced-xlr", "15"),
	NewLabel("balanced-xlr", "50"),
	NewLabel("balanced-xlr", "50"),
	NewLabel("combi-xlr,shuko", "15"),
	NewLabel("combi-xlr,shuko", "20"),
	NewLabel("combi-xlr-shuko,xlr-iec-c13", "15"),
	NewLabel("dmx-3-pin-xlr", "10"),
	NewLabel("dmx-3-pin-xlr", "15"),
	NewLabel("shuko-230V-16A", "10"),
	NewLabel("shuko-230V-20A", "3"),
	NewLabel("shuko-230V-20A", "5"),
	NewLabel("shuko-230V-20A", "10"),
	NewLabel("shuko-230V-20A", "20"),
	NewLabel("shuko-iec-c13", "1"),
	NewLabel("shuko-iec-c13", "1.5"),
	NewLabel("shuko-iec-c13", "3"),
	NewLabel("shuko-iec-c13", "7.5"),
	NewLabel("unbalanced-jack", "1.5"),
	NewLabel("unbalanced-jack", "3"),
	NewLabel("unbalanced-jack", "5"),
	NewLabel("unbalanced-jack", "10"),
}

type Label struct {
	Length     string
	Type       string
	Properties string
	Background string
	Foreground string
	Defs       []string
}

func NewLabel(tp string, length string) Label {
	var properties string
	var defs []string
	switch tp {
	case "balanced-xlr":
		properties = balancedXlr
		defs = []string{xlrF, xlrM}
	case "combi-xlr,shuko":
		properties = combiXlrShuko
		defs = []string{xlrF, xlrM, shukoM, shukoF}
	case "combi-xlr-shuko,xlr-iec-c13":
		properties = combiXlrShukoXlrIecC13
		defs = []string{xlrF, xlrM, shukoM, iecC13}
	case "dmx-3-pin-xlr":
		properties = dmx3PinXlr
		defs = []string{xlrF, xlrM}
	case "shuko-230V-16A":
		properties = shuko230V16A
		defs = []string{shukoM, shukoF}
	case "shuko-230V-20A":
		properties = shuko230V20A
		defs = []string{shukoM, shukoF}
	case "shuko-iec-c13":
		properties = shukoIecC13
		defs = []string{shukoM, iecC13}
	case "unbalanced-jack":
		properties = unbalancedJack
		defs = []string{jack}
	default:
		panic(fmt.Errorf("no match for properties %s", tp))
	}
	colors := Colors[length]
	return Label{
		Length:     length,
		Type:       tp,
		Properties: properties,
		Background: colors.Background,
		Foreground: colors.Foreground,
		Defs:       defs,
	}
}

func main() {
	for _, label := range labels {
		svg := strings.ReplaceAll(d6template, "$BACKGROUND", label.Background)
		svg = strings.ReplaceAll(svg, "$FOREGROUND", label.Foreground)
		svg = strings.ReplaceAll(svg, "$PROPERTIES", label.Properties)
		svg = strings.ReplaceAll(svg, "$LENGTH", label.Length)
		svg = strings.ReplaceAll(svg, "$DEFS", strings.Join(label.Defs, "\r\n\r\n"))

		fileName := label.Type + "-" + label.Length + ".svg"
		if err := os.WriteFile(fileName, []byte(svg), 0777); err != nil {
			panic(err)
		}
	}
}
