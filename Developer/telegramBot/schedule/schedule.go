package schedule

import (
	"fmt"
	"time"
)

type Event struct {
	Name  string
	Room  string
	Start time.Time
	End   time.Time
}

var UpComming = []Event{
	{Name: "Архитектура предприятий", Room: "3-221/1", Start: time.Date(2022, time.April, 26, 8, 15, 0, 0, time.Local), End: time.Date(2022, time.April, 26, 9, 45, 0, 0, time.Local)},
	{Name: "Архитектура предприятий", Room: "3-221/1", Start: time.Date(2022, time.April, 26, 9, 55, 0, 0, time.Local), End: time.Date(2022, time.April, 26, 11, 25, 0, 0, time.Local)},
	{Name: "Менеджмент", Room: "3-321", Start: time.Date(2022, time.April, 26, 11, 50, 0, 0, time.Local), End: time.Date(2022, time.April, 26, 13, 20, 0, 0, time.Local)},
	{Name: "Менеджмент", Room: "3-321", Start: time.Date(2022, time.April, 26, 13, 45, 0, 0, time.Local), End: time.Date(2022, time.April, 26, 15, 15, 0, 0, time.Local)},
	{Name: "Правовые аспекты деятельности в области информационно-коммуникационных технологий", Room: "3-221", Start: time.Date(2022, time.April, 27, 15, 25, 0, 0, time.Local), End: time.Date(2022, time.April, 27, 16, 55, 0, 0, time.Local)},
	{Name: "Правовые аспекты деятельности в области информационно-коммуникационных технологий", Room: "3-221", Start: time.Date(2022, time.April, 27, 17, 05, 0, 0, time.Local), End: time.Date(2022, time.April, 27, 18, 35, 0, 0, time.Local)},
	{Name: "Архитектура предприятий", Room: "3-203", Start: time.Date(2022, time.April, 28, 8, 15, 0, 0, time.Local), End: time.Date(2022, time.April, 28, 9, 45, 0, 0, time.Local)},
	{Name: "Проектирование информационных систем", Room: "3-203", Start: time.Date(2022, time.April, 28, 9, 55, 0, 0, time.Local), End: time.Date(2022, time.April, 28, 11, 25, 0, 0, time.Local)},
	{Name: "Исследование операций и методы оптимизации", Room: "3-203", Start: time.Date(2022, time.April, 28, 11, 50, 0, 0, time.Local), End: time.Date(2022, time.April, 28, 13, 20, 0, 0, time.Local)},
	{Name: "Менеджмент", Room: "3-203", Start: time.Date(2022, time.April, 28, 13, 45, 0, 0, time.Local), End: time.Date(2022, time.April, 28, 15, 15, 0, 0, time.Local)},
	{Name: "Правовые аспекты деятельности в области информационно-коммуникационных технологий", Room: "3-212", Start: time.Date(2022, time.April, 29, 8, 15, 0, 0, time.Local), End: time.Date(2022, time.April, 29, 9, 45, 0, 0, time.Local)},
	{Name: "Информационные технологии и системы в сфере цифровой экономики", Room: "3-212", Start: time.Date(2022, time.April, 29, 9, 55, 0, 0, time.Local), End: time.Date(2022, time.April, 29, 11, 25, 0, 0, time.Local)},
	{Name: "Информационные технологии и системы в сфере цифровой экономики", Room: "3-407", Start: time.Date(2022, time.April, 30, 15, 25, 0, 0, time.Local), End: time.Date(2022, time.April, 30, 16, 55, 0, 0, time.Local)},
}

func GetSchedule(day string, count int, events []Event) string {
	var result string
	var preResult []Event
	for i := 0; i < len(events); i++ {
		if events[i].Start.Day() == time.Now().Day()+count {
			preResult = append(preResult, events[i])
		}
	}

	switch {
	case len(preResult) > 5:
		result += fmt.Sprintf("%s будет %d пар\n\n", day, len(preResult))
	case len(preResult) > 1 && len(preResult) < 5:
		result += fmt.Sprintf("%s будет %d пары\n\n", day, len(preResult))
	case len(preResult) == 1:
		result += fmt.Sprintf("%s будет %d парa\n\n", day, len(preResult))
	case len(preResult) == 0:
		return fmt.Sprintf("%s нет пар!!1!!1!!!1", day)
	}

	for i := 0; i < len(preResult); i++ {
		result += fmt.Sprintf("%s\nАудитория: %s\nНачало: %s\nКонец: %s\n\n", preResult[i].Name, preResult[i].Room, preResult[i].Start.Format("15:04:05.999"), preResult[i].End.Format("15:04:05.999"))
	}
	return result
}

func Current(events []Event) string {
	var result string
	var preResult []Event

	for i := 0; i < len(events); i++ {
		if events[i].Start.Day() == time.Now().Day() && events[i].Start.Before(time.Now()) && events[i].End.After(time.Now().Local()) {
			preResult = append(preResult, events[i])
		}
	}

	// for i := 0; i < len(events); i++ {
	// 	if events[i].Start.Day() == time.Now().Day() {
	// 		preResult = append(preResult, events[i])
	// 	}
	// }
	if len(preResult) != 0 {
		if preResult[0].End.Unix()-time.Now().Unix() <= 5400 {
			result = fmt.Sprintf("Сейчас идет %s\nЗакончится через: %v", preResult[0].Name, (preResult[0].End.Unix()-time.Now().Unix())/60)
		}
	} else {
		result = "Сейчас нет пары"
	}
	return result
}
