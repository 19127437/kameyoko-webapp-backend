package repository

import (
	"advanced-webapp-project/model"
	"context"
	"database/sql"
	"errors"
	"time"
)

type IPresRepo interface {
	FindPresentationById(presId string) (*model.Pres, error)
	FindAllPresentations() ([]*model.Pres, error)
	InsertPresentation(pres *model.Pres, userId string) (int64, error)
	UpdatePresentation(presId string, data model.Pres) (int64, error)
	DeletePresentation(presId string) (int64, error)
}

type presRepo struct {
	conn *sql.DB
}

func NewPresRepo(sqldb *sql.DB) *presRepo {
	return &presRepo{
		conn: sqldb,
	}
}

func (db *presRepo) FindPresentationById(presId string) (*model.Pres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	pres := model.Pres{}
	user := model.User{}
	err := db.conn.QueryRowContext(ctx, stmtSelectPresentationById, presId).Scan(
		&pres.Id,
		&pres.Name,
		&user.Id,
		&pres.ModifiedAt,
		&pres.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	pres.Owner = &user
	return &pres, nil
}

func (db *presRepo) FindAllPresentations() ([]*model.Pres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	rows, err := db.conn.QueryContext(ctx, stmtSelectAllPresentations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presList []*model.Pres
	for rows.Next() {
		var pres model.Pres
		var user model.User
		if err = rows.Scan(
			&pres.Id,
			&pres.Name,
			&user.Id,
			&pres.ModifiedAt,
			&pres.CreatedAt); err != nil {
			return nil, errors.New("error scanning")
		}

		pres.Owner = &user
		presList = append(presList, &pres)
	}

	return presList, nil
}

func (db *presRepo) InsertPresentation(pres *model.Pres, userId string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	pres.ModifiedAt = time.Now()
	pres.CreatedAt = time.Now()
	result, err := db.conn.ExecContext(ctx, stmtInsertPresentation,
		pres.Name,
		userId,
		pres.ModifiedAt,
		pres.CreatedAt,
	)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (db *presRepo) UpdatePresentation(presId string, data model.Pres) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	res, err := db.conn.ExecContext(ctx, stmtUpdatePresentation, data.Name, time.Now(), presId)
	if err != nil {
		return -1, nil
	}

	return res.RowsAffected()
}

func (db *presRepo) DeletePresentation(presId string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	res, err := db.conn.ExecContext(ctx, stmtDeletePresentation, presId)
	if err != nil {
		return -1, err
	}

	return res.RowsAffected()
}
