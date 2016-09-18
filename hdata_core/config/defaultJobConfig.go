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

func (d *DefaultJobConfig) GetReaderConfig() *PluginConfig {

	return d.ReaderConfig

}

func (d *DefaultJobConfig) GetWriterConfig() *PluginConfig {
	return d.WriterConfig
}

func (d *DefaultJobConfig) GetReaderName() string {
	return d.ReaderName
}

func (d *DefaultJobConfig) GetWriterName() string {
	return d.WriterName
}

func (d *DefaultJobConfig) NewReader() Reader {

	return d
}

func (d *DefaultJobConfig) NewSplitter() Splitter {

}

func (d *DefaultJobConfig) NewWriter() Writer {

}
