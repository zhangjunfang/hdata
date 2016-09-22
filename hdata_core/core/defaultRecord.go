package core

import (
	"github.com/zhangjunfang/hdata/hdata_api"
)

var fields []interface{} = make([]interface{}, 0)
var cursor int

type DefaultRecord struct {
}

func (d *DefaultRecord) Add(v interface{}) {
	d.AddWithIndex(cursor, v)
}

func (d *DefaultRecord) AddWithIndex(k uint, v interface{}) {
	fields[k] = v
	cursor++
}

func (d *DefaultRecord) Get(k uint) interface{} {
	return fields[k]
}

func (d *DefaultRecord) Size() uint {
	return len(fields)
}
