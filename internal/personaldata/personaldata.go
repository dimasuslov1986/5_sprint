package personaldata

import (
	"fmt"
)

// Ниже создайте структуру Personal
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Ниже создайте метод Print()
func (pers Personal) Print() {
	fmt.Println("Имя:", pers.Name)
	fmt.Println("Вес:", pers.Weight)
	fmt.Println("Рост:", pers.Height)
}
