package product

import (
	"reflect"
	"testing"

	"github.com/soerjadi/stockist/internal/repository/product"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository product.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &productUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
