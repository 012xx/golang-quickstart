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
	var todos []Todo // 型のスライス（リスト）を宣言する
	db.Order("created_at desc").Find(&todos) // created_at カラム（作成日時）の降順で並べ替える & Todoテーブルから全てのレコードを取得
	db.Close() // DBを閉じる
	return todos // 取得したTodoを返す
}

// DBひとつ取得
func dbGetOne(id int) Todo {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbGetOne）")
	}
	var todo Todo // Todo型の変数を宣言
	db.First(&todo, id) // データベースのTodoテーブルから、指定された id を持つ最初のレコードを取得
	db.Close() // DBを閉じる
	return todo // 取得したTodoを返す
}

// UPDATE
func dbUpdate(id int, text string, states string) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbUpdate）")
	}
	var todo Todo
	db.First(&todo, id) // idを指定して、Todoを取得
	todo.Text = text // 取得したTodoのテキストを上書き
	todo.States = states // 取得したTodoのステータスを上書き
	db.Save(&todo) // 取得したTodoを保存
	db.Close() // DBを閉じる

// DELETE
func dbDelete(id int) {
	db, err := gorm.Open("sqlite3", "test.sqlite3")
	if err != nil {
		panic("データベース開けず！（dbDelete）"	)
	}
	var todo Todo
	db.First(&todo, id) // idを指定して、Todoを取得
	db.Delete(&todo) // 取得したTodoを削除
	db.Close() // DBを閉じる
	}

}
