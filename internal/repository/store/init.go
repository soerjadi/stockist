package store

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.insertStore, err = db.Preparex(insertStore)
	if err != nil {
		return q, err
	}

	q.getByID, err = db.Preparex(getByID)
	if err != nil {
		return q, err
	}

	return q, nil
}

func GetRepository(db *sqlx.DB) (Repository, error) {
	query, err := prepareQueries(db)
	if err != nil {
		return nil, err
	}

	return &storeRepository{
		query: query,
	}, nil
}
