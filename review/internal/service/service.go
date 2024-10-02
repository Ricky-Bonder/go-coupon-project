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
	coupon, err := s.FindByCode(code)
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

func (s Service) GetSpecificCoupon(code string) (Coupon, error) {
	coupon, err := s.FindByCode(code)
	if err != nil {
		return Coupon{}, err
	}
	return coupon, nil
}

func (s Service) GetCoupons() ([]Coupon, error) {
	coupon, err := s.repo.GetAll()
	if err != nil {
		return []Coupon{}, err
	}
	return coupon, nil
}

func (s Service) FindByCode(code string) (Coupon, error) {
	coupons, err := s.repo.Get(Coupon{Code: code})
	if err != nil {
		return Coupon{}, err
	}
	if len(coupons) == 0 {
		return Coupon{}, fmt.Errorf("coupon not found")
	}
	return coupons[0], nil
}
