package hdata_api

import (
	"fmt"
)

//var
type Configuration struct{}

var (
	Map map[string]interface{}
)

func init() {

	Map = make(map[string]interface{})
}

type Configuration_method interface {
	GetStringWithDefaultValue(string, string) (string, error) //

	GetString(string) (string, error) //

	SetString(string, string) //

	GetInt(string, int) (int, error) //

	SetInt(string, int) error //

	GetLong(string, int64) (int64, error) //

	SetLong(string, int64) //

	GetDouble(string, float64) (float64, error) //

	SetDouble(string, float64) //

	GetBoolean(string, bool) (bool, error) //

	SetBoolean(string, bool) //

	GetFloat(string, float32) (float32, error) //

	SetFloat(string, float32) //
}

func (c Configuration) GetStringWithDefaultValue(key string, defaultValue string) (string, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(string); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}
}

func (c Configuration) GetString(key string) (string, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(string); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	}
	return "", nil
}

func (c Configuration) SetString(key string, value string) error {

	Map[key] = value

	return nil
}
func (c Configuration) GetInt(key string, defaultValue int) (int, error) {

	if v, ok := Map[key]; ok {
		if value, b := v.(int); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}

	return 0, nil
}
func (c Configuration) SetInt(key string, value int) error {
	Map[key] = value
	return nil
}
func (c Configuration) GetLong(key string, defaultValue int64) (int64, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(int64); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}
}
func (c Configuration) SetLong(key string, value int64) {
	Map[key] = value
}

func (c Configuration) GetDouble(key string, defaultValue float64) (float64, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(float64); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}
}
func (c Configuration) SetDouble(key string, value float64) {
	Map[key] = value
}
func (c Configuration) GetBoolean(key string, defaultValue bool) (bool, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(bool); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}
}
func (c Configuration) SetBoolean(key string, value bool) {
	Map[key] = value
}
func (c Configuration) GetFloat(key string, defaultValue float32) (float32, error) {
	if v, ok := Map[key]; ok {
		if value, b := v.(float32); b {
			return value, nil
		} else {
			return value, fmt.Errorf("%s", " Type conversion error ")
		}
	} else {
		return defaultValue, nil
	}
}
func (c Configuration) SetFloat(key string, v float32) {
	Map[key] = v
}
