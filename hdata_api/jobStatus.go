package hdata_api

type JobStatus int

const (
	SUCCESS = JobStatus(iota)
	FAILED
)
