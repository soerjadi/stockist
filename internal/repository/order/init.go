package order

import "github.com/jmoiron/sqlx"

func prepareQueries(db *sqlx.DB) (prepareQuery, error) {
	var (
		q   prepareQuery
		err error
	)

	q.createOrder, err = db.Preparex(createOrder)
	if err != nil {
		return q, err
	}

	q.createOrderItem, err = db.Preparex(createOrderItem)
	if err != nil {
		return q, err
	}

	q.updateOrderStatus, err = db.Preparex(updateOrderStatus)
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

	return &orderRepository{
		query: query,
	}, nil
}
