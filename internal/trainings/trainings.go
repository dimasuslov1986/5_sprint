package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	sliseStrings := strings.Split(datastring, ",")
	if len(sliseStrings) != 3 {
		return fmt.Errorf("ошибочный формат пакета")
	}
	var steps int
	steps, err = strconv.Atoi(sliseStrings[0])
	if err != nil {
		return fmt.Errorf("ошибочный формат пакета")
	}
	t.Steps = steps

	switch {
	case sliseStrings[1] == "Ходьба":
		t.TrainingType = sliseStrings[1]
	case sliseStrings[1] == "Бег":
		t.TrainingType = sliseStrings[1]
	default:
		return fmt.Errorf("неизвестный тип тренировки")
	}

	duration, err := time.ParseDuration(sliseStrings[2])
	if err != nil {
		return fmt.Errorf("ошибочный формат времени")
	}
	t.Duration = duration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps)
	if dist <= 0 {
		return "", fmt.Errorf("ошибка")
	}
	meanSp := spentenergy.MeanSpeed(t.Steps, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		ccal := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		info := fmt.Sprintf(`Тип тренировки: %s
	Длительность: %.2f ч.
	Дистанция: %.2f км.
	Скорость: %.2f км/ч
	Сожгли калорий: %.2f`, t.TrainingType, t.Duration.Hours(), dist, meanSp, ccal)
		return info, nil
	case "Бег":
		ccal := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		info := fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, t.TrainingType, t.Duration.Hours(), dist, meanSp, ccal)
		return info, nil
	default:
		return "", fmt.Errorf("unknown training type")
	}
}
