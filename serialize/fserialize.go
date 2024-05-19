package serialize

import (
	"encoding/json"
	"fmt"
	"github.com/fdataflow/fcommon"
	"github.com/fdataflow/fiface"
	"reflect"
)

type DefaultSerialize struct {
}

var DefaultSerializeInstance = &DefaultSerialize{}

func IsSerialize(paramType reflect.Type) bool {
	return paramType.Implements(reflect.TypeOf((*fiface.ISerialize)(nil)).Elem())
}

func (s *DefaultSerialize) UnMarshal(arr fcommon.DataFlowRowArr, r reflect.Type) (reflect.Value, error) {
	if r.Kind() != reflect.Slice {
		return reflect.Value{}, fmt.Errorf("")
	}
	slice := reflect.MakeSlice(r, 0, len(arr))
	for _, row := range arr {
		var elem reflect.Value
		var err error
		elem, err = unMarshalStruct(row, r.Elem())
		if err == nil {
			slice = reflect.Append(slice, elem)
			continue
		}
		elem, err = unMarshalJSONString(row, r.Elem())
		if err == nil {
			slice = reflect.Append(slice, elem)
			continue
		}
		elem, err = unMarshalJSONStruct(row, r.Elem())
		if err == nil {
			slice = reflect.Append(slice, elem)
			continue
		} else {
			return reflect.Value{}, fmt.Errorf("")
		}
	}
	return slice, nil
}

func (s *DefaultSerialize) Marshal(data interface{}) (fcommon.DataFlowRowArr, error) {
	var arr fcommon.DataFlowRowArr
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice, reflect.Array:
		slice := reflect.ValueOf(data)
		for i := 0; i < slice.Len(); i++ {
			jsonData, err := json.Marshal(slice.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			arr = append(arr, string(jsonData))
		}
	default:
		jsonData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		arr = append(arr, string(jsonData))
	}
	return arr, nil
}

func unMarshalStruct(row fcommon.DataFlowRow, elemType reflect.Type) (reflect.Value, error) {
	rowType := reflect.TypeOf(row)
	if rowType == nil {
		return reflect.Value{}, fmt.Errorf("")
	}
	if rowType.Kind() != reflect.Struct && rowType.Kind() != reflect.Ptr {
		return reflect.Value{}, fmt.Errorf("")
	}
	if rowType.Kind() == reflect.Ptr {
		if reflect.ValueOf(row).IsNil() {
			return reflect.Value{}, fmt.Errorf("")
		}
		row = reflect.ValueOf(row).Elem().Interface()
		rowType = reflect.TypeOf(row)
	}
	if !rowType.AssignableTo(elemType) {
		return reflect.Value{}, fmt.Errorf("")
	}
	return reflect.ValueOf(row), nil
}

func unMarshalJSONString(row fcommon.DataFlowRow, elemType reflect.Type) (reflect.Value, error) {
	s, ok := row.(string)
	if !ok {
		return reflect.Value{}, fmt.Errorf("")
	}
	elem := reflect.New(elemType).Elem()
	err := json.Unmarshal([]byte(s), elem.Addr().Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return elem, nil
}

func unMarshalJSONStruct(row fcommon.DataFlowRow, elemType reflect.Type) (reflect.Value, error) {
	jsonData, err := json.Marshal(row)
	if err != nil {
		return reflect.Value{}, err
	}
	elem := reflect.New(elemType).Interface()
	err = json.Unmarshal(jsonData, elem)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(elem), nil
}
