package db

import (
	"context"
	"database/sql"
)

type DBTX interface {
	PrepareContext(context.Context, string) (*sql.Stmt, error)
}

func Query(ctx context.Context, dbtx DBTX, query string, args []interface{}) (data *sql.Rows, err error) {
	stmt, err := dbtx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	data, err = stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func QueryRow(ctx context.Context, dbtx DBTX, query string, args []interface{}) (row *sql.Row, err error) {
	stmt, err := dbtx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row = stmt.QueryRowContext(ctx, args...)

	return row, nil
}

func Exec(ctx context.Context, dbtx DBTX, query string, args []interface{}) (insertID, rowsAffected int64, err error) {
	stmt, err := dbtx.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	insertID, err = result.LastInsertId()
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	return
}
