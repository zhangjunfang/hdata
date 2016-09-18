package util

/**
 * 数据类型转
 */
func ConvertInt64(v interface{}) (int64, bool) {

	val, ok := v.(int64)
	return val, ok

}
func ConvertFloat64(v interface{}) (float64, bool) {

	val, ok := v.(float64)
	return val, ok

}
func ConvertBool(v interface{}) (bool, bool) {
	val, ok := v.(bool)
	return val, ok

}
