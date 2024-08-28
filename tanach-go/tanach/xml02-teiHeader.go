package tanach


import (
	"encoding/xml"
)



//details in structure for teiHeader tags

//name element used more than once
type Name []struct {
	XMLName xml.Name `xml:"name"`
	Type string `xml:"type,attr"`
}

type BibleItem struct {
	XMLName xml.Name `xml:"bibleItem"`
	Title []string `xml:"title"`
	Editor []string  `xml:"editor"`
	Edition string  `xml:"edition"`
	Imprint struct {
		XMLName xml.Name `xml:"imprint"`
		Publisher string   `xml:"publisher"`
		PubPlace []string   `xml:"pubPlace"`
		Date string   `xml:"date"`
	}
	Idno struct {
		XMLName xml.Name `xml:"idno"`
		Type string `xml:"type,attr"`
	}
}

type Correction struct {
    Citation string `xml:"citation"`
    Description  string    `xml:"description"`
    Author  string    `xml:"author"`
    Filedate  string    `xml:"filedate"`
    Date  string    `xml:"date"`
    N  string    `xml:"n"`
}


	type teiHeader struct {
		XMLName xml.Name `xml:"teiHeader"`
		FileDesc struct {
			XMLName xml.Name `xml:"fileDesc"`
			TitleStmt struct {
				XMLName xml.Name `xml:"titleStmt"`
				//Title []string `xml:"title"`
				Title []struct {
					XMLName xml.Name `xml:"title"`
					Level string `xml:"level,attr"`
					Type string `xml:"type,attr"`
				}
				Editor []string `xml:"editor"`
				RespStmt struct {
					XMLName xml.Name `xml:"respStmt"`
					Resp string `xml:"resp"`
					Name []Name `xml:"name"`
					//Name []string `xml:"name"`
				}
			}
			EditionStmt struct {
				XMLName xml.Name `xml:"editionStmt"`
				Edition struct {
					XMLName xml.Name `xml:"edition"`
					Version string `xml:"version"`
					Date string `xml:"date"`
					Build string `xml:"build"`
					BuildDateTime string `xml:"buildDateTime"`
				}
				RespStmt struct {
					XMLName xml.Name `xml:"respStmt"`
					Resp []string  `xml:"resp"`
				}
			}
			Extent string `xml:"extent"`
			PublicationStmt struct {
				XMLName xml.Name `xml:"publicationStmt"`
				Distributor struct {
					XMLName xml.Name `xml:"distributor"`
					Name []Name `xml:"name"`
					//Name []string  `xml:"name"`
					//Availability string  `xml:"availability"`
					Availability struct {
						XMLName xml.Name `xml:"availability"`
						Status string `xml:"status,attr"`
					}
			}
			NotesStmt struct {
				XMLName xml.Name `xml:"notesStmt"`
				Note string  `xml:"note"`
				Correction []Correction  `xml:"correction"`
			}
		}
		SourceDesc struct {
			XMLName xml.Name `xml:"sourceDesc"`
			BibleItem []BibleItem `xml:"bibleItem"`
		}
	}
	EncodingDesc string   `xml:"encodingDesc"`
	ProfileDesc struct {
		XMLName xml.Name `xml:"profileDesc"`
		Creation string  `xml:"creation"`
		Date string  `xml:"date"`
		LangUsage struct {
			XMLName xml.Name `xml:"langUsage"`
			//Language string   `xml:"language"`
			Language struct {
				XMLName xml.Name `xml:"language"`
				Ident string   `xml:"ident,attr"`
			}
		}
	}
}


//var Tanach T //global variable for r(R)esult.

var htmlTable string //same as r above, but string format. string format is easier to write out to file since no need to loop through all the variables in r.
