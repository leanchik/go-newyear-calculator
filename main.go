package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Введите вашу дату (в формате ДД.ММ.ГГ)")
	date1 := readDate()
	date2 := NewYearCalculator()

	days := calculateDaysBetween(date1, date2)
	fmt.Println("До Нового года осталось", days, "дней")
}

func readDate() time.Time {
	var input string
	fmt.Scan(&input)
	layout := "02.01.2006"
	parsed, err := time.Parse(layout, input)

	if err != nil {
		fmt.Println("Неверный формат даты. Попробуйте еще раз")
		return readDate()
	}

	newYear := NewYearCalculator()
	if parsed.After(newYear) {
		fmt.Println("Эй, ты уже в Новом году!")
		return readDate()
	}
	return parsed
}

func calculateDaysBetween(date1, date2 time.Time) int {
	if date1.After(date2) {
		date1, date2 = date2, date1
	}
	diff := date2.Sub(date1)
	return int(diff.Hours() / 24)
}

func NewYearCalculator() time.Time {
	now := time.Now()
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)
	return newYear
}
