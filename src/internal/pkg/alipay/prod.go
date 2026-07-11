package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/smartwalle/alipay/v3"
)

type prodClient struct {
	client    *alipay.Client
	notifyURL string
	returnURL string
	quitURL   string
}

func newProdClient(cfg Config) (*prodClient, error) {
	privateKey, err := os.ReadFile(cfg.PrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load alipay private key: %w", err)
	}

	publicKey, err := os.ReadFile(cfg.AlipayPublicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load alipay public key: %w", err)
	}

	client, err := alipay.New(cfg.AppID, string(privateKey), true)
	if err != nil {
		return nil, fmt.Errorf("create alipay client: %w", err)
	}
	if err := client.LoadAliPayPublicKey(string(publicKey)); err != nil {
		return nil, fmt.Errorf("load alipay public key into client: %w", err)
	}

	return &prodClient{
		client:    client,
		notifyURL: cfg.NotifyURL,
		returnURL: cfg.ReturnURL,
		quitURL:   cfg.QuitURL,
	}, nil
}

func (c *prodClient) CreateWapOrder(ctx context.Context, input WapOrderInput) (*WapOrderResult, error) {
	var p alipay.TradeWapPay
	p.NotifyURL = c.notifyURL
	p.ReturnURL = c.returnURL
	if c.quitURL != "" {
		p.QuitURL = c.quitURL
	}
	p.Subject = input.Subject
	p.OutTradeNo = input.OrderNo
	p.TotalAmount = fenToYuanString(input.AmountFen)
	p.ProductCode = "QUICK_WAP_WAY"
	p.TimeoutExpress = formatTimeoutExpress(input.ExpireAt)

	payURL, err := c.client.TradeWapPay(p)
	if err != nil {
		return nil, fmt.Errorf("alipay wap pay: %w", err)
	}
	if payURL == nil || strings.TrimSpace(payURL.String()) == "" {
		return nil, fmt.Errorf("alipay wap pay: empty pay_url")
	}

	return &WapOrderResult{PayURL: payURL.String()}, nil
}

func (c *prodClient) ParseNotify(req *http.Request) (*NotifyResult, error) {
	if err := req.ParseForm(); err != nil {
		return nil, fmt.Errorf("parse alipay notify form: %w", err)
	}

	notification, err := c.client.DecodeNotification(context.Background(), req.Form)
	if err != nil {
		return nil, fmt.Errorf("decode alipay notify: %w", err)
	}

	switch notification.TradeStatus {
	case alipay.TradeStatusSuccess, alipay.TradeStatusFinished:
	default:
		return nil, fmt.Errorf("alipay notify: unexpected trade_status %s", notification.TradeStatus)
	}

	if notification.OutTradeNo == "" {
		return nil, fmt.Errorf("alipay notify: missing out_trade_no")
	}
	if notification.TradeNo == "" {
		return nil, fmt.Errorf("alipay notify: missing trade_no")
	}

	amountFen, err := yuanStringToFen(notification.TotalAmount)
	if err != nil {
		return nil, fmt.Errorf("alipay notify: parse total_amount: %w", err)
	}

	successTime := time.Now()
	if notification.GmtPayment != "" {
		if parsed, err := time.ParseInLocation("2006-01-02 15:04:05", notification.GmtPayment, time.Local); err == nil {
			successTime = parsed
		}
	}

	rawJSON, _ := json.Marshal(notification)

	return &NotifyResult{
		OutTradeNo:    notification.OutTradeNo,
		TransactionID: notification.TradeNo,
		SuccessTime:   successTime,
		AmountFen:     amountFen,
		RawJSON:       string(rawJSON),
	}, nil
}

func (c *prodClient) CloseOrder(ctx context.Context, orderNo string) error {
	var p alipay.TradeClose
	p.OutTradeNo = orderNo

	_, err := c.client.TradeClose(ctx, p)
	if err != nil {
		return fmt.Errorf("alipay close order: %w", err)
	}
	return nil
}

func fenToYuanString(fen int64) string {
	yuan := float64(fen) / 100.0
	return strconv.FormatFloat(yuan, 'f', 2, 64)
}

func yuanStringToFen(yuan string) (int64, error) {
	value, err := strconv.ParseFloat(strings.TrimSpace(yuan), 64)
	if err != nil {
		return 0, err
	}
	return int64(math.Round(value * 100)), nil
}

func formatTimeoutExpress(expireAt time.Time) string {
	remaining := time.Until(expireAt)
	if remaining <= 0 {
		return "1m"
	}
	minutes := int(remaining.Minutes())
	if minutes < 1 {
		minutes = 1
	}
	return fmt.Sprintf("%dm", minutes)
}
