package order

import (
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/soerjadi/stockist/internal/repository/order"
	"github.com/soerjadi/stockist/internal/repository/product"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository  order.Repository
		productRepo product.Repository
		redis       *redis.Client
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &orderUsecase{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsecase(tt.args.repository, tt.args.productRepo, tt.args.redis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
