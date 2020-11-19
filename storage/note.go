package storage

import (
	"context"
	"notes/model"
	"notes/storage/util"
	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	queryNotes    = `SELECT id, content, start_time, update_time FROM note`
	queryNoteByID = `SELECT id, content, start_time, update_time FROM note where id = ?`
	addNote       = `INSERT INTO note (content) VALUES (?)`
	removeNote    = `DELETE FROM note WHERE id = ?`
)

func GetNoteByID(ctx context.Context, id string) (model.NewNote, error) {
	var (
		err   error
		res   model.NewNote
		param []interface{}
	)

	param = append(param, id)

	row, err := util.DBConnector.Query(queryNoteByID, param...)
	if err != nil {
		return res, err
	}
	iter := sqlbuilder.NewIterator(row)
	err = iter.One(&res)

	return res, err
}

func RemoveNote(ctx context.Context, id string) error {

	var param []interface{}

	param = append(param, id)

	_, err := util.DBConnector.Exec(removeNote, param...)

	return err
}

func AddNote(ctx context.Context, content string) error {

	var param []interface{}

	param = append(param, content)

	_, err := util.DBConnector.Exec(addNote, param...)

	return err
}

func GetAllNotes(ctx context.Context) ([]model.NewNote, error) {

	var (
		err   error
		res   []model.NewNote
		param []interface{}
	)

	rows, err := util.DBConnector.Query(queryNotes, param...)
	if err != nil {
		return res, err
	}

	iter := sqlbuilder.NewIterator(rows)
	err = iter.All(&res)

	return res, err
}