package hdata_api

type DefaultRecord struct {
	Fields     []interface{}
	Cursor     uint
	FieldCount uint
}

func (d *DefaultRecord) Add(v interface{}) {
	d.AddWithIndex(d.Cursor, v)
}
func (d *DefaultRecord) AddWithIndex(k uint, v interface{}) {
	if d.Fields == nil {
		d.Fields = make([]interface{}, d.FieldCount)
	}
	d.Fields[k] = v
	d.Cursor = d.Cursor + 1
}

func (d *DefaultRecord) Get(k uint) interface{} {
	return d.Fields[k]
}

func (d *DefaultRecord) Size() int {
	return len(d.Fields)
}
