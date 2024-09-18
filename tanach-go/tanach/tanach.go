// Tanach is assumed to mean Hebrew bible
//
//File xml01* has the top level structure.
//
//File xml02* has structure details for teiHeader tags
//
//File xml03* has structure details for tanach tags
//
//File xml04* has structure details for notes tags
package tanach

import (
	"fmt"
	"encoding/xml"
	//"strings" //from tutorial on reading xml
	"os"
    //"encoding/json"
	//"sort"
	 //"strconv"
)

//function check is an attempt to handle possible errors instead of having a bunch of error handlers
func check(e error) {
	if e != nil {
        panic(e)
    } //else {
	//	fmt.Println("\nno error for e")
	//}
}

//This function is just left over from a previous project and to be recycled or deleted
func RemoveFile(path *string) { 
	fmt.Println("\ntanach.RemoveFile() called on path: ", *path)
	/*
	if _, err := os.Stat(*path); err == nil {
		check(err)
		fmt.Println("\nRemoveOldL14(): old .l14 exists planning to remove it")
		e := os.Remove(*path)
		check(e)
	}
	*/
}

//A function to do the Go method of transforming json to program variables. The Go term is unmarshal
func ReadUnmarshalXml(path *string) {
	fmt.Println("\ntanach.ReadFile() called for path: ", *path)

	xmlAsString, err := os.ReadFile(*path)
	check(err)
	//fmt.Printf("\nxmlAsString debugging=%s", xmlAsString) //Successfully read in file to string

	if err := xml.Unmarshal([]byte(xmlAsString), &tanachUsXml1); err != nil {
		//fmt.Println("\nsome error duing Unmarshal=", err)
		panic(err)
	}

}

//Initial printing of selected tei Header tags to verify the unmarshal worked.
func PrintTeiHeader() {
	fmt.Println("following are selected tags from within the teiHeader ...")
	fmt.Println("\ttitle ...")
	for _, value := range tanachUsXml1.TeiHeader.FileDesc.TitleStmt.Title {
			fmt.Printf("\t\t%s\n", value)
	}

	fmt.Println("\tCreation ...")
	fmt.Println("\t\ttanachUsXml1.TeiHeader.ProfileDesc.Creation=", tanachUsXml1.TeiHeader.ProfileDesc.Creation)

}


//Initial printing of selected scriptures to verify the unmarshal worked.
func PrintTanach() {
	fmt.Println("following are selected tags from within the tanach ...")
	fmt.Println("\ttanachUsXml1.Tanach.Book.Names.Name=", tanachUsXml1.Tanach.Book.Names.Name)

	fmt.Println("\tverses for each chapter ...")
	for index, value := range tanachUsXml1.Tanach.Book.Chapter {
			fmt.Printf("\t\t%d: %s; %s\n", index + 1, value.Vs, value.V[0])
			var tmp1 string = ""
			for _, v2 := range value.V {
				tmp1 += v2.W[0] + " "
			}
			println(tmp1)
	}
}


//Initial printing of selected notes to verify the unmarshal worked.
func PrintNotes() {
	fmt.Println("following are selected tags from within the notes ...")
	fmt.Println("\ttanachUsXml1.Notes.Note[0].Code=", tanachUsXml1.Notes.Note[0].Code)

	fmt.Println("\tcodes for each note ...")
	for index, value := range tanachUsXml1.Notes.Note {
			fmt.Printf("\t\t%d: %s\n", index + 1, value.Code)
	}
}


//This function is just leftover from previous project waiting to be recycled or deleted.
func SortStruct(sortCode *int) {
	fmt.Println("\ntanach.SortStruct() called for int:", *sortCode)

	/*
	// Sort by one way by removing comments
	if *sortCode == 1 { 
		fmt.Println("SortStruct() sortCode == 1")
		sort.Slice(r.Agencies, func(i, j int) bool {
			return r.Agencies[i].Agency < r.Agencies[j].Agency
		})
	} else if *sortCode == 2 {
		fmt.Println("SortStruct() sortCode == 2")
		sort.Slice(r.Agencies, func(i, j int) bool {
			return r.Agencies[i].AgencyRef < r.Agencies[j].AgencyRef
		})
	} else {
		//fmt.Println("SortStruct() received unknown sortCode.")
		panic("SortStruct() received unknown sortCode.")
	}
	*/
}

//Another function leftover from previous project copy, to be recycled or deleted.
func PrintStruct() {
	fmt.Println("\ntanach.PrintStruct() called")

	/*
	htmlTable = "" //clear global variable

	//metadata
		htmlTable += "\n<p>source file:" + r.Filename + "</p>"     //fmt.Printf("\nr.Filename=%s", r.Filename)
		htmlTable += "\n<p>version:" + r.Version + "</p>"    //fmt.Printf("\nr.Version=%s", r.Version)
		htmlTable += "\n<p>comments:" + r.Comments + "</p>"    //fmt.Printf("\nr.Comments=%s", r.Comments)

	//table setup
		htmlTable += "\n<table border=1><caption>Agencies and liaisons as processed by Go language (Google)</caption>"
		htmlTable += "\n\t<tr><th>Agency</th>"
		htmlTable += "\n\t<th>Legal Ref</th>"
		htmlTable += "\n\t<th>liaisons</th>"
		htmlTable += "\n\t<th>repositories</th>"
		htmlTable += "\n\t<th>lastModified</th>"
		htmlTable += "\n\t<th>offboarding</th>"
		htmlTable += "\n</tr>"

	//fmt.Printf("\nr.Agencies=%s", r.Agencies)

	//loop thru agencies one at a time
	for i := 0; i < len(r.Agencies); i++ {
		htmlTable += "\n<tr>"
		//fmt.Printf("\na.Agency=%s; LegalRef=%s", r.Agencies[i].Agency, r.Agencies[i].AgencyRef)
		htmlTable += "\n\t<td>" + r.Agencies[i].Agency + "</td>" //name
		htmlTable += "\n\t<td>" + r.Agencies[i].AgencyRef + "</td>" //legal ref

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Liaisons); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Liaisons[j].Name + "; " +  r.Agencies[i].Liaisons[j].Email + "; " +  r.Agencies[i].Liaisons[j].Phone
				if j != len(r.Agencies[i].Liaisons) { htmlTable += "<br>" }
			}
			htmlTable += "\n\t</td>"

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Repos); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Repos[j].Name + "; " +  r.Agencies[i].Repos[j].Url
				if j != len(r.Agencies[i].Repos) { htmlTable += "<br>" }
			}
		htmlTable += "\n\t</td>"

		htmlTable += "\n\t<td>" + r.Agencies[i].LastModified + "</td>"

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Offboarding); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Offboarding[j].Title + "; " +  r.Agencies[i].Offboarding[j].Url
				if j != len(r.Agencies[i].Offboarding) { htmlTable += "<br>" }
			}
		htmlTable += "\n\t</td>"

		htmlTable += "\n</tr>"
	}
		htmlTable += "\n</table>"
	*/
}

func WriteFile(path *string) { //write data to a .l14 file so a web page can load it
	fmt.Println("\ntanach.WriteFile() was called for path: ", *path)
	/*
	err := os.WriteFile(*path, []byte(htmlTable), 0644)
	check(err)
	*/
}
