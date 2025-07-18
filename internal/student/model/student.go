package model

type Student struct {
	ID            int64
	Name          string
	Age           int32
	RollNumber    string
	Class         string
	Section       string
	ParentName    string
	ParentPhone   string
	AdmissionDate string
	Gender        string
	Address       string
	Marks         []SubjectMark
}

type SubjectMark struct {
	SubjectName      string
	Mark             int32
	GivenByTeacherID int64
}
