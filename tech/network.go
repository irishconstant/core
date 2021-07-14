package tech

import "time"

//Node элемент схемы энергоснабжения
type Node struct {
	Key      int
	Name     string
	Type     NodeType
	Location string // В случаях, когда расположение нельзя указать через Объект
	Object   Object
	IsFirst  bool
}

//NodeType Тип элемента схемы энергоснабжения
type NodeType struct {
	Key       int
	Name      string
	ShortName string

	IsSource bool
	IsEnergy bool

	AfterTypes []NodeType // Типы, которые могут идти за данным типом
}

//Edge линия энергоснабжения
type Edge struct {
	Key int

	ParentNode Node
	ChildNode  Node

	EdgeType

	Pipelines []Pipeline
	Events    []Event
}

// EdgeType (назначение линии): магистральная, внутриквартальная, распределительная
type EdgeType struct {
	Key  int
	Name string
}

//Event переключение. Включения/отключения/ограничения линии
type Event struct {
	Key   int
	Date  time.Time
	Value int // Процент активности линии
}
