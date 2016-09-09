package hdata_api

const (
	PARALLELISM_KEY     string = "parallelism"
	DEFAULT_PARALLELISM int    = 1

	FLOWLIMIT_KEY     string = "flow.limit"
	DEFAULT_FLOWLIMIT int64  = 0
)

type PluginConfig struct {
	Configuration
}

func (p *PluginConfig) GetParallelism() int {
	parallelism, _ := p.Configuration.GetInt(PARALLELISM_KEY, DEFAULT_PARALLELISM)
	if parallelism < 1 {
		panic("Reader and Writer parallelism must be >= 1.")
	}
	return parallelism
}
func (p *PluginConfig) GetFlowLimit() int64 {
	flowLimit, _ := p.Configuration.GetLong(FLOWLIMIT_KEY, DEFAULT_FLOWLIMIT)
	if flowLimit < 0 {
		panic("Reader and Writer FLowLimit must be >= 0.")
	}
	return flowLimit
}
