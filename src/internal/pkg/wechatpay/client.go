package wechatpay

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type NativeOrderInput struct {
	OrderNo     string
	Description string
	AmountFen   int64
	ExpireAt    time.Time
}

type NativeOrderResult struct {
	CodeURL string
}

type H5OrderInput struct {
	OrderNo     string
	Description string
	AmountFen   int64
	ExpireAt    time.Time
	ClientIP    string
}

type H5OrderResult struct {
	H5URL string
}

type NotifyResult struct {
	OutTradeNo    string
	TransactionID string
	SuccessTime   time.Time
	AmountFen     int64
	RawJSON       string
}

type Client interface {
	CreateNativeOrder(ctx context.Context, input NativeOrderInput) (*NativeOrderResult, error)
	CreateH5Order(ctx context.Context, input H5OrderInput) (*H5OrderResult, error)
	ParseNotify(req *http.Request) (*NotifyResult, error)
	CloseOrder(ctx context.Context, orderNo string) error
}

type Config struct {
	DevMode        bool
	MchID          string
	AppID          string
	APIv3Key       string
	CertSerialNo   string
	PrivateKeyPath string
	NotifyURL      string
	H5WapName      string
	H5WapURL       string
	H5ReturnURL    string
}

func NewClient(cfg Config) (Client, error) {
	if cfg.DevMode {
		return &DevClient{h5ReturnURL: cfg.H5ReturnURL}, nil
	}

	if cfg.MchID == "" || cfg.AppID == "" || cfg.APIv3Key == "" {
		return nil, fmt.Errorf("wechat_pay: mch_id, app_id and api_v3_key are required when dev_mode is false")
	}
	if cfg.CertSerialNo == "" || cfg.PrivateKeyPath == "" {
		return nil, fmt.Errorf("wechat_pay: cert_serial_no and private_key_path are required when dev_mode is false")
	}
	if cfg.NotifyURL == "" {
		return nil, fmt.Errorf("wechat_pay: notify_url is required when dev_mode is false")
	}

	return newProdClient(cfg)
}

func appendH5RedirectURL(h5URL, returnURL string) string {
	if returnURL == "" {
		return h5URL
	}
	sep := "&"
	if !containsQuery(h5URL) {
		sep = "?"
	}
	return h5URL + sep + "redirect_url=" + url.QueryEscape(returnURL)
}

func containsQuery(rawURL string) bool {
	u, err := url.Parse(rawURL)
	return err == nil && u.RawQuery != ""
}
