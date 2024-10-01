package service

import (
	"coupon_service/internal/repository/memdb"
	. "coupon_service/internal/service/entity"
	"fmt"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Service struct {
	repo memdb.Repository[CouponID, Coupon]
}

func New(repo *gorm.DB) Service {
	return Service{
		repo: *memdb.MakeRepository[CouponID, Coupon](repo),
	}
}

func (s Service) ApplyCoupon(basket Basket, code string) (b *Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}
	return applyDiscountToBasket(b, coupon.Discount)
}

func applyDiscountToBasket(basket *Basket, discount int) (b *Basket, e error) {
	if basket.Value > 0 {
		basket.AppliedDiscount = discount
		basket.ApplicationSuccessful = true
		return basket, nil
	} else if basket.Value < 0 {
		return basket, fmt.Errorf("tried to apply discount to negative value")
	}
	return basket, nil
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) (Coupon, error) {
	coupon := Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		CouponID:       uuid.New(),
	}

	if _, err := s.repo.Save(&coupon); err != nil {
		return Coupon{}, err
	}
	return coupon, nil
}

func (s Service) GetCoupons(codes []string) ([]Coupon, error) {
	coupons := make([]Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
		}
		coupons = append(coupons, coupon)
	}

	return coupons, e
}

func (s Service) FindByCode(code string) (Coupon, error) {
	return s.repo.FindByCode(code)
}
