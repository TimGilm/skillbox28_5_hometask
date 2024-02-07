package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"skillbox28_5_hometask/pkg/student"
	"strconv"
	"strings"
)

type Storage interface {
	Get(student *student.Student) bool
	Put(student *student.Student)
	Map() map[string]*student.Student
}

type App struct {
	repository Storage
}

func NewStorage(storage Storage) *App {
	return &App{
		repository: storage,
	}
}

func (a *App) Run() {
	for {
		if student, ok := a.inputNextStudent(); ok {
			a.storeStudent(student)
		} else {
			break
		}
	}
}

func (a *App) inputNextStudent() (*student.Student, bool) {
	fmt.Println("Введите имя, возраст и оценку студента через пробел или `end` для выхода (либо ctrl+D в терминале)")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		err := io.EOF
		if len(parts) == 3 {
			name := parts[0]
			age, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Повторите ввод, введен неверный формат данных")
				continue
			}
			grade, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Повторите ввод, введен неверный формат данных")
				continue
			}
			return student.NewStudent(name, age, grade), true
		} else if (len(parts) == 1 && parts[0] == "end") || err != nil {
			fmt.Println("Завершение ввода данных, выход из программы")
			fmt.Println("Вывод на печать введенных данных")
			counter := 1
			for _, v := range a.repository.Map() {
				fmt.Printf("%v.student name: %v, age: %v, grade: %v\n", counter, v.Name, v.Age, v.Grade)
				counter += 1
			}
			break
		} else {
			fmt.Println("Повторите ввод, введены недостоверные данные")
			continue
		}
	}
	return nil, false
}

func (a *App) storeStudent(student *student.Student) {
	msg := "Студент с именем %v уже присутствует в хранилище\n"
	if a.repository.Get(student) {
		msg = "Данные студента с именем %v успешно добавлены в хранилище\n"
		a.repository.Put(student)
	}
	fmt.Printf(msg, student.Name)
}

func (a *App) printStudent(student *student.Student) {
	msg := "Студент с именем %v уже присутствует в хранилище\n"
	if a.repository.Get(student) {
		msg = "Данные студента с именем %v успешно добавлены в хранилище\n"
		a.repository.Put(student)
	}
	fmt.Printf(msg, student.Name)
}
