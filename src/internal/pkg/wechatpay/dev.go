package wechatpay

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type DevClient struct {
	h5ReturnURL string
}

func (d *DevClient) CreateNativeOrder(_ context.Context, input NativeOrderInput) (*NativeOrderResult, error) {
	codeURL := fmt.Sprintf("weixin://wxpay/bizpayurl?pr=dev_%s", input.OrderNo)
	log.Printf("[dev] wechat native order %s amount=%d code_url=%s", input.OrderNo, input.AmountFen, codeURL)
	return &NativeOrderResult{CodeURL: codeURL}, nil
}

func (d *DevClient) CreateH5Order(_ context.Context, input H5OrderInput) (*H5OrderResult, error) {
	h5URL := fmt.Sprintf("https://wx.tenpay.com/cgi-bin/mmpayweb-bin/checkmweb?prepay_id=dev_%s", input.OrderNo)
	h5URL = appendH5RedirectURL(h5URL, d.h5ReturnURL)
	log.Printf("[dev] wechat h5 order %s amount=%d client_ip=%s h5_url=%s", input.OrderNo, input.AmountFen, input.ClientIP, h5URL)
	return &H5OrderResult{H5URL: h5URL}, nil
}

func (d *DevClient) ParseNotify(_ *http.Request) (*NotifyResult, error) {
	return nil, fmt.Errorf("wechat notify is not used in dev_mode; use POST /v1/payments/dev/simulate-pay")
}

func (d *DevClient) CloseOrder(_ context.Context, orderNo string) error {
	log.Printf("[dev] wechat close order %s", orderNo)
	return nil
}
