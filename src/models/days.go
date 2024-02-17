package models

import (
	"gingonic/src/config"

	"github.com/jinzhu/gorm"
)

type Day struct {
	gorm.Model
	Name string
}

func SelectAllDay() []Day {
	items := []Day{}
	config.DB.Raw("SELECT * FROM days").Scan(&items)
	return items
}

func SelectDayById(ID int) Day {
	item := Day{}
	config.DB.Raw("SELECT * FROM days WHERE id = ?", ID).Scan(&item)
	return item
}

func CreateDay(Name string) []Day {
	items := []Day{}
	config.DB.Raw("INSERT INTO days (created_at, updated_at, deleted_at, name) VALUES(NULL, NULL, NULL, ?)", Name).Scan(&items)
	return items
}

func UpdateDay(ID int, Name string) []Day {
	items := []Day{}
	config.DB.Raw("UPDATE days SET name = ? WHERE id = ?", Name, ID).Scan(&items)
	return items
}

func DeleteDay(ID int) []Day {
	items := []Day{}
	config.DB.Raw("DELETE FROM days WHERE id = ?", ID).Scan(&items)
	return items
}
