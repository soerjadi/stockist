package user

import (
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/soerjadi/stockist/internal/repository/user"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository user.Repository
		validate   *validator.Validate
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &userUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository, tt.args.validate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
