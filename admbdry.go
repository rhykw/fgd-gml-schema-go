// Code generated by zek 0.1.23; DO NOT EDIT.
package fgdgmlschema

import "encoding/xml"

// AdmBdry was generated automatically by zek 0.1.23. DO NOT EDIT.
type AdmBdry struct {
	XMLName  xml.Name `xml:"AdmBdry"`
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
	Loc      Loc `xml:"loc"`
	Type string `xml:"type"`
}
