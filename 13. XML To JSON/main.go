package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
)

type Developer struct {
	XMLName   xml.Name `xml:"developer" json:"-"`
	ID        int      `xml:"id" json:"id"`
	FirstName string   `xml:"firstname" json:"firstName"`
	LastName  string   `xml:"lastname" json:"lastName"`
	UserName  string   `xml:"username" json:"userName"`
}

type Developers struct {
	XMLName    xml.Name    `xml:"developers" json:"-"`
	Developers []Developer `xml:"developer" json:"developers"`
}

func (d Developer) String() string {
	return fmt.Sprintf(
		"\t ID : %d - FirstName : %s - LastName : %s - UserName : %s",
		d.ID,
		d.FirstName,
		d.LastName,
		d.UserName,
	)
}

func main() {
	in := flag.String("in", "developers.xml", "XML file to read")
	out := flag.String("out", "developers.json", "JSON file to write")
	indent := flag.Bool("indent", true, "indent the JSON output")
	flag.Parse()

	xmlData, err := os.ReadFile(*in)
	if err != nil {
		log.Fatalf("Reading '%s' failed: %v", *in, err)
	}

	var devs Developers
	if err := xml.Unmarshal(xmlData, &devs); err != nil {
		log.Fatalf("Parsing '%s' failed: %v", *in, err)
	}

	if len(devs.Developers) == 0 {
		log.Fatalf("No <developer> elements found in '%s'.", *in)
	}

	// Write XML on screen
	for _, developer := range devs.Developers {
		fmt.Println(developer)
	}

	// Convert to JSON
	var jsonData []byte
	if *indent {
		jsonData, err = json.MarshalIndent(devs.Developers, "", "  ")
	} else {
		jsonData, err = json.Marshal(devs.Developers)
	}
	if err != nil {
		log.Fatalf("Encoding JSON failed: %v", err)
	}

	// Write JSON on screen
	fmt.Printf("\n%s\n", jsonData)

	// Write to JSON file
	if err := os.WriteFile(*out, jsonData, 0o644); err != nil {
		log.Fatalf("Writing '%s' failed: %v", *out, err)
	}

	fmt.Printf("\nConverted %d developer(s) from '%s' to '%s'.\n", len(devs.Developers), *in, *out)
}
