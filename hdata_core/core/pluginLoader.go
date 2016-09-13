package core

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/zhangjunfang/hdata/config"
	//	"github.com/zhangjunfang/hdata/hdata_core/common"
	//	"github.com/zhangjunfang/hdata/hdata_core/util"
)

var (
	ReaderMap map[string]string
	WriterMap map[string]string
)

func init() {
	ReaderMap = make(map[string]string, 0)
	WriterMap = make(map[string]string, 0)
	//path := util.GetConfigDir() + common.PLUGINS_XML
	//HdataPluginsXML(path)
}

func HdataPluginsXML(path string) {
	path = strings.Replace(path, "\\", "/", -1)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	r := make([]*config.Reader, 0)
	w := make([]*config.Writer, 0)
	r = append(r, &config.Reader{
		Name:  "rrname",
		Class: "ccClass",
	})

	w = append(w, &config.Writer{
		Name:  "wwname",
		Class: "wwClass",
	})
	mm := &config.Plugins{
		Readers: r,
		Writers: w,
	}
	data := &config.HdataPlugins{
		Plugins: mm,
	}
	b, _ := xml.Marshal(mm)
	fmt.Println(string(b))
	xml.Unmarshal(content, data)
	//fmt.Println(len(data.p.Readers))
	for _, v := range mm.Readers {
		ReaderMap[v.Name] = v.Class
		fmt.Println(v.Class)
	}
	for _, v := range mm.Writers {
		WriterMap[v.Name] = v.Class
		fmt.Println(v.Name)
	}

}
