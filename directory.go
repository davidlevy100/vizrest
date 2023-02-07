package main

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"feed" json:"feed,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Xmlns   string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Base    string   `xml:"base,attr" json:"base,omitempty"`
	ID      string   `xml:"id" json:"id,omitempty"`
	Title   string   `xml:"title" json:"title,omitempty"`
	Updated string   `xml:"updated" json:"updated,omitempty"`
	Link    Link     `xml:"link" json:"link,omitempty"`
	Entry   Entry    `xml:"entry" json:"entry,omitempty"`
}

type Link struct {
	Text string `xml:",chardata" json:"text,omitempty"`
	Href string `xml:"href,attr" json:"href,omitempty"`
	Rel  string `xml:"rel,attr" json:"rel,omitempty"`
	Type string `xml:"type,attr" json:"type,omitempty"`
}

type Author struct {
	Text string `xml:",chardata" json:"text,omitempty"`
	Name string `xml:"name" json:"name,omitempty"`
}

type Entry struct {
	Text     string   `xml:",chardata" json:"text,omitempty"`
	Xmlns    string   `xml:"xmlns,attr" json:"xmlns,omitempty"`
	Viz      string   `xml:"viz,attr" json:"viz,omitempty"`
	Category Category `xml:"category" json:"category,omitempty"`
	Title    string   `xml:"title" json:"title,omitempty"`
	Author   Author   `xml:"author" json:"author,omitempty"`
	Updated  string   `xml:"updated" json:"updated,omitempty"`
	ID       string   `xml:"id" json:"id,omitempty"`
	Summary  string   `xml:"summary" json:"summary,omitempty"`
	Link     []Link   `xml:"link" json:"link,omitempty"`
	Empty    string   `xml:"empty" json:"empty,omitempty"`
}
