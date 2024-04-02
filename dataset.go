package fgdgmlschema

import "encoding/xml"

type Dataset struct {
	XMLName        xml.Name `xml:"Dataset"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Gml            string   `xml:"gml,attr"`
	Xsi            string   `xml:"xsi,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	ID             string   `xml:"id,attr"`
	Description    string   `xml:"description"`
	Name           string   `xml:"name"`
        AdmArea   []AdmArea   `xml:"AdmArea"`
        AdmBdry   []AdmBdry   `xml:"AdmBdry"`
        AdmPt     []AdmPt     `xml:"AdmPt"`
        BldA      []BldA      `xml:"BldA"`
        BldL      []BldL      `xml:"BldL"`
        Cntr      []Cntr      `xml:"Cntr"`
        CommBdry  []CommBdry  `xml:"CommBdry"`
        CommPt    []CommPt    `xml:"CommPt"`
        ElevPt    []ElevPt    `xml:"ElevPt"`
        GCP       []GCP       `xml:"GCP"`
        RailCL    []RailCL    `xml:"RailCL"`
        RdCompt   []RdCompt   `xml:"RdCompt"`
        RdEdg     []RdEdg     `xml:"RdEdg"`
        WA        []WA        `xml:"WA"`
        WL        []WL        `xml:"WL"`
        WStrA     []WStrA     `xml:"WStrA"`
        WStrL     []WStrL     `xml:"WStrL"`	
} 
