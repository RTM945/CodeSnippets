package cluster

// Action type for enum
type Action int

// Action values
const (
	ADD Action = iota
	DEL
)

// SDListener interface
type SDListener interface {
	AddServer(*Server)
	RemoveServer(*Server)
}
