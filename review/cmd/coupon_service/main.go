package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"flag"
	"go.uber.org/zap"
	"time"
)

var (
	cfg    = config.New()
	dbFile = flag.String("db", "test.db?_pragma=busy_timeout(5000)", "db file")
	logger *zap.SugaredLogger
)

const (
	HOURS_IN_YEAR = 24 * 365
)

func main() {
	setupLogger()
	dbConn := memdb.InitDb(*dbFile)
	//ctx, cancel := context.WithCancel(
	//	context.WithValue(
	//		context.Background(),
	//		"log", logger,
	//	),
	//)

	svc := service.New(dbConn)
	couponService := api.New(cfg.API, svc)
	couponService.Start()
	logger.Infof("Starting Coupon service server")

	coupon, err := svc.CreateCoupon(10, "10OFF", 100)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("Created coupon: %v", coupon)

	foundCoupon, err := svc.FindByCode(coupon.Code)
	if err != nil {
		return
	}

	logger.Infof("Found coupon: %v", foundCoupon)

	<-time.After(HOURS_IN_YEAR * time.Hour)
	logger.Infof("Coupon service server alive for a year, closing")
	couponService.Close()
}

func setupLogger() {
	l, _ := zap.NewDevelopment()
	logger = l.Sugar()
	zap.RedirectStdLog(logger.Desugar())
	zap.ReplaceGlobals(logger.Desugar())
	defer logger.Sync()
}
