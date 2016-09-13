package config

import (
	"encoding/xml"
	"io/ioutil"
	"strings"

	"github.com/zhangjunfang/hdata/config"
	"github.com/zhangjunfang/hdata/hdata_api"
	"github.com/zhangjunfang/hdata/hdata_core/common"
	"github.com/zhangjunfang/hdata/hdata_core/util"
)

type DefaultEngineConfig struct {
	*hdata_api.EngineConfig
}

func Create() *DefaultEngineConfig {

	path := util.GetConfigDir() + common.HDATA_XML
	return HdataXML(path)
}

func HdataXML(path string) *DefaultEngineConfig {

	path = strings.Replace(path, "\\", "/", -1)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	hdata := &config.Hdata{}
	xml.Unmarshal(content, hdata)
	configuration := &hdata_api.Configuration{}
	for _, v := range hdata.Configuration.Properties {
		configuration.SetString(v.Name, v.Value)
	}
	defaultEngineConfig := &DefaultEngineConfig{}
	engineConfig := &hdata_api.EngineConfig{}
	engineConfig.Configuration = configuration
	defaultEngineConfig.EngineConfig = engineConfig
	return defaultEngineConfig
}
