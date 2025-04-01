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
		// добавить ошибку парсинга
		if dp.Parse(v) != nil {
			fmt.Println(dp.Parse(v))
			continue
		}
	}
	fmt.Println(dp.ActionInfo())
}
