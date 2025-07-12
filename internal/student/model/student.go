package model

type Student struct {
	ID           int64 // auto-increment or UUID
	Name         string
	Age          int32
	RollNo       string
	Class        string
	Section      string
	ClassSection string // e.g., "7B"
}
