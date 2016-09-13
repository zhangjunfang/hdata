package main

import (
	"encoding/xml"
	"fmt"

	"github.com/zhangjunfang/hdata/config"
	"github.com/zhangjunfang/hdata/hdata_core/core"
)

func main() {

	b, _ := xml.Marshal(config.Plugins{})
	fmt.Println(string(b))
	core.HdataPluginsXML("C:/workspace/golang/src/github.com/zhangjunfang/hdata/config/plugins.xml")
}
