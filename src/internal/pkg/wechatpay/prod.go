package wechatpay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type prodClient struct {
	mchID       string
	appID       string
	notifyURL   string
	h5WapName   string
	h5WapURL    string
	h5ReturnURL string
	client      *core.Client
	nativeSvc   native.NativeApiService
	h5Svc       h5.H5ApiService
	notifyHdlr  *notify.Handler
}

func newProdClient(cfg Config) (*prodClient, error) {
	privateKey, err := utils.LoadPrivateKeyWithPath(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load wechat private key: %w", err)
	}

	wechatPublicKey, err := utils.LoadPublicKeyWithPath(cfg.WechatPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load wechat public key: %w", err)
	}

	ctx := context.Background()
	opts := []core.ClientOption{
		option.WithWechatPayPublicKeyAuthCipher(
			cfg.MchID,
			cfg.CertSerialNo,
			privateKey,
			cfg.WechatPublicKeyID,
			wechatPublicKey,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("create wechat pay client: %w", err)
	}

	notifyHandler := notify.NewNotifyHandler(
		cfg.APIv3Key,
		verifiers.NewSHA256WithRSAPubkeyVerifier(cfg.WechatPublicKeyID, *wechatPublicKey),
	)

	return &prodClient{
		mchID:       cfg.MchID,
		appID:       cfg.AppID,
		notifyURL:   cfg.NotifyURL,
		h5WapName:   cfg.H5WapName,
		h5WapURL:    cfg.H5WapURL,
		h5ReturnURL: cfg.H5ReturnURL,
		client:      client,
		nativeSvc:   native.NativeApiService{Client: client},
		h5Svc:       h5.H5ApiService{Client: client},
		notifyHdlr:  notifyHandler,
	}, nil
}

func (c *prodClient) CreateNativeOrder(ctx context.Context, input NativeOrderInput) (*NativeOrderResult, error) {
	resp, _, err := c.nativeSvc.Prepay(ctx, native.PrepayRequest{
		Appid:       core.String(c.appID),
		Mchid:       core.String(c.mchID),
		Description: core.String(input.Description),
		OutTradeNo:  core.String(input.OrderNo),
		NotifyUrl:   core.String(c.notifyURL),
		Amount: &native.Amount{
			Total:    core.Int64(input.AmountFen),
			Currency: core.String("CNY"),
		},
		TimeExpire: core.Time(input.ExpireAt),
	})
	if err != nil {
		return nil, fmt.Errorf("wechat native prepay: %w", err)
	}
	if resp.CodeUrl == nil || *resp.CodeUrl == "" {
		return nil, fmt.Errorf("wechat native prepay: empty code_url")
	}
	return &NativeOrderResult{CodeURL: *resp.CodeUrl}, nil
}

func (c *prodClient) CreateH5Order(ctx context.Context, input H5OrderInput) (*H5OrderResult, error) {
	if c.h5WapURL == "" {
		return nil, fmt.Errorf("wechat_pay: h5_wap_url is required for h5 payment")
	}
	if input.ClientIP == "" {
		return nil, fmt.Errorf("wechat h5 prepay: client ip is required")
	}

	resp, _, err := c.h5Svc.Prepay(ctx, h5.PrepayRequest{
		Appid:       core.String(c.appID),
		Mchid:       core.String(c.mchID),
		Description: core.String(input.Description),
		OutTradeNo:  core.String(input.OrderNo),
		NotifyUrl:   core.String(c.notifyURL),
		Amount: &h5.Amount{
			Total:    core.Int64(input.AmountFen),
			Currency: core.String("CNY"),
		},
		TimeExpire: core.Time(input.ExpireAt),
		SceneInfo: &h5.SceneInfo{
			PayerClientIp: core.String(input.ClientIP),
			H5Info: &h5.H5Info{
				Type:    core.String("Wap"),
				AppName: core.String(c.h5WapName),
				AppUrl:  core.String(c.h5WapURL),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("wechat h5 prepay: %w", err)
	}
	if resp.H5Url == nil || *resp.H5Url == "" {
		return nil, fmt.Errorf("wechat h5 prepay: empty h5_url")
	}

	h5URL := appendH5RedirectURL(*resp.H5Url, c.h5ReturnURL)
	return &H5OrderResult{H5URL: h5URL}, nil
}

func (c *prodClient) ParseNotify(req *http.Request) (*NotifyResult, error) {
	transaction := new(payments.Transaction)
	notifyReq, err := c.notifyHdlr.ParseNotifyRequest(context.Background(), req, transaction)
	if err != nil {
		return nil, fmt.Errorf("parse wechat notify: %w", err)
	}

	if transaction.OutTradeNo == nil || *transaction.OutTradeNo == "" {
		return nil, fmt.Errorf("wechat notify: missing out_trade_no")
	}
	if transaction.TransactionId == nil || *transaction.TransactionId == "" {
		return nil, fmt.Errorf("wechat notify: missing transaction_id")
	}
	if transaction.Amount == nil || transaction.Amount.Total == nil {
		return nil, fmt.Errorf("wechat notify: missing amount")
	}

	successTime := time.Now()
	if transaction.SuccessTime != nil {
		if parsed, err := time.Parse(time.RFC3339, *transaction.SuccessTime); err == nil {
			successTime = parsed
		}
	}

	rawJSON, _ := json.Marshal(notifyReq)

	return &NotifyResult{
		OutTradeNo:    *transaction.OutTradeNo,
		TransactionID: *transaction.TransactionId,
		SuccessTime:   successTime,
		AmountFen:     *transaction.Amount.Total,
		RawJSON:       string(rawJSON),
	}, nil
}

func (c *prodClient) CloseOrder(ctx context.Context, orderNo string) error {
	_, err := c.nativeSvc.CloseOrder(ctx, native.CloseOrderRequest{
		OutTradeNo: core.String(orderNo),
		Mchid:      core.String(c.mchID),
	})
	if err != nil {
		return fmt.Errorf("wechat close order: %w", err)
	}
	return nil
}
