package hdata_api

type Record interface {
	Add(v interface{})

	AddWithIndex(k uint, v interface{})

	Get(uint) interface{}

	Size() uint
}
