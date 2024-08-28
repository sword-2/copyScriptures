package tanach


//This is the structure of top level storage for tanach.us xml files. The details may be divided to separate files
import (
	"encoding/xml"
)

//This is top level variable used to Unmarshall XML file into
var tanachUsXml1 struct {
	XMLName xml.Name `xml:"Tanach"`
	TeiHeader teiHeader `xml:"teiHeader"`
	Tanach Tanach `xml:"tanach"`
	Notes Notes `xml:"notes"`
}


var xmlAsString string //temporary storage of xml file prior to unmarshalling
