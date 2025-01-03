package schema_member

import "fmt"

type SchemaMember interface {
	Validate(data interface{}) error
	IsRequired() bool
}

type IntDataType struct {
	name       string
	minVal     int32
	maxVal     int32
	isRequired bool
}

func (intDT *IntDataType) Validate(data interface{}) (err error) {
	intData, ok := data.(int32)
	if !ok {
		return fmt.Errorf("not an int32 type")
	}
	if intData < intDT.minVal || intData > intDT.maxVal {
		return fmt.Errorf("range invalid")
	}
	return nil
}
func (intDT *IntDataType) IsRequired() bool {
	return intDT.isRequired
}

func NewIntSchema(name string, minLen, maxLen int32, isRequired bool) SchemaMember {
	return &IntDataType{
		name: name,
		minVal: minLen,
		maxVal: maxLen,
		isRequired: isRequired,
	}
}

type StringDataType struct {
	name       string
	minLen     uint32
	maxLen     uint32
	isRequired bool
}

func (strDT *StringDataType) Validate(data interface{}) (err error) {
	strData, ok := data.(string)
	if !ok {
		return fmt.Errorf("not a string")
	}
	if len(strData) < int(strDT.minLen) || len(strData) > int(strDT.maxLen) {
		return fmt.Errorf("not valid range")
	}
	return
}
func (strDT *StringDataType) IsRequired() bool {
	return strDT.isRequired
}

func NewStrSchema(name string, minLen, maxLen uint32, isRequired bool) SchemaMember {
	return &StringDataType{
		name: name,
		minLen: minLen,
		maxLen: maxLen,
		isRequired: isRequired,
	}
}