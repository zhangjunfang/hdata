package hdata_api

type Fields struct {
	List []string
}

func (f *Fields) AddAll(fields []string) {
	if f.List == nil {
		f.List = make([]string, 0)
	}
	f.List = append(f.List, fields...)
}
