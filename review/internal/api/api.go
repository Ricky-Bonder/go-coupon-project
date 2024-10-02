package api

import (
	"context"
	"coupon_service/internal/service/entity"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) (entity.Coupon, error)
	GetSpecificCoupon(string) (entity.Coupon, error)
	GetCoupons() ([]entity.Coupon, error)
	FindByCode(string) (entity.Coupon, error)
}

type Config struct {
	Host string
	Port int
}

type API struct {
	srv *http.Server
	MUX *gin.Engine
	svc Service
	CFG Config
}

func New[T Service](cfg Config, svc T) API {
	gin.SetMode(gin.ReleaseMode)
	r := new(gin.Engine)
	r = gin.New()
	r.Use(gin.Recovery())

	return API{
		MUX: r,
		CFG: cfg,
		svc: svc,
	}.withServer().withRoutes()
}

func (a API) withServer() API {

	ch := make(chan API)
	go func() {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", a.CFG.Port),
			Handler: a.MUX,
		}
		ch <- a
	}()

	return <-ch
}

func (a API) withRoutes() API {
	apiGroup := a.MUX.Group("/api")
	apiGroup.GET("/coupon/all", a.GetAll)
	apiGroup.POST("/coupon/:code", a.Get)
	apiGroup.POST("/create", a.Create)
	apiGroup.POST("/apply", a.Apply)
	return a
}

func (a API) Start() {
	go func() {
		err := a.srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (a API) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
