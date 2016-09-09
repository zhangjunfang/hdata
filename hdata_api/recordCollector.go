package hdata_api

type RecordCollector interface {
	Send(*Record)

	Sends([]*Record)
}
