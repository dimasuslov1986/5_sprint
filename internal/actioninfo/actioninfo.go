package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		dp.Parse(v) // добавить ошибку парсинга
	}
	fmt.Println(dp.ActionInfo())
}
