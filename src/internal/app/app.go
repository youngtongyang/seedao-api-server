package app

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/handler"
	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/aliyunsms"
	"github.com/seedao/seedao-api-server/internal/pkg/jwtutil"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
	"github.com/seedao/seedao-api-server/internal/pkg/wechatpay"
	"github.com/seedao/seedao-api-server/internal/service"
)

func NewEcho(cfg *config.Config, db *gorm.DB, redis *redisstore.Store) (*echo.Echo, error) {
	smsSender, err := aliyunsms.NewClient(cfg.AliyunSMS)
	if err != nil {
		return nil, err
	}

	jwtMgr := jwtutil.NewManager(cfg.JWT.Secret, cfg.JWT.ExpHours)

	wxClient, err := wechatpay.NewClient(wechatpay.Config{
		DevMode:        cfg.WechatPay.DevMode,
		MchID:          cfg.WechatPay.MchID,
		AppID:          cfg.WechatPay.AppID,
		APIv3Key:       cfg.WechatPay.APIv3Key,
		CertSerialNo:        cfg.WechatPay.CertSerialNo,
		PrivateKeyPath:      cfg.WechatPay.PrivateKeyPath,
		WechatPublicKeyID:   cfg.WechatPay.WechatPublicKeyID,
		WechatPublicKeyPath: cfg.WechatPay.WechatPublicKeyPath,
		NotifyURL:           cfg.WechatPay.NotifyURL,
		H5WapName:      cfg.WechatPay.H5WapName,
		H5WapURL:       cfg.WechatPay.H5WapURL,
		H5ReturnURL:    cfg.WechatPay.H5ReturnURL,
	})
	if err != nil {
		return nil, err
	}

	smsSvc := service.NewSMSService(db, redis, smsSender)
	snsSvc := service.NewSNSService(db, redis)
	walletSvc := service.NewWalletService(redis, cfg.SIWE)
	authSvc := service.NewAuthService(db, smsSvc, snsSvc, walletSvc, jwtMgr)
	profileSvc := service.NewProfileService(db)
	statsSvc := service.NewStatsService(db)
	membershipSvc := service.NewMembershipService(db)
	paymentSvc := service.NewPaymentService(db, redis, wxClient, cfg.WechatPay)

	authHandler := handler.NewAuthHandler(authSvc, smsSvc)
	snsHandler := handler.NewSNSHandler(snsSvc)
	walletHandler := handler.NewWalletHandler(walletSvc, snsSvc)
	profileHandler := handler.NewProfileHandler(profileSvc)
	statsHandler := handler.NewStatsHandler(statsSvc)
	membershipHandler := handler.NewMembershipHandler(membershipSvc)
	paymentHandler := handler.NewPaymentHandler(paymentSvc, cfg.WechatPay.DevMode)

	e := echo.New()
	e.HideBanner = true
	e.Use(echomw.Recover())
	e.Use(echomw.RequestID())
	e.Use(echomw.Logger())
	e.Use(middleware.DebugMode(cfg.Server.Debug))

	e.GET("/health", handler.Health)

	v1 := e.Group("/v1")
	v1.POST("/auth/sms/send", authHandler.SendSMS)
	v1.POST("/auth/login", authHandler.Login)
	v1.POST("/auth/register", authHandler.Register)
	v1.GET("/sns/check", snsHandler.Check)
	v1.POST("/wallet/nonce", walletHandler.Nonce)
	v1.POST("/wallet/verify-sns", walletHandler.VerifySNS)
	v1.GET("/stats/user-count", statsHandler.UserCount)
	v1.GET("/user/profile", profileHandler.Get, middleware.JWT(jwtMgr))
	v1.POST("/user/membership/renew", membershipHandler.Renew, middleware.JWT(jwtMgr))
	v1.POST("/payments/orders", paymentHandler.CreateOrder, middleware.JWT(jwtMgr))
	v1.GET("/payments/orders", paymentHandler.ListOrders, middleware.JWT(jwtMgr))
	v1.GET("/payments/orders/:order_no", paymentHandler.GetOrder, middleware.JWT(jwtMgr))
	v1.POST("/payments/orders/:order_no/close", paymentHandler.CloseOrder, middleware.JWT(jwtMgr))
	v1.POST("/payments/wechat/notify", paymentHandler.WechatNotify)
	if cfg.WechatPay.DevMode {
		v1.POST("/payments/dev/simulate-pay", paymentHandler.SimulatePayDev, middleware.JWT(jwtMgr))
	}

	return e, nil
}
