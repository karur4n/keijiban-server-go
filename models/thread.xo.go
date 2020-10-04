// Package models contains the types for schema 'go_keijiban'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Thread represents a row from 'go_keijiban.threads'.
type Thread struct {
	ID        string    `json:"id"`         // id
	Title     string    `json:"title"`      // title
	CreatedAt time.Time `json:"created_at"` // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Thread exists in the database.
func (t *Thread) Exists() bool {
	return t._exists
}

// Deleted provides information if the Thread has been deleted from the database.
func (t *Thread) Deleted() bool {
	return t._deleted
}

// Insert inserts the Thread to the database.
func (t *Thread) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO go_keijiban.threads (` +
		`id, title, created_at` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.ID, t.Title, t.CreatedAt)
	_, err = db.Exec(sqlstr, t.ID, t.Title, t.CreatedAt)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Thread in the database.
func (t *Thread) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE go_keijiban.threads SET ` +
		`title = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, t.Title, t.CreatedAt, t.ID)
	_, err = db.Exec(sqlstr, t.Title, t.CreatedAt, t.ID)
	return err
}

// Save saves the Thread to the database.
func (t *Thread) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Thread from the database.
func (t *Thread) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM go_keijiban.threads WHERE id = ?`

	// run query
	XOLog(sqlstr, t.ID)
	_, err = db.Exec(sqlstr, t.ID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// ThreadByID retrieves a row from 'go_keijiban.threads' as a Thread.
//
// Generated from index 'threads_id_pkey'.
func ThreadByID(db XODB, id string) (*Thread, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, created_at ` +
		`FROM go_keijiban.threads ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	t := Thread{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&t.ID, &t.Title, &t.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
