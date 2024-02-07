// Отрефакторить код прошлого домашнего задания. Декомпозируйте его так, чтобы логике одной
// сущности соответствовал один пакет. Для того, чтобы вы могли использовать методы и переменные,
// которые объявлены в других пакетах, сделайте их экспортируемыми.
package main

import (
	"skillbox28_5_hometask/pkg/app"
	mem_storage "skillbox28_5_hometask/pkg/storage"
)

func main() {
	//repository := stub_storage.NewStubStore()
	repository := mem_storage.NewMemStore()
	app := app.NewStorage(repository)
	app.Run()
}
