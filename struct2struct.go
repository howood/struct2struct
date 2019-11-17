package struct2struct

import (
	"fmt"
	"reflect"
)

var BOOLVAL = true

// Convert Struct to Struct With StructTag
func ConvertStructToStruct(fromData interface{}, toData interface{}, convertFromTag, convertToTag string) interface{} {
	fromElem := reflect.ValueOf(fromData).Elem()
	fromSize := fromElem.NumField()
	toElem := reflect.ValueOf(toData).Elem()

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
		if value != nil && toElemField.IsValid() {
			if toElemField.Type().Kind() == reflect.Ptr {
				if reflect.ValueOf(value).Type().Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
					continue
				} else if reflect.ValueOf(value).Type().Kind() != reflect.Ptr && reflect.ValueOf(value).IsValid() == false {
					continue
				}
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
					switch reflect.ValueOf(value).Type().Kind() {
					case reflect.Ptr:
						switch reflect.ValueOf(value).Type() {
						case reflect.ValueOf(&BOOLVAL).Type():
							toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
						default:
							if reflect.ValueOf(value).Type().Elem().Kind() == reflect.Struct &&
								reflect.ValueOf(value).Type().Elem().PkgPath() != toElemField.Type().Elem().PkgPath() {
								//Different package
								continue
							} else {
								toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
							}
						}
					default:
						toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
					}
				}
			} else {
				if reflect.ValueOf(value).Type().Kind() == reflect.Ptr && toElemField.Type().Kind() != reflect.Ptr {
					if reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
						continue
					}
					switch reflect.ValueOf(value).Type().Kind() {
					case reflect.Ptr:
						toElemField.Set(reflect.Indirect(reflect.ValueOf(value)).Convert(toElemField.Type()))
					default:
						toElemField.Set(reflect.Indirect(reflect.ValueOf(value)).Convert(toElemField.Type()))
					}
					toElemField.Set(reflect.Indirect(reflect.ValueOf(value)).Convert(toElemField.Type()))
				} else {
					toElemField.Set(reflect.ValueOf(value).Convert(toElemField.Type()))
				}
			}
		}
	}
	return toElem.Interface()
}

// Merge Struct to Struct With StructTag
func MergeStructToStruct(source interface{}, destination interface{}, convertFromTag, convertToTag string) interface{} {
	sourceelem := reflect.ValueOf(source).Elem()
	sourcesize := sourceelem.NumField()
	destinationelem := reflect.ValueOf(destination).Elem()

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
		if destinationelemfield.Type().Kind() == reflect.Ptr {
			if reflect.ValueOf(value).Type().Kind() == reflect.Ptr && reflect.Indirect(reflect.ValueOf(value)).IsValid() == false {
				continue
			} else if reflect.ValueOf(value).Type().Kind() != reflect.Ptr && reflect.ValueOf(value).IsValid() == false {
				continue
			}
			switch reflect.ValueOf(value).Type() {
			case reflect.ValueOf(&BOOLVAL).Type():
				destinationelemfield.Set(reflect.ValueOf(value).Convert(destinationelemfield.Type()))
			}
		} else {
			if fmt.Sprintf("%v", value) != "0" && fmt.Sprintf("%v", value) != "" {
				destinationelemfield.Set(reflect.ValueOf(value).Convert(destinationelemfield.Type()))
			}
		}
	}
	return destinationelem.Interface()
}
