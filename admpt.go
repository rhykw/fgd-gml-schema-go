// Code generated by zek 0.1.23; DO NOT EDIT.
package fgdgmlschema

import "encoding/xml"

// AdmPt was generated automatically by zek 0.1.23. DO NOT EDIT.
type AdmPt struct {
	XMLName  xml.Name `xml:"AdmPt"`
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
	Vis      string `xml:"vis"`
	Pos      struct {
		Text  string `xml:",chardata"`
		Point struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id,attr"`
			SrsName string `xml:"srsName,attr"`
			Pos     string `xml:"pos"`
		} `xml:"Point"`
	} `xml:"pos"`
	Type    string `xml:"type"`
	Name    string `xml:"name"`
	AdmCode string `xml:"admCode"`
}