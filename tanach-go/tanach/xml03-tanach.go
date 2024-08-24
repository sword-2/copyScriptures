package tanach


import (
	"encoding/xml"
)

//details in structure for tanach tags

type Tanach struct {
	XMLName xml.Name `xml:"tanach"`
	Book struct {
		XMLName xml.Name `xml:"book"`
		Names struct {
			XMLName xml.Name `xml:"names"`
			Name string   `xml:"name"`
			Abbrev []string   `xml:"abbrev"`
			Number string   `xml:"number"`
			Filename string   `xml:"filename"`
			Hebrewname string   `xml:"hebrewnanme"`
		}
		Chapter []Chapter `xml:"c"`
		Vs string `xml:"vs"`  //verse count for whole book
		Cs string `xml:"cs"`  //chapter count
	}
}

type Chapter struct {
	XMLName xml.Name `xml:"c"`
	V []Verse  `xml:"v"`
    Vs  string `xml:"vs"` //verse count for this chapter
}

type Verse struct {
		XMLName xml.Name `xml:"v"`
		W []string `xml:"w"` //each individual word
}
