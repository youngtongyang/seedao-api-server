package alipay

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type DevClient struct {
	returnURL string
}

func (d *DevClient) CreateWapOrder(_ context.Context, input WapOrderInput) (*WapOrderResult, error) {
	payURL := fmt.Sprintf("https://openapi.alipay.com/gateway.do?dev_order=%s", url.QueryEscape(input.OrderNo))
	if d.returnURL != "" {
		payURL += "&return_url=" + url.QueryEscape(d.returnURL)
	}
	log.Printf("[dev] alipay wap order %s amount=%d pay_url=%s", input.OrderNo, input.AmountFen, payURL)
	return &WapOrderResult{PayURL: payURL}, nil
}

func (d *DevClient) ParseNotify(_ *http.Request) (*NotifyResult, error) {
	return nil, fmt.Errorf("alipay notify is not used in dev_mode; use POST /v1/payments/dev/simulate-pay")
}

func (d *DevClient) CloseOrder(_ context.Context, orderNo string) error {
	log.Printf("[dev] alipay close order %s", orderNo)
	return nil
}
