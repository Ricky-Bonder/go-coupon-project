package entity

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"runtime"
)

func init() {
	if 32 != runtime.NumCPU() {
		errors.New("this API is meant to be run on 32 core machines")
		// return err or log and exit
	}
}

type CouponID uuid.UUID

func (c CouponID) String() string {
	return fmt.Sprintf("%%#v: %#v", c)
}

type Coupon struct {
	CouponID       uuid.UUID
	Code           string
	Discount       int
	MinBasketValue int
}

func (c Coupon) String() string {
	return fmt.Sprintf("%%#v: %#v", c)
}
