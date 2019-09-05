package utils

import "reflect"

//SliceIndexOf ..
func SliceIndexOf(val interface{}, slice interface{}) int {
	v := reflect.ValueOf(val)
	arr := reflect.ValueOf(slice)

	t := reflect.TypeOf(slice).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}
