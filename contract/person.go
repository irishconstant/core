package contract

import (
	"time"

	"github.com/irishconstant/core/auth"
	"github.com/irishconstant/core/doc"
	"github.com/irishconstant/core/ref"
)

//Person Физическое лицо
type Person struct {
	Key int

	User           auth.User // Ответственный пользователь
	Name           string
	PatronymicName string
	FamilyName     string
	Citizenship    ref.Citizenship
	Sex            bool
	DateBirth      time.Time
	DateDeath      time.Time

	PossibleUsers []auth.User // Доступные пользователи для привязки
	Contacts      []Contact   // Контакты
	Docs          []doc.Doc   // Документы
}

// Contact Контакт
type Contact struct {
	Key     int
	Type    ContactType
	Value   string
	Address ref.Address
}

// ContactType Тип контакта. Справочник
type ContactType struct {
	Key        int
	Name       string
	Validation string
	IsAddress  bool
}
