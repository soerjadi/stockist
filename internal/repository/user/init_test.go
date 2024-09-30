package user

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

type prepareQueryMock struct {
	insertUser  *sqlmock.ExpectedPrepare
	getUserByID *sqlmock.ExpectedPrepare
}

func expectPrepareMock(mock sqlmock.Sqlmock) prepareQueryMock {
	prepareQuery := prepareQueryMock{}

	prepareQuery.insertUser = mock.ExpectPrepare(`
	INSERT INTO users \(
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt
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
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	`)

	prepareQuery.getUserByID = mock.ExpectPrepare(`
	SELECT
		id,
		name,
		email,
		phone_number,
		address,
		role,
		password,
		salt,
		created_at
	FROM
		users
	WHERE
		id = (.*)
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

				return &userRepository{
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
