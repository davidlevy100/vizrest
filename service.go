package main

import "encoding/xml"

type Service struct {
	XMLName   xml.Name  `xml:"service" json:"service,omitempty"`
	Text      string    `xml:",chardata" json:"text,omitempty"`
	Xmlns     string    `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Atom      string    `xml:"atom,attr" json:"atom,omitempty"`
	Viz       string    `xml:"viz,attr" json:"viz,omitempty"`
	Workspace Workspace `xml:"workspace" json:"workspace,omitempty"`
}

type Workspace struct {
	Text       string     `xml:",chardata" json:"text,omitempty"`
	Title      string     `xml:"title" json:"title,omitempty"`
	Collection Collection `xml:"collection" json:"collection,omitempty"`
}

type Collection []struct {
	Text       string     `xml:",chardata" json:"text,omitempty"`
	Href       string     `xml:"href,attr" json:"href,omitempty"`
	Title      string     `xml:"title" json:"title,omitempty"`
	Categories Categories `xml:"categories" json:"categories,omitempty"`
}

type Categories struct {
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Category Category `xml:"category" json:"category,omitempty"`
}

type Category struct {
	Text   string `xml:",chardata" json:"text,omitempty"`
	Term   string `xml:"term,attr" json:"term,omitempty"`
	Scheme string `xml:"scheme,attr" json:"scheme,omitempty"`
}
