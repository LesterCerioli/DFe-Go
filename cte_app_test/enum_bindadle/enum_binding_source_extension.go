package enumbindable

import (
	"reflect"
)

type EnumBindingSourceExtension struct {
	enumType reflect.Type
}

func NewEnumBindingSourceExtension() *EnumBindingSourceExtension {
	return &EnumBindingSourceExtension{}
}

func NewEnumBindingSourceExtensionWithType(enumType reflect.Type) *EnumBindingSourceExtension {
	e := NewEnumBindingSourceExtension()
	e.SetEnumType(enumType)
	return e
}

func (e *EnumBindingSourceExtension) GetEnumType() reflect.Type {
	return e.enumType
}

func (e *EnumBindingSourceExtension) SetEnumType(value reflect.Type) {
	if value == e.enumType {
		return
	}

	if value != nil {
		enumType := value
		if value.Kind() == reflect.Ptr {
			enumType = value.Elem()
		}
		if enumType.Kind() != reflect.Int {
			panic("Type must be for an Enum.")
		}
	}

	e.enumType = value
}

func (e *EnumBindingSourceExtension) ProvideValue() interface{} {
	if e.enumType == nil {
		panic("The EnumType must be specified.")
	}

	actualEnumType := e.enumType
	if e.enumType.Kind() == reflect.Ptr {
		actualEnumType = e.enumType.Elem()
	}

	enumValues := reflect.New(reflect.SliceOf(actualEnumType)).Elem()
	for i := 0; i < actualEnumType.NumField(); i++ {
		field := actualEnumType.Field(i)
		if field.Type == actualEnumType {
			enumValues = reflect.Append(enumValues, reflect.ValueOf(field.Name))
		}
	}

	if actualEnumType == e.enumType {
		return enumValues.Interface()
	}

	tempSlice := reflect.MakeSlice(reflect.SliceOf(actualEnumType), enumValues.Len()+1, enumValues.Len()+1)
	reflect.Copy(tempSlice.Slice(1, tempSlice.Len()), enumValues)
	return tempSlice.Interface()
}
