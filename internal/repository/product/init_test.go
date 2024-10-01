package product

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type prepareQueryMock struct {
	getByID       *sqlmock.ExpectedPrepare
	getList       *sqlmock.ExpectedPrepare
	createProduct *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQuery := prepareQueryMock{}

	prepareQuery.getByID = mock.ExpectPrepare(`
	SELECT
		id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	FROM 
		products
	WHERE 
		id = (.*)
	`)

	prepareQuery.getList = mock.ExpectPrepare(`
	SELECT
		id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	FROM
		products
	ORDER BY id DESC
	LIMIT (.*)
	OFFSET (.*)
	`)

	prepareQuery.createProduct = mock.ExpectPrepare(`
	INSERT INTO products \(
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images
	\) VALUES \(
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		(.*),
		(.*)
	\) RETURNING
	 	id,
		name,
		description,
		weight,
		price,
		store_id,
		stock,
		images,
		created_at
	`)

	return prepareQuery
}

func TestGetRepository(t *testing.T) {

	tests := []struct {
		name     string
		initMock func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock)
		want     func(db *sqlx.DB) Repository
		wantErr  bool
	}{
		{
			name: "success",
			initMock: func() (*sqlx.DB, *sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				expectPrepareMock(mock)
				expectPrepareMock(mock)
				return sqlx.NewDb(db, "postgres"), db, mock
			},
			want: func(db *sqlx.DB) Repository {
				q, _ := prepareQueries(db)

				return &productRepository{
					query: q,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, dbMock, mock := tt.initMock()
			defer dbMock.Close()

			got, err := GetRepository(db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			want := tt.want(db)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetRepository() = %v, want %v", got, want)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err.Error())
			}
		})
	}
}
