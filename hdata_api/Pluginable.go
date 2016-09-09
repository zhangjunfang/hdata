package hdata_api

type Pluginable interface {
	GetPluginName() string
	SetPluginName(string)
}
