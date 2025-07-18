package repository

import (
	"sync"
	"time"

	"github.com/karanparmar2o/student-dashboard/api/remark"
)

type RemarkRepo interface {
	AddRemark(r *remark.Remark) (int64, error)
	GetRemarksForStudent(studentID int64) ([]*remark.Remark, error)
	ApproveRemark(remarkID int64) error
}

type memoryRemarkRepo struct {
	mu      sync.Mutex
	remarks map[int64]*remark.Remark
	nextID  int64
}

func NewMemoryRemarkRepo() RemarkRepo {
	return &memoryRemarkRepo{
		remarks: make(map[int64]*remark.Remark),
		nextID:  1,
	}
}

func (m *memoryRemarkRepo) AddRemark(r *remark.Remark) (int64, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	r.Id = m.nextID
	r.Date = time.Now().Format("2006-01-02")
	m.remarks[m.nextID] = r
	m.nextID++
	return r.Id, nil
}

func (m *memoryRemarkRepo) GetRemarksForStudent(studentID int64) ([]*remark.Remark, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	var list []*remark.Remark
	for _, r := range m.remarks {
		if r.StudentId == studentID {
			list = append(list, r)
		}
	}
	return list, nil
}

func (m *memoryRemarkRepo) ApproveRemark(remarkID int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if r, ok := m.remarks[remarkID]; ok {
		r.Approved = true
		return nil
	}
	return nil // or return error if not found
}
