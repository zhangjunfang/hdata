package core

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"github.com/zhangjunfang/hdata/config"
	"github.com/zhangjunfang/hdata/hdata_core/common"
	"github.com/zhangjunfang/hdata/hdata_core/util"
)

var (
	ReaderMap map[string]string
	WriterMap map[string]string
)

func init() {
	ReaderMap = make(map[string]string, 0)
	WriterMap = make(map[string]string, 0)
	path := util.GetConfigDir() + common.PLUGINS_XML
	HdataPluginsXML(path)
}

func HdataPluginsXML(path string) {
	path = strings.Replace(path, "\\", "/", -1)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	data := &config.Plugins{}
	xml.Unmarshal(content, data)
	for _, v := range data.R.ReaderList {
		ReaderMap[v.Name] = v.Class
		//fmt.Println(v.Class)
	}
	for _, v := range data.W.WriterList {
		WriterMap[v.Name] = v.Class
		//fmt.Println(v.Name)
	}

}
