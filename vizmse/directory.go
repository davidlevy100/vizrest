package vizmse

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Base    string   `xml:"base,attr"`
	ID      string   `xml:"id"`
	Title   string   `xml:"title"`
	Updated string   `xml:"updated"`
	Link    struct {
		Text string `xml:",chardata"`
		Href string `xml:"href,attr"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
	} `xml:"link"`
	Entry []struct {
		Text     string `xml:",chardata"`
		Xmlns    string `xml:"xmlns,attr"`
		Viz      string `xml:"viz,attr"`
		Category struct {
			Text   string `xml:",chardata"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"category"`
		Title  string `xml:"title"`
		Author struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
		} `xml:"author"`
		Updated string `xml:"updated"`
		ID      string `xml:"id"`
		Summary string `xml:"summary"`
		Link    []struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Empty string `xml:"empty"`
	} `xml:"entry"`
}
