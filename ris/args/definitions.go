package args

import (
	"encoding/xml"
)

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Value   string   `xml:"value,attr"`
}

type ShaderType struct {
	XMLName xml.Name `xml:"shaderType"`
	Tag     Tag      `xml:"tag"`
}

type Tags struct {
	XMLName xml.Name `xml:"tags"`
	Tags    []Tag    `xml:"tag"`
}

type Help struct {
	XMLName xml.Name `xml:"help"`
	Value   string   `xml:",innerxml"`
}

type String struct {
	XMLName xml.Name `xml:"string"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type HintDict struct {
	XMLName    xml.Name `xml:"hintdict"`
	Attributes []String `xml:"string"`
	Name       string   `xml:"name,attr"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Tags    Tags     `xml:"tags"`
	Help    Help     `xml:"help"`

	Label       string `xml:"label,attr"`
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"`
	Default     string `xml:"default,attr"`
	Min         string `xml:"min,attr"`
	Max         string `xml:"max,attr"`
	Widget      string `xml:"widget,attr"`
	Connectable string `xml:"connectable,attr"`
}

type Output struct {
	XMLName xml.Name `xml:"output"`
	Name    string   `xml:"name,attr"`
	Tags    Tags     `xml:"tags"`
}

type Rfmdata struct {
	XMLName xml.Name `xml:"rfmdata"`

	NodeId         string `xml:"nodeid,attr"`
	Classification string `xml:"classification,attr"`
}

type Page struct {
	XMLName xml.Name `xml:"page"`
	Name    string   `xml:"name,attr"`
	Open    string   `xml:"open,attr"`
	Params  []Param  `xml:"param"`
}

type Args struct {
	XMLName xml.Name   `xml:"args"`
	Shader  ShaderType `xml:"shaderType"`
	Pages   []Page     `xml:"page"`
	Params  []Param    `xml:"param"`
	Outputs []Output   `xml:"output"`
	Rfmdata Rfmdata    `xml:"rfmdata"`
	Help    Help       `xml:"help"`
	Format  string     `xml:"format,attr"`
}
