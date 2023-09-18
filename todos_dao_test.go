package main

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func TestGetAll(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test' ,'test', false, now(), now())")

	dao := NewTodoDao(conn)
	todos, err := dao.GetAll()
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	if len(todos) == 0 {
		cleanup(conn)
		t.Fatal("Expected at least one todo")
	}
	conn.Exec("delete from todos where id = 'test'")

}

func TestGet(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test' ,'test', false, now(), now())")

	dao := NewTodoDao(conn)
	todos, err := dao.Get("test")
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	if todos == nil {
		cleanup(conn)
		t.Fatal("Expected at least one todo")
	}
	if todos.Title != "test" {
		cleanup(conn)
		t.Fatal("Expected title to be test")
	}

	conn.Exec("delete from todos where title = 'test'")

}

func TestCreate(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	dao := NewTodoDao(conn)
	todo := &Todo{
		ID:        "test",
		Title:     "test",
		Completed: false,
	}

	err = dao.Create(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}
	cleanup(conn)
}

func TestUpdate(t *testing.T) {
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
	todo := &Todo{
		ID:        "test",
		Title:     "test1",
		Completed: true,
	}
	err = dao.Update(todo)
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}
	cleanup(conn)
}

func TestDelete(t *testing.T) {
	t.Logf("Testing Delete")
	conn, err := sql.Open("pgx", "postgres://todos:T0d05!@localhost:5432/todos?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("insert into todos (id, title, completed, created_at, updated_at) values ('test', 'test', false, now(), now())")

	dao := NewTodoDao(conn)
	err = dao.Delete("test")
	if err != nil {
		cleanup(conn)
		t.Fatal(err)
	}

	cleanup(conn)
}

func cleanup(conn *sql.DB) {
	conn.Exec("delete from todos where title = 'test'")
}
