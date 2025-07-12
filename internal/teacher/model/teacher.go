package model

type Teacher struct {
	ID            int64
	Name          string
	Gender        string
	Subject       []string
	ClassSections []string
}
