// Code generated by zek 0.1.23; DO NOT EDIT.
package fgdgmlschema

import "encoding/xml"

// BldL was generated automatically by zek 0.1.23. DO NOT EDIT.
type BldL struct {
	XMLName  xml.Name `xml:"BldL"`
	Text     string   `xml:",chardata"`
	ID       string   `xml:"id,attr"`
	Fid      string   `xml:"fid"`
	LfSpanFr struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		TimePosition string `xml:"timePosition"`
	} `xml:"lfSpanFr"`
	DevDate struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id,attr"`
		TimePosition string `xml:"timePosition"`
	} `xml:"devDate"`
	OrgGILvl string `xml:"orgGILvl"`
	Loc      struct {
		Text  string `xml:",chardata"`
		Curve struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id,attr"`
			SrsName  string `xml:"srsName,attr"`
			Segments struct {
				Text              string `xml:",chardata"`
				LineStringSegment struct {
					Text    string `xml:",chardata"`
					PosList string `xml:"posList"`
				} `xml:"LineStringSegment"`
			} `xml:"segments"`
		} `xml:"Curve"`
	} `xml:"loc"`
	Type string `xml:"type"`
}
