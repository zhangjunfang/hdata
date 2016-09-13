package config

import (
	"encoding/xml"
)

type Hdata struct {
	Configuration `xml:"configuration"`
}
type Configuration struct {
	Properties []*Property `xml:"property"`
}
type Property struct {
	Name        string `xml:"name"`
	Value       string `xml:"value"`
	Description string `xml:"description"`
}

//------------------------------------
type Reader struct {
	Name  string `xml:"name"`
	Class string `xml:"class"`
}
type Writer struct {
	Name  string `xml:"name"`
	Class string `xml:"class"`
}
type Readers struct {
	ReaderList []*Reader `xml:"reader"`
}
type Writers struct {
	WriterList []*Writer `xml:"writer"`
}
type Plugins struct {
	XMLName xml.Name `xml:"plugins"`
	R       *Readers `xml:"readers"`
	W       *Writers `xml:"writers"`
}
