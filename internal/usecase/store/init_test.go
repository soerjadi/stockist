package store

import (
	"reflect"
	"testing"

	"github.com/soerjadi/stockist/internal/repository/store"
)

func TestGetUsecase(t *testing.T) {
	type args struct {
		repository store.Repository
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "soerja",
			want: &storeUsecase{},
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
