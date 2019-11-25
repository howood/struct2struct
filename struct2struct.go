package struct2struct

import (
	"errors"
	"fmt"
	"reflect"
)

// ConvertStructToStruct Converts Struct to Struct With StructTag
func ConvertStructToStruct(fromData interface{}, toData interface{}, convertFromTag, convertToTag string) (interface{}, error) {
	fromElem, err := getReflectElem(fromData)
	if err != nil {
		return nil, err
	}
	fromSize := fromElem.NumField()
	toElem, err := getReflectElem(toData)
	if err != nil {
		return nil, err
	}

	for i := 0; i < fromSize; i++ {
		field := fromElem.Type().Field(i).Name
		// with convertFromTag ,replace using fromstruct field name to tag name
		if convertFromTag != "" {
			if tagval := fromElem.Type().Field(i).Tag.Get(convertFromTag); tagval != "" {
				field = tagval
			}
		}
		value := fromElem.Field(i).Interface()
		toElemField := toElem.FieldByName(field)
		// with convertToTag ,replace using tostruct field name to tag name
		if convertToTag != "" {
			for j := 0; j < toElem.NumField(); j++ {
				if tagval := toElem.Type().Field(j).Tag.Get(convertToTag); tagval == field {
					toElemField = toElem.Field(j)
				}
			}
		}
		if value != nil && toElemField.IsValid() && reflect.ValueOf(value).Type().Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
			continue
		}
		if value != nil && toElemField.IsValid() && reflect.ValueOf(value).Type().Kind() != reflect.Ptr && reflect.ValueOf(value).IsValid() == false {
			continue
		}
		if value != nil && toElemField.IsValid() {
			// To pointer
			if toElemField.Type().Kind() == reflect.Ptr {
				convertValueToPointer(value, &toElemField)
			} else {
				if reflect.ValueOf(value).Type().Kind() == reflect.Ptr && toElemField.Type().Kind() != reflect.Ptr {
					if reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
						continue
					}
					toElemField.Set(reflect.Indirect(reflect.ValueOf(value)).Convert(toElemField.Type()))
				} else {
					toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
				}
			}
		}
	}
	return toElem.Interface(), nil
}

// MergeStructToStruct Merges Struct to Struct With StructTag
// overwrite destination data with source data
func MergeStructToStruct(source interface{}, destination interface{}, convertFromTag, convertToTag string) (interface{}, error) {
	sourceelem, err := getReflectElem(source)
	if err != nil {
		return nil, err
	}
	sourcesize := sourceelem.NumField()
	destinationelem, err := getReflectElem(destination)
	if err != nil {
		return nil, err
	}

	for i := 0; i < sourcesize; i++ {
		field := sourceelem.Type().Field(i).Name
		if convertFromTag != "" {
			if tagval := sourceelem.Type().Field(i).Tag.Get(convertFromTag); tagval != "" {
				field = tagval
			}
		}
		value := sourceelem.Field(i).Interface()
		destinationelemfield := destinationelem.FieldByName(field)
		if convertToTag != "" {
			for j := 0; j < sourceelem.NumField(); j++ {
				if tagval := sourceelem.Type().Field(j).Tag.Get(convertToTag); tagval == field {
					destinationelemfield = destinationelem.Field(j)
				}
			}
		}
		if destinationelemfield.Type().Kind() == reflect.Ptr && reflect.ValueOf(value).Type().Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
			continue
		}
		if destinationelemfield.Type().Kind() == reflect.Ptr && reflect.ValueOf(value).Type().Kind() != reflect.Ptr && reflect.ValueOf(value).IsValid() == false {
			continue
		}
		if destinationelemfield.Type().Kind() != reflect.Ptr && (fmt.Sprintf("%v", value) == "0" || fmt.Sprintf("%v", value) == "") {
			continue
		}
		if destinationelemfield.Type().Kind() != reflect.Ptr && reflect.ValueOf(value).Type().Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
			continue
		}
		if destinationelemfield.Type().Kind() == reflect.Ptr {
			convertValueToPointer(value, &destinationelemfield)
		} else {
			if reflect.ValueOf(value).Type().Kind() == reflect.Ptr {
				destinationelemfield.Set(reflect.Indirect(reflect.ValueOf(value)).Convert(destinationelemfield.Type()))
			} else {
				destinationelemfield.Set(reflect.ValueOf(value).Convert(destinationelemfield.Type()))
			}
		}
	}
	return destinationelem.Interface(), nil
}

func getReflectElem(source interface{}) (reflect.Value, error) {
	switch reflect.ValueOf(source).Type().Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(source).Elem(), nil
	case reflect.Struct:
		r := reflect.Indirect(reflect.ValueOf(source)).Convert(reflect.ValueOf(source).Type())
		if r.CanSet() == false {
			return reflect.Value{}, errors.New("Not Valid Struct ot Struct pointer: CanSet")
		}
		if r.CanAddr() == false {
			return reflect.Value{}, errors.New("Not Valid Struct ot Struct pointer: CanAddr")
		}
		return r, nil
	default:
		return reflect.Value{}, errors.New("Not Struct ot Struct pointer")
	}
}

func convertValueToPointer(value interface{}, toElemField *reflect.Value) {
	switch converted := value.(type) {
	// simple type
	case int:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case int32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case int64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case uint:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case uint32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case uint64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case float32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case float64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case bool:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case string:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	// Map Type
	case map[string]string:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]int:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]int32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]int64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]uint:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]uint32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]uint64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]float32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]float64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]bool:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]string:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]int:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]int32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]int64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]uint:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]uint32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]uint64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]float32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]float64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]bool:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	// Map Type Pointer
	case map[string]*string:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*int:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*int32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*int64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*uint:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*uint32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*uint64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*float32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*float64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]*bool:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*string:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*int:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*int32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*int64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*uint:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*uint32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*uint64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*float32:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*float64:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]*bool:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string]interface{}:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	case map[string][]interface{}:
		toElemField.Set(reflect.ValueOf(&converted).Convert(toElemField.Type()))
	default:
		if reflect.ValueOf(value).Type().Kind() == reflect.Ptr &&
			reflect.ValueOf(value).Type().Elem().Kind() == reflect.Struct &&
			reflect.ValueOf(value).Type().Elem().PkgPath() != toElemField.Type().Elem().PkgPath() {
			return
		}
		toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
	}
}
