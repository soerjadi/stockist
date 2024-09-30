package user

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.insertUser, err = db.Preparex(insertUser)
	if err != nil {
		return q, err
	}

	q.getUserByID, err = db.Preparex(getUserByID)
	if err != nil {
		return q, err
	}

	q.getUserByPhoneNumber, err = db.Preparex(getUserByPhoneNumber)
	if err != nil {
		return q, err
	}

	q.getUserByEmail, err = db.Preparex(getUserByEmail)
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

	return &userRepository{
		query: query,
	}, nil
}
