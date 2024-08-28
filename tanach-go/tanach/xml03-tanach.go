package tanach


import (
	"encoding/xml"
)


type Verse struct {
		XMLName xml.Name `xml:"v"`
		N string `xml:"n,attr"`  //verse number
		W []string `xml:"w"` //array of words
}

type Chapter struct {
	XMLName xml.Name `xml:"c"`
	N string `xml:"n,attr"`  //chapter number
	V []Verse  `xml:"v"`
    Vs  string `xml:"vs"` //verse count for this chapter
}


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

