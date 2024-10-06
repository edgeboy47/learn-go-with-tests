package reflection

import "reflect"

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, fn func(string)) {
	val := getValue(x)
	numValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
		break

	case reflect.Struct:
		numValues = val.NumField()
		getField = val.Field
		break

	case reflect.Slice, reflect.Array:
		numValues = val.Len()
		getField = val.Index
		break

	case reflect.String:
		fn(val.String())
		break

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}

	case reflect.Func:
		result := val.Call(nil)
		for _, val := range result {
			walk(val.Interface(), fn)
		}
	}

	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}
