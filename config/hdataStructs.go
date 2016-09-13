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
