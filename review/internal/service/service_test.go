package service

import (
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service/entity"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

var db, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

func TestNew(t *testing.T) {
	type args struct {
		repo memdb.Repository[entity.CouponID, entity.Coupon]
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: memdb.Repository[entity.CouponID, entity.Coupon]{}},
			Service{repo: *memdb.MakeRepository[entity.CouponID, entity.Coupon](db)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	type fields struct {
		repo memdb.Repository[entity.CouponID, entity.Coupon]
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo memdb.Repository[entity.CouponID, entity.Coupon]
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
	}{
		{"Apply 10%", fields{*memdb.MakeRepository[entity.CouponID, entity.Coupon](db)}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			_, err := s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			assert.NoError(t, err)
		})
	}
}
