package aliyunsms

import (
	"fmt"
	"log"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/seedao/seedao-api-server/internal/config"
)

type Sender interface {
	Send(phone, code string) error
}

type Client struct {
	client       *dysmsapi.Client
	signName     string
	templateCode string
}

func NewClient(cfg config.AliyunSMSConfig) (Sender, error) {
	if cfg.DevMode {
		return &DevSender{}, nil
	}

	if cfg.AccessKeyID == "" || cfg.AccessKeySecret == "" {
		return nil, fmt.Errorf("aliyun_sms: access_key_id and access_key_secret are required when dev_mode is false")
	}
	if cfg.SignName == "" || cfg.TemplateCode == "" {
		return nil, fmt.Errorf("aliyun_sms: sign_name and template_code are required when dev_mode is false")
	}

	c, err := dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(cfg.AccessKeyID),
		AccessKeySecret: tea.String(cfg.AccessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	})
	if err != nil {
		return nil, fmt.Errorf("create aliyun sms client: %w", err)
	}

	return &Client{
		client:       c,
		signName:     cfg.SignName,
		templateCode: cfg.TemplateCode,
	}, nil
}

func (c *Client) Send(phone, code string) error {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(c.signName),
		TemplateCode:  tea.String(c.templateCode),
		TemplateParam: tea.String(fmt.Sprintf(`{"code":"%s"}`, code)),
	}
	resp, err := c.client.SendSms(req)
	if err != nil {
		return err
	}
	if resp.Body == nil || tea.StringValue(resp.Body.Code) != "OK" {
		msg := "unknown error"
		if resp.Body != nil && resp.Body.Message != nil {
			msg = tea.StringValue(resp.Body.Message)
		}
		return fmt.Errorf("aliyun sms failed: %s", msg)
	}
	return nil
}

type DevSender struct{}

func (d *DevSender) Send(phone, code string) error {
	log.Printf("[dev] sms to %s code=%s", phone, code)
	return nil
}
