package entity

import "coupon_service/internal/service/entity"

type ApplicationRequest struct {
	Code   entity.CouponID
	Basket entity.Basket
}
