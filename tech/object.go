package tech

import (
	"github.com/irishconstant/core/contract"
	"github.com/irishconstant/core/ref"
)

//Object Объект
type Object struct {
	Key          int
	Name         string
	BuildAddress string
	ref.Address
	SubObjects []SubObject
	Type       ObjectType
}

//SubObject Помещение в Объекте
type SubObject struct {
	Key    int
	Name   string
	Number string
}

//ObjectType Тип объекта
type ObjectType struct {
	Key      int
	Name     string
	Is354    bool
	IsEnergy bool
}

//SubObjectType Тип Подобъекта
type SubObjectType struct {
	Key    int
	Name   string
	IsFlat bool
}

//SupplyPoint отражает Точку поставки
type SupplyPoint struct {
	Key            int
	Number         int
	Resource       ref.EnergyResource
	RegisterPoints []contract.RegisterPoint
}
