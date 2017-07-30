package dbHandler

import (
	"span/go-swagger-example/examples/database/store"
	"span/go-swagger-example/examples/database/models"
	"time"
	"strconv"
	"github.com/go-openapi/strfmt"
)

// AnalyticsStore has event specific calls
type TodoStore struct {
	store.DataStore
}

func (store *TodoStore) GetTodo(id string) (bool, *models.TodoFull) {
	var todoFull models.TodoFull

	if err := store.Get(&todoFull, `
		SELECT name, created_at, updated_at, completed_at, id, completed
		FROM todo WHERE id = ?`, id); err != nil {
		return false, nil
	}

	return true, &todoFull
}

func (store *TodoStore) AddTodo(todo *models.TodoPartial) (bool, *models.TodoFull) {
	var completedDate *strfmt.DateTime

	date := strfmt.DateTime(time.Now())

	if todo.Completed {
		completedDate = &date
	}

	res, err := store.Exec(`
		INSERT INTO todo (name, completed, created_at, updated_at, completed_at)
		VALUES(?, ?, ?, ?, ?)
		`, todo.Name, todo.Completed, date, date, completedDate)

	if err != nil {
		return false, nil
	}
	if rows, err := res.RowsAffected(); rows == 0 || err != nil {
		return false, nil
	}

	id, _ := res.LastInsertId()

	if success, todoFull := store.GetTodo(strconv.Itoa(int(id))); success {
		return true, todoFull
	}

	return false, nil
}
