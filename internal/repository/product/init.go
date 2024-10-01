package product

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.getByID, err = db.Preparex(getByID)
	if err != nil {
		return q, err
	}

	q.getList, err = db.Preparex(getList)
	if err != nil {
		return q, err
	}

	q.createProduct, err = db.Preparex(createProduct)
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

	return &productRepository{
		query: query,
	}, nil
}
