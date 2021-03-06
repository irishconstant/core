package doc

import "github.com/irishconstant/core/ref"

//Application Заявка
type Application struct {
}

//ApplicationType Тип заявки
type ApplicationType struct {
}

// DocType Тип документа. Справочник
type DocType struct {
	Key            int
	Name           string
	Citizenship    ref.Citizenship
	IsSerialNumber bool
	IsNumber       bool
	IsFromCode     bool
	IsDateBegin    bool
	IsDateEnd      bool
}

type Doc struct {
	Key  int
	Name string
}
