package app

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"

	_ "github.com/seedao/seedao-api-server/docs"
	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/handler"
	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/aliyunsms"
	"github.com/seedao/seedao-api-server/internal/pkg/alipay"
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
		DevMode:             cfg.WechatPay.DevMode,
		MchID:               cfg.WechatPay.MchID,
		AppID:               cfg.WechatPay.AppID,
		APIv3Key:            cfg.WechatPay.APIv3Key,
		CertSerialNo:        cfg.WechatPay.CertSerialNo,
		PrivateKeyPath:      cfg.WechatPay.PrivateKeyPath,
		WechatPublicKeyID:   cfg.WechatPay.WechatPublicKeyID,
		WechatPublicKeyPath: cfg.WechatPay.WechatPublicKeyPath,
		NotifyURL:           cfg.WechatPay.NotifyURL,
		H5WapName:           cfg.WechatPay.H5WapName,
		H5WapURL:            cfg.WechatPay.H5WapURL,
		H5ReturnURL:         cfg.WechatPay.H5ReturnURL,
	})
	if err != nil {
		return nil, err
	}

	alipayClient, err := alipay.NewClient(alipay.Config{
		DevMode:             cfg.Alipay.DevMode,
		AppID:               cfg.Alipay.AppID,
		PrivateKeyPath:      cfg.Alipay.PrivateKeyPath,
		AlipayPublicKeyPath: cfg.Alipay.AlipayPublicKeyPath,
		NotifyURL:           cfg.Alipay.NotifyURL,
		ReturnURL:           cfg.Alipay.ReturnURL,
		QuitURL:             cfg.Alipay.QuitURL,
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
	couponSvc := service.NewCouponService(db)
	payDevMode := cfg.WechatPay.DevMode || cfg.Alipay.DevMode
	courseSvc := service.NewCourseService(db, wxClient, alipayClient, cfg.Membership, payDevMode)
	paymentSvc := service.NewPaymentService(db, redis, wxClient, alipayClient, cfg.Membership, payDevMode, couponSvc, courseSvc)
	proposalSvc := service.NewProposalService(db, snsSvc, cfg.Governance)

	authHandler := handler.NewAuthHandler(authSvc, smsSvc)
	snsHandler := handler.NewSNSHandler(snsSvc)
	walletHandler := handler.NewWalletHandler(walletSvc, snsSvc)
	profileHandler := handler.NewProfileHandler(profileSvc)
	statsHandler := handler.NewStatsHandler(statsSvc)
	membershipHandler := handler.NewMembershipHandler(membershipSvc)
	paymentHandler := handler.NewPaymentHandler(paymentSvc, payDevMode)
	courseHandler := handler.NewCourseHandler(courseSvc, paymentSvc, payDevMode)
	proposalHandler := handler.NewProposalHandler(proposalSvc, profileSvc)

	e := echo.New()
	e.HideBanner = true
	e.Use(echomw.Recover())
	e.Use(echomw.RequestID())
	e.Use(echomw.Logger())
	e.Use(middleware.DebugMode(cfg.Server.Debug))

	e.GET("/health", handler.Health)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := e.Group("/v1")
	v1.POST("/auth/sms/send", authHandler.SendSMS)
	v1.GET("/auth/phone/check", authHandler.CheckPhone)
	v1.POST("/auth/login", authHandler.Login)
	v1.POST("/auth/register", authHandler.Register)
	v1.GET("/sns/check", snsHandler.Check)
	v1.GET("/sns/by-wallet", snsHandler.LookupByWallet)
	v1.POST("/wallet/nonce", walletHandler.Nonce)
	v1.POST("/wallet/verify-sns", walletHandler.VerifySNS)
	v1.GET("/stats/user-count", statsHandler.UserCount)
	v1.GET("/user/profile", profileHandler.Get, middleware.JWT(jwtMgr))
	v1.POST("/user/membership/renew", membershipHandler.Renew, middleware.JWT(jwtMgr))
	v1.POST("/payments/coupons/preview", paymentHandler.PreviewCoupon, middleware.OptionalJWT(jwtMgr))
	v1.POST("/payments/orders", paymentHandler.CreateOrder, middleware.JWT(jwtMgr))
	v1.GET("/payments/orders", paymentHandler.ListOrders, middleware.JWT(jwtMgr))
	v1.GET("/payments/orders/:order_no", paymentHandler.GetOrder, middleware.JWT(jwtMgr))
	v1.POST("/payments/orders/:order_no/close", paymentHandler.CloseOrder, middleware.JWT(jwtMgr))
	v1.POST("/payments/wechat/notify", paymentHandler.WechatNotify)
	v1.POST("/payments/alipay/notify", paymentHandler.AlipayNotify)
	v1.POST("/courses/checkout/preview", courseHandler.PreviewCheckout, middleware.OptionalJWT(jwtMgr))
	v1.POST("/courses/orders", courseHandler.CreateOrder, middleware.OptionalJWT(jwtMgr))
	v1.GET("/courses/orders/:order_no", courseHandler.GetOrder, middleware.OptionalJWT(jwtMgr))
	v1.POST("/courses/orders/:order_no/close", courseHandler.CloseOrder, middleware.OptionalJWT(jwtMgr))
	v1.POST("/courses/orders/claim", courseHandler.ClaimOrder, middleware.JWT(jwtMgr))
	if cfg.WechatPay.DevMode || cfg.Alipay.DevMode {
		v1.POST("/payments/dev/simulate-pay", paymentHandler.SimulatePayDev, middleware.JWT(jwtMgr))
		v1.POST("/courses/dev/orders/:order_no/simulate-pay", courseHandler.SimulatePayDev, middleware.OptionalJWT(jwtMgr))
	}

	v1.GET("/governance/nodes", proposalHandler.GovernanceNodes)
	v1.GET("/governance/proposal-permissions", proposalHandler.ProposalCreatePermissions, middleware.JWT(jwtMgr))
	v1.GET("/proposals", proposalHandler.List)
	v1.GET("/proposals/:id/vote-summary", proposalHandler.VoteSummary)
	v1.GET("/proposals/:id/votes/me", proposalHandler.GetMyVote, middleware.JWT(jwtMgr))
	v1.GET("/proposals/:id/votes", proposalHandler.ListVotes)
	v1.GET("/proposals/:id/comments", proposalHandler.ListComments)
	v1.GET("/proposals/:id", proposalHandler.Get)
	v1.POST("/proposals", proposalHandler.Create, middleware.JWT(jwtMgr))
	v1.PATCH("/proposals/:id", proposalHandler.Update, middleware.JWT(jwtMgr))
	v1.POST("/proposals/:id/close", proposalHandler.Close, middleware.JWT(jwtMgr))
	v1.POST("/proposals/:id/publish", proposalHandler.Publish, middleware.JWT(jwtMgr))
	v1.POST("/proposals/:id/finalize", proposalHandler.Finalize, middleware.JWT(jwtMgr))
	v1.POST("/proposals/:id/votes", proposalHandler.CastVote, middleware.JWT(jwtMgr))
	v1.POST("/proposals/:id/comments", proposalHandler.CreateComment, middleware.JWT(jwtMgr))

	return e, nil
}
