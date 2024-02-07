package mem_storage

import "skillbox28_5_hometask/pkg/student"

type StubStorage struct{}

func NewStubStore() *StubStorage {
	return &StubStorage{}
}

func (ss *StubStorage) Get(student *student.Student) bool {
	return true
}
func (ss *StubStorage) Put(student *student.Student) {
	return
}
func (ss *StubStorage) Map() map[string]*student.Student {
	s := make(map[string]*student.Student)
	s["Timur"] = &student.Student{
		Name:  "Timur",
		Age:   45,
		Grade: 5,
	}
	return s
}
