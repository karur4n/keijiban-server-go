package models

import (
	sq "github.com/Masterminds/squirrel"
)

func FindAllThreads() ([]*Thread, error) {
	db := newDB()

	sql, _, err := sq.Select("id, title, created_at").From("threads").ToSql()
	if err != nil {
		return nil, err
	}

	var threads []*Thread

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		thread := &Thread{}
		if err = rows.Scan(&thread.ID, &thread.Title, &thread.CreatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	return threads, nil
}
