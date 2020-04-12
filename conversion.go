package odoo

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

const (
	dateFormat     = "2006-01-02"
	datetimeFormat = "2006-01-02 15:04:05"
	tagName        = "xmlrpc"
)

func convertFromStaticToDynamic(static interface{}) map[string]interface{} {
	var dynamic = make(map[string]interface{})
	sv := reflect.ValueOf(static).Elem()
	st := reflect.TypeOf(static).Elem()
	for i := 0; i < sv.NumField(); i++ {
		field := sv.Field(i)
		if field.IsNil() {
			continue
		}
		key, _ := st.Field(i).Tag.Lookup(tagName)
		if dynamicValue := convertFromStaticToDynamicValue(field.Interface()); dynamicValue != nil {
			dynamic[strings.Split(key, ",")[0]] = dynamicValue
		}
	}
	return dynamic
}

func convertFromStaticToDynamicValue(staticValue interface{}) interface{} {
	var v interface{}
	switch staticValue.(type) {
	case *String:
		v = staticValue.(*String).v
	case *Int:
		v = staticValue.(*Int).v
	case *Bool:
		v = staticValue.(*Bool).v
	case *Selection:
		v = staticValue.(*Selection).v
	case *Time:
		v = staticValue.(*Time).v.Format(datetimeFormat)
	case *Float:
		v = staticValue.(*Float).v
	case *Many2One:
		v = staticValue.(*Many2One).ID
	case *Relation:
		v = staticValue.(*Relation).v
	default:
		v = staticValue
	}
	return v
}

func convertFromDynamicToStatic(dynamic interface{}, static interface{}) error {
	model := reflect.TypeOf(static).Elem()
	var sv reflect.Value
	switch dynamic.(type) {
	case []interface{}:
		if model.Kind() != reflect.Slice {
			return fmt.Errorf("cannot convert dynamic model to static model %s", model.Name())
		}
		sv = convertFromDynamicToStaticSlice(dynamic.([]interface{}), model)
	default:
		if model.Kind() == reflect.Slice {
			return fmt.Errorf("cannot convert dynamic model to static model %s", model.Name())
		}
		sv = convertFromDynamicToStaticOne(dynamic.(map[string]interface{}), model)
	}
	reflect.ValueOf(static).Elem().Set(sv)
	return nil
}

func convertFromDynamicToStaticSlice(dynamic []interface{}, sliceModel reflect.Type) reflect.Value {
	lenSlice := len(dynamic)
	ss := reflect.MakeSlice(sliceModel, lenSlice, lenSlice)
	for i := 0; i < lenSlice; i++ {
		ss.Index(i).Set(convertFromDynamicToStaticOne(dynamic[i].(map[string]interface{}), sliceModel.Elem()))
	}
	return ss
}

func convertFromDynamicToStaticOne(dynamic map[string]interface{}, oneModel reflect.Type) reflect.Value {
	s := reflect.New(oneModel).Elem()
	staticValues := scanStaticModelValues(oneModel, s)
	for key, dynamicValue := range dynamic {
		if _, ok := staticValues[key]; ok {
			staticField := staticValues[key]
			staticValue := convertFromDynamicToStaticValue(staticField.Type(), dynamicValue)
			if staticValue != nil {
				staticField.Set(reflect.ValueOf(staticValue))
			}
		}
	}
	return s
}

func convertFromDynamicToStaticValue(staticType reflect.Type, dynamicValue interface{}) interface{} {
	var staticValue interface{}
	if staticType.Kind() == reflect.Ptr {
		staticType = staticType.Elem()
	}
	typeName := staticType.Name()
	if !(dynamicValue == nil || (reflect.ValueOf(dynamicValue).Kind() == reflect.Bool && typeName != "Bool")) {
		switch typeName {
		case "String":
			staticValue = NewString(dynamicValue.(string))
		case "Int":
			staticValue = NewInt(dynamicValue.(int64))
		case "Selection":
			staticValue = NewSelection(dynamicValue)
		case "Float":
			staticValue = NewFloat(dynamicValue.(float64))
		case "Time":
			format := dateFormat
			if len(dynamicValue.(string)) > 10 {
				format = datetimeFormat
			}
			t, _ := time.Parse(format, dynamicValue.(string))
			staticValue = NewTime(t)
		case "Many2One":
			staticValue = NewMany2One(dynamicValue.([]interface{})[0].(int64), dynamicValue.([]interface{})[1].(string))
		case "Relation":
			staticValue = NewRelation()
			staticValue.(*Relation).ids = sliceInterfaceToInt64Slice(dynamicValue.([]interface{}))
		case "Bool":
			staticValue = NewBool(dynamicValue.(bool))
		default:
			staticValue = dynamicValue
		}
	}
	return staticValue
}

func scanStaticModelValues(typ reflect.Type, s reflect.Value) map[string]reflect.Value {
	fields := make(map[string]reflect.Value)
	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		key, _ := typ.Field(i).Tag.Lookup(tagName)
		fields[strings.Split(key, ",")[0]] = field
	}
	return fields
}

func sliceInterfaceToInt64Slice(s []interface{}) []int64 {
	i64 := make([]int64, len(s))
	for i := 0; i < len(s); i++ {
		i64[i] = s[i].(int64)
	}
	return i64
}
