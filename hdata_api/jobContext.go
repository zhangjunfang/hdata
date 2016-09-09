package hdata_api

type JobContext struct {
	EngineConfig
	JobConfig
	*OutputFieldsDeclarer
	Storage
	*Metric
	Readers          []*Reader
	Writers          []*Writer
	IsReaderFinished bool
	IsReaderError    bool
	IsWriterFinished bool
	IsWriterError    bool
	JobStatus
}
