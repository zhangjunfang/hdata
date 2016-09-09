package hdata_api

type Reader interface {
	Prepare(JobContext PluginConfig)

	Execute(RecordCollector)

	Close()

	DeclareOutputFields(OutputFieldsDeclarer)

	NewSplitter() Splitter
}
