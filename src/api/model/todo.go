package model

import (
	"database/sql"
	"strconv"
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func TodoAll(db *sql.DB) ([]*Todo, error) {
	rows, err := db.Query(`SELECT id, name, done FROM todo`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*Todo
	for rows.Next() {
		t := &Todo{}
		if err := rows.Scan(&t.ID, &t.Name, &t.Done); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func TodoByID(db *sql.DB, id string) (*Todo, error) {
	t := &Todo{}

	if err := db.QueryRow(`SELECT id, name, done FROM todo WHERE id=?`, id).Scan(&t.ID, &t.Name, &t.Done); err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Todo) Insert(db *sql.DB) (*Todo, error) {
	res, err := db.Exec(`INSERT INTO todo (name, done) VALUES (?, ?)`, t.Name, false)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Todo{
		ID:   int(id),
		Name: t.Name,
		Done: false,
	}, nil
}

func (t *Todo) Toggle(db *sql.DB) (*Todo, error) {
	_, err := db.Exec(`UPDATE todo SET done=? WHERE id = ?`, !t.Done, t.ID)
	if err != nil {
		return nil, err
	}

	id := strconv.Itoa(t.ID)
	todo, err := TodoByID(db, id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
