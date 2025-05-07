package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
    Steps         int
    TrainingType  string
    Duration      time.Duration
    personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных")
	}

	// Парсинг и валидация шагов
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil || steps <= 0 {
		return errors.New("некорректное количество шагов")
	}
	t.Steps = steps

	// Парсинг типа тренировки
	t.TrainingType = strings.TrimSpace(parts[1])

	// Парсинг и валидация длительности
	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil || duration <= 0 {
		return errors.New("некорректная длительность")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return "", errors.New("неизвестный тип тренировки")
	}

	// Расчет показателей
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	// Форматирование с завершающим \n
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil
}
