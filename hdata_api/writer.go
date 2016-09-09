package hdata_api

type Writer interface {
	Prepare(JobContext, PluginConfig)
	Execute(Record)
	Close()
}
