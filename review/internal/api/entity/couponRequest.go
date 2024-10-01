package entity

import "coupon_service/internal/service/entity"

type CouponRequest struct {
	Codes []entity.CouponID `json:"codes"`
}
