package entity

import "reflect"

type Param struct {
	inGroups      []string
	rawData       map[string]interface{}
	reflectObject reflect.Value
	reflectType   reflect.Type
}

func (p *Param) InGroups() []string {
	return p.inGroups
}

func (p *Param) SetInGroups(inGroups []string) *Param {
	p.inGroups = inGroups

	return p
}

func (p *Param) RawData() map[string]interface{} {
	return p.rawData
}

func (p *Param) SetRawData(rawData map[string]interface{}) *Param {
	p.rawData = rawData

	return p
}

func (p *Param) ReflectObject() reflect.Value {
	return p.reflectObject
}

func (p *Param) SetReflectObject(reflectObject reflect.Value) *Param {
	p.reflectObject = reflectObject

	return p
}

func (p *Param) ReflectType() reflect.Type {
	return p.reflectType
}

func (p *Param) SetReflectType(reflectType reflect.Type) *Param {
	p.reflectType = reflectType

	return p
}

func (p *Param) Init(
	inGroups []string,
	rawData map[string]interface{},
	reflectObject reflect.Value,
	reflectType reflect.Type,
) *Param {
	return p.
		SetInGroups(inGroups).
		SetRawData(rawData).
		SetReflectObject(reflectObject).
		SetReflectType(reflectType)
}
