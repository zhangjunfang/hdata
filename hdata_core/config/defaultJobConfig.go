package config

import (
	"github.com/zhangjunfang/hdata/hdata_api"
)

type DefaultJobConfig struct {
	*hdata_api.JobConfig
	ReaderConfig *hdata_api.PluginConfig
	WriterConfig *hdata_api.PluginConfig
	ReaderName   string
	WriterName   string
}
