package mem_storage

import (
	"skillbox28_5_hometask/pkg/student"
)

type MemStorage struct {
	students map[string]*student.Student
}

func NewMemStore() *MemStorage {
	return &MemStorage{
		students: make(map[string]*student.Student),
	}
}

func (ms *MemStorage) Get(student *student.Student) bool {
	if ms.contains(student) {
		return false //т.е. дословно, если ms.contains(studentName) - true, то не добавляем и возвращаем false
	}
	return true
}

func (ms *MemStorage) Put(student *student.Student) {
	students := student.Name
	ms.students[students] = student
}

func (ms *MemStorage) Map() map[string]*student.Student {
	return ms.students
}

func (ms *MemStorage) contains(student *student.Student) bool {
	for st, _ := range ms.students {
		if student.Name == st {
			return true
		}
	}
	return false
}
