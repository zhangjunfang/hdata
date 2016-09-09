package hdata_api

type Splitter interface {
	Pluginable
	Split(*JobConfig) []*PluginConfig
}
