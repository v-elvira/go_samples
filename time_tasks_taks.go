package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, _ := parseDate(lines[0])
	tasks, _ := parseTasks(date, lines[1:])

	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	return time.Parse("02.01.2006", src)
	//return time.Time{}, errors.New("not implemented")
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	re, err := regexp.Compile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0, len(lines))
	day_tasks := make(map[string]int, len(lines))

	for _, line := range lines {
		var start_time, stop_time time.Time
		var err error
		parts := re.FindStringSubmatch(line)
		if len(parts) != 4 {
			return nil, errors.New("wrong task format")
		}
		if start_time, err = time.Parse("15:04", parts[1]); err != nil {
			return nil, err
		}
		if stop_time, err = time.Parse("15:04", parts[2]); err != nil {
			return nil, err
		}
		dur := stop_time.Sub(start_time)
		if dur <= 0 {
			return nil, errors.New("negative task duration")
		}

		title := parts[3]
		if idx, ok := day_tasks[title]; ok {
			tasks[idx].Dur += stop_time.Sub(start_time)
		} else {
			tasks = append(tasks, Task{Date: date, Dur: dur, Title: title})
			day_tasks[title] = len(tasks) - 1
		}
	}
	return tasks, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i].Dur > tasks[j].Dur })
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
