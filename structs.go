package vizrest

import "encoding/xml"

type Service struct {
	XMLName   xml.Name  `xml:"service"`
	Text      string    `xml:",chardata"`
	Xmlns     string    `xml:"xmlns,attr"`
	Atom      string    `xml:"atom,attr"`
	Viz       string    `xml:"viz,attr"`
	Workspace Workspace `xml:"workspace"`
}

type Workspace struct {
	Text       string     `xml:",chardata"`
	Title      string     `xml:"title"`
	Collection Collection `xml:"collection"`
}

type Collection []struct {
	Text       string     `xml:",chardata"`
	Href       string     `xml:"href,attr"`
	Title      string     `xml:"title"`
	Categories Categories `xml:"categories"`
}

type Categories struct {
	Text     string   `xml:",chardata"`
	Category Category `xml:"category"`
}

type Category struct {
	Text   string `xml:",chardata"`
	Term   string `xml:"term,attr"`
	Scheme string `xml:"scheme,attr"`
}
