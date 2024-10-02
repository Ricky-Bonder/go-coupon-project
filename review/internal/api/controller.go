package api

import (
	. "coupon_service/internal/api/entity"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (a *API) Apply(c *gin.Context) {
	apiReq := ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, basket)

	}
}

func (a *API) Create(c *gin.Context) {
	apiReq := Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	_, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusCreated)
	}
}

func (a *API) Get(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		log.Println("Error parsing code")
		c.Status(http.StatusBadRequest)
	} else {
		apiReq := CouponRequest{Code: code}
		if err := c.ShouldBindJSON(&apiReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		coupon, err := a.svc.GetSpecificCoupon(apiReq.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, coupon)
		}
	}
}

func (a *API) GetAll(c *gin.Context) {
	coupons, err := a.svc.GetCoupons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, coupons)
	}
}
