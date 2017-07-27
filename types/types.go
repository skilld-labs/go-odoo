package types

import (
	"reflect"
	"time"
)

type Relations []Relation

type Relation struct {
	FieldName               string
	CreateNewFields         []map[string]interface{}
	CreateExistantFieldsIds []int
	UpdateFields            []UpdateField
	DeleteFieldsIds         []int
	RemoveFieldsIds         []int
	RemoveAll               bool
	ReplaceFieldsIds        []int
}

type UpdateField struct {
	Id     int
	Values map[string]interface{}
}

type NilableType interface {
	Type_() interface{}
}

type Type interface {
	NilableType_() interface{}
}

type Many2One struct {
	Id   int64
	Name string
}

func load(ns interface{}, s interface{}) interface{} {
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
		if se.Field(i).Type().Name() == "Time" {
			var t time.Time
			d := f.Elem().String()
			if len(d) == 10 {
				t, _ = time.Parse("2006-01-02", d)
			} else {
				t, _ = time.Parse("2006-01-02 15:04:05", d)
			}
			se.Field(i).Set(reflect.ValueOf(t))
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
				se.Field(i).Set(reflect.ValueOf(Many2One{Id: f.Elem().Index(0).Elem().Int(), Name: f.Elem().Index(1).Elem().String()}))
				continue
			}
		}
		se.Field(i).Set(f.Elem())
	}
	return s
}

func HandleRelations(fields *map[string]interface{}, args *Relations) {
	for _, arg := range *args {
		var am []interface{}
		if len(arg.CreateNewFields) > 0 {
			for _, field := range arg.CreateNewFields {
				am = append(am, AppendValues(0, nil, field))
			}
		}
		if len(arg.CreateExistantFieldsIds) > 0 {
			for _, id := range arg.CreateExistantFieldsIds {
				am = append(am, AppendValues(4, id, nil))
			}
		}
		if len(arg.UpdateFields) > 0 {
			for _, field := range arg.UpdateFields {
				am = append(am, AppendValues(1, field.Id, field.Values))
			}
		}
		if len(arg.DeleteFieldsIds) > 0 {
			for _, id := range arg.DeleteFieldsIds {
				am = append(am, AppendValues(2, id, nil))
			}
		}
		if len(arg.RemoveFieldsIds) > 0 {
			for _, id := range arg.RemoveFieldsIds {
				am = append(am, AppendValues(3, id, nil))
			}
		}
		if arg.RemoveAll {
			am = append(am, AppendValues(5, nil, nil))
		}
		if len(arg.ReplaceFieldsIds) > 0 {
			am = append(am, AppendValues(6, nil, arg.ReplaceFieldsIds))
		}
		(*fields)[arg.FieldName] = am
	}
}

func AppendValues(actionId int, Ids interface{}, values interface{}) []interface{} {
	var m []interface{}
	m = append(m, actionId)
	m = append(m, Ids)
	m = append(m, values)
	return m
}
