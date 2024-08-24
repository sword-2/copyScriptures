package tanach

import (
	"encoding/xml"
)

//details in structure for tanach tags

type Notes struct {
	XMLName xml.Name `xml:"notes"`
	Note []Note `xml:"note"`
}

type Note struct {
	Code string  `xml:"code"`
    Gccode string    `xml:"gccode"`
    Note string    `xml:"note"`
}
