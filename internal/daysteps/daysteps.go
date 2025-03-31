package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	sliseStrings := strings.Split(datastring, ",")
	if len(sliseStrings) != 2 {
		return fmt.Errorf("ошибочный формат пакета")
	}
	var steps int
	steps, err = strconv.Atoi(sliseStrings[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(sliseStrings[1])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	ds.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("ошибка продолжительности прогулки")
	}
	dist := float64(ds.Steps) * StepLength / 1000
	ccal := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Personal.Weight), float64(ds.Personal.Height), ds.Duration) // не понял какую ошибку тут искать
	info := fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`, ds.Steps, dist, ccal)
	return info, nil
}
