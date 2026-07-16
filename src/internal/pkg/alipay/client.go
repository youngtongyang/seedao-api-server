package alipay

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type WapOrderInput struct {
	OrderNo   string
	Subject   string
	AmountFen int64
	ExpireAt  time.Time
}

type WapOrderResult struct {
	PayURL string
}

// PageOrderInput is for PC 电脑网站支付 (alipay.trade.page.pay).
type PageOrderInput struct {
	OrderNo   string
	Subject   string
	AmountFen int64
	ExpireAt  time.Time
}

type PageOrderResult struct {
	PayURL string
}

type NotifyResult struct {
	OutTradeNo    string
	TransactionID string
	SuccessTime   time.Time
	AmountFen     int64
	RawJSON       string
}

type Client interface {
	CreateWapOrder(ctx context.Context, input WapOrderInput) (*WapOrderResult, error)
	CreatePageOrder(ctx context.Context, input PageOrderInput) (*PageOrderResult, error)
	ParseNotify(req *http.Request) (*NotifyResult, error)
	CloseOrder(ctx context.Context, orderNo string) error
}

type Config struct {
	DevMode             bool
	AppID               string
	PrivateKeyPath      string
	AlipayPublicKeyPath string
	NotifyURL           string
	ReturnURL           string
	QuitURL             string
}

func NewClient(cfg Config) (Client, error) {
	if cfg.DevMode {
		return &DevClient{returnURL: cfg.ReturnURL}, nil
	}

	if cfg.AppID == "" {
		return nil, fmt.Errorf("alipay: app_id is required when dev_mode is false")
	}
	if cfg.PrivateKeyPath == "" || cfg.AlipayPublicKeyPath == "" {
		return nil, fmt.Errorf("alipay: private_key_path and alipay_public_key_path are required when dev_mode is false")
	}
	if cfg.NotifyURL == "" || cfg.ReturnURL == "" {
		return nil, fmt.Errorf("alipay: notify_url and return_url are required when dev_mode is false")
	}

	return newProdClient(cfg)
}
