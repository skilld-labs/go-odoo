package types

import (
	"reflect"
)
func convertNil(s interface{}, ns interface{}) interface{} {
	nse := reflect.ValueOf(ns).Elem()
	se := reflect.ValueOf(s).Elem()
	for i := 0; i < nse.NumField(); i++ {
		f := nse.Field(i)
		if f.Kind() == reflect.Bool {
			se.Field(i).SetBool(f.Bool())
			continue
		}
		if f.IsNil() || f.Elem().Kind() == reflect.Bool {
			continue
		}
		if f.Elem().Kind() == reflect.Slice {
			if se.Field(i).Kind() == reflect.Slice {
				interfaceSlice := f.Interface().([]interface{})
				int64Slice := make([]int64, len(interfaceSlice))
				for j := range interfaceSlice {
					int64Slice[j] = interfaceSlice[j].(int64)
				}
				se.Field(i).Set(reflect.ValueOf(int64Slice))
				continue
			}
			if se.Field(i).Kind() == reflect.Struct {
				var many2One struct {
					Id   int64
					Name string
				}
				many2One.Id = f.Elem().Index(0).Elem().Int()
				many2One.Name = f.Elem().Index(1).Elem().String()
				se.Field(i).Set(reflect.ValueOf(many2One))
				continue
			}
		}
		se.Field(i).Set(f.Elem())
	}
	return s
}
