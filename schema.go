package fgdgmlschema

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// Pos position
type Pos [2]float32

// PosList position list
type PosList []Pos

// Area strunct of area
type Area struct {
	PosList PosList `xml:"http://www.opengis.net/gml/3.2 Surface>patches>PolygonPatch>exterior>Ring>curveMember>Curve>segments>LineStringSegment>posList"`
}

// AdmArea struct of AdmArea
type AdmArea struct {
	OrgGILvl string `xml:"orgGILvl"`
	Fid      string `xml:"fid"`
	Vis      string `xml:"vis"`
	Name     string `xml:"name"`
	Code     string `xml:"admCode"`
	Area     Area   `xml:"area"`
}

// Dataset struct of Dataset
type Dataset struct {
	Data []AdmArea `xml:"AdmArea"`
}

// UnmarshalXML UnmarshalXML for PosList
func (pl *PosList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	p := struct {
		Content string `xml:",chardata"`
	}{}
	if err := d.DecodeElement(&p, &start); err != nil {
		return err
	}

	t := strings.Fields(strings.TrimSpace(p.Content))
	*pl = PosList{}
	for i := 0; i < len(t); i++ {
		y, _ := strconv.ParseFloat(t[i], 32)
		x, _ := strconv.ParseFloat(t[i+1], 32)
		*pl = append(*pl, Pos{float32(x), float32(y)})
		i++
	}
	return nil
}
