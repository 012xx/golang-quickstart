package main

import (
	_ "github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	gorm.Model
	Text   string
	States string
}

// DB初期化
func dbInit() {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInsert）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}

// CREATE(INSERT)
func dbInsert(text string, states string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbInsert）")
	}
	db.Create(&Todo{Text: text, States: states})
	defer db.Close()
}

// READ(SERECT)
// DB全取得
func dbGetAll() []Todo{
	db, err:= gorm.Open("sqlite3","test.sqlite3")
	if err != nil {
		panic("データベース開けず！(dbGetAll())")
	}
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DBひとつ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbGetOne）")
	}
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

}
