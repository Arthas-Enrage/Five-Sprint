package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
    Steps    int
    Duration time.Duration
    personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("неверный формат данных")
	}

	// Обработка шагов с учетом возможного знака '+'
	stepsInput := parts[0]
	stepsStr := strings.TrimPrefix(stepsInput, "+")
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}

	// Проверяем соответствие ввода
	expectedInput := strconv.Itoa(steps)
	if strings.HasPrefix(stepsInput, "+") {
		expectedInput = "+" + expectedInput
	}
	if expectedInput != stepsInput {
		return errors.New("некорректный формат шагов")
	}

	if steps <= 0 {
		return errors.New("шаги должны быть положительным числом")
	}
	ds.Steps = steps

	// Парсинг длительности
	durationStr := parts[1]
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("ошибка парсинга времени: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if ds.Steps <= 0 {
		return "", errors.New("количество шагов должно быть положительным")
	}
	if ds.Duration <= 0 {
		return "", errors.New("длительность должна быть положительной")
	}
	if ds.Weight <= 0 {
		return "", errors.New("вес должен быть положительным")
	}
	if ds.Height <= 0 {
		return "", errors.New("рост должен быть положительным")
	}

	// Расчет показателей
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
