package hdata_api

type JobConfig struct {
	*Configuration
	JobConfiger
}

type JobConfiger interface {
	GetReaderConfig() *PluginConfig

	GetWriterConfig() *PluginConfig

	GetReaderName() string

	GetWriterName() string

	NewReader() Reader

	NewSplitter() Splitter

	NewWriter() Writer
}
