package main

import (
	"encoding/xml"
	"fmt"
	FGD_GML_Schema "github.com/rhykw/fgd-gml-schema-go"
)

func main() {
	var sr FGD_GML_Schema.Dataset
	data := getData()
	err := xml.Unmarshal([]byte(data), &sr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", sr)
}

func getData() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<Dataset xsi:schemaLocation="http://fgd.gsi.go.jp/spec/2008/FGD_GMLSchema FGD_GMLSchema.xsd"
        xmlns:gml="http://www.opengis.net/gml/3.2"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        xmlns="http://fgd.gsi.go.jp/spec/2008/FGD_GMLSchema"
        gml:id="Dataset1">
  <gml:description>基盤地図情報メタデータ ID=fmdid:16-1001</gml:description>
  <gml:name>基盤地図情報ダウンロードデータ（GML版）</gml:name>
<AdmArea gml:id="K4_503005_1">
<fid>503005-40203-a-1-20160923</fid>
<lfSpanFr gml:id="K4_503005_1-1">
<gml:timePosition>2016-09-23</gml:timePosition>
</lfSpanFr>
<devDate gml:id="K4_503005_1-2">
<gml:timePosition>2016-09-24</gml:timePosition>
</devDate>
<orgGILvl>25000</orgGILvl>
<vis>表示</vis>
<area>
<gml:Surface gml:id="K4_503005_1-g" srsName="fguuid:jgd2011.bl">
<gml:patches>
<gml:PolygonPatch>
<gml:exterior>
<gml:Ring>
<gml:curveMember>
<gml:Curve gml:id="K4_503005_1-3">
<gml:segments>
<gml:LineStringSegment>
<gml:posList>
33.352839721 130.631333061
33.353220279 130.631023606
33.354276667 130.630163333
</gml:posList>
</gml:LineStringSegment>
</gml:segments>
</gml:Curve>
</gml:curveMember>
</gml:Ring>
</gml:exterior>
</gml:PolygonPatch>
</gml:patches>
</gml:Surface>
</area>
<type>郡市・東京都の区</type>
<name>久留米市</name>
<admCode>40203</admCode>
</AdmArea>
<AdmArea gml:id="K4_503005_2">
<fid>503005-40503-a-2-20160923</fid>
<lfSpanFr gml:id="K4_503005_2-1">
<gml:timePosition>2016-09-23</gml:timePosition>
</lfSpanFr>
<devDate gml:id="K4_503005_2-2">
<gml:timePosition>2016-09-24</gml:timePosition>
</devDate>
<orgGILvl>25000</orgGILvl>
<vis>表示</vis>
<area>
<gml:Surface gml:id="K4_503005_2-g" srsName="fguuid:jgd2011.bl">
<gml:patches>
<gml:PolygonPatch>
<gml:exterior>
<gml:Ring>
<gml:curveMember>
<gml:Curve gml:id="K4_503005_2-3">
<gml:segments>
<gml:LineStringSegment>
<gml:posList>
33.358162225 130.628290389
33.357792225 130.628445279
</gml:posList>
</gml:LineStringSegment>
</gml:segments>
</gml:Curve>
</gml:curveMember>
</gml:Ring>
</gml:exterior>
</gml:PolygonPatch>
</gml:patches>
</gml:Surface>
</area>
<type>町村・指定都市の区</type>
<name>大刀洗町</name>
<admCode>40503</admCode>
</AdmArea>
</Dataset>
`
}
