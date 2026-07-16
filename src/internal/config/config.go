package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	Redis      RedisConfig      `yaml:"redis"`
	JWT        JWTConfig        `yaml:"jwt"`
	AliyunSMS  AliyunSMSConfig  `yaml:"aliyun_sms"`
	SIWE       SIWEConfig       `yaml:"siwe"`
	Membership MembershipConfig `yaml:"membership"`
	WechatPay  WechatPayConfig  `yaml:"wechat_pay"`
	Alipay     AlipayConfig     `yaml:"alipay"`
	Governance GovernanceConfig `yaml:"governance"`
}

type ServerConfig struct {
	Addr  string `yaml:"addr"`
	Debug bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret   string `yaml:"secret"`
	ExpHours int    `yaml:"exp_hours"`
}

type AliyunSMSConfig struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	SignName        string `yaml:"sign_name"`
	TemplateCode    string `yaml:"template_code"`
	DevMode         bool   `yaml:"dev_mode"`
}

type SIWEConfig struct {
	ChainID int64 `yaml:"chain_id"`
}

// MembershipConfig holds product pricing / order metadata shared by WeChat and Alipay.
type MembershipConfig struct {
	MembershipPriceFen int64  `yaml:"membership_price_fen"`
	OrderExpireMinutes int    `yaml:"order_expire_minutes"`
	Description        string `yaml:"description"`
}

type AlipayConfig struct {
	DevMode             bool   `yaml:"dev_mode"`
	AppID               string `yaml:"app_id"`
	PrivateKeyPath      string `yaml:"private_key_path"`
	AlipayPublicKeyPath string `yaml:"alipay_public_key_path"`
	NotifyURL           string `yaml:"notify_url"`
	ReturnURL           string `yaml:"return_url"`
	QuitURL             string `yaml:"quit_url"`
}

type WechatPayConfig struct {
	DevMode             bool   `yaml:"dev_mode"`
	MchID               string `yaml:"mch_id"`
	AppID               string `yaml:"app_id"`
	APIv3Key            string `yaml:"api_v3_key"`
	CertSerialNo        string `yaml:"cert_serial_no"`
	PrivateKeyPath      string `yaml:"private_key_path"`
	WechatPublicKeyID   string `yaml:"wechat_public_key_id"`
	WechatPublicKeyPath string `yaml:"wechat_public_key_path"`
	NotifyURL           string `yaml:"notify_url"`
	H5WapName           string `yaml:"h5_wap_name"`
	H5WapURL            string `yaml:"h5_wap_url"`
	H5ReturnURL         string `yaml:"h5_return_url"`
}

type GovernanceConfig struct {
	QuarterNodeTotal           int      `yaml:"quarter_node_total"`
	HallProposalSNS            []string `yaml:"hall_proposal_sns"`
	NodeConsensusProposalSNS   []string `yaml:"node_consensus_proposal_sns"`
	HallVoterSNS               []string `yaml:"hall_voter_sns"`
	NodeVoterSNS               []string `yaml:"node_voter_sns"`
	DefaultPublicitySeconds    int      `yaml:"default_publicity_seconds"`
	DefaultVoteDurationSeconds int      `yaml:"default_vote_duration_seconds"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if cfg.Server.Addr == "" {
		cfg.Server.Addr = ":8080"
	}
	if cfg.JWT.ExpHours == 0 {
		cfg.JWT.ExpHours = 168
	}
	if cfg.SIWE.ChainID == 0 {
		cfg.SIWE.ChainID = 1
	}
	if cfg.Membership.MembershipPriceFen == 0 {
		cfg.Membership.MembershipPriceFen = 9900
	}
	if cfg.Membership.OrderExpireMinutes == 0 {
		cfg.Membership.OrderExpireMinutes = 30
	}
	if cfg.Membership.Description == "" {
		cfg.Membership.Description = "SEEDAO 年度会员"
	}
	if cfg.WechatPay.H5WapName == "" {
		cfg.WechatPay.H5WapName = "SeeDAO"
	}
	if cfg.WechatPay.MchID == "" {
		cfg.WechatPay.DevMode = true
	}
	if cfg.Alipay.AppID == "" {
		cfg.Alipay.DevMode = true
	}
	if cfg.Alipay.ReturnURL == "" {
		cfg.Alipay.ReturnURL = cfg.WechatPay.H5ReturnURL
	}
	if cfg.Redis.Password == "" {
		return nil, fmt.Errorf("redis.password is required")
	}
	if cfg.Governance.DefaultPublicitySeconds == 0 {
		cfg.Governance.DefaultPublicitySeconds = 432000
	}
	if cfg.Governance.DefaultVoteDurationSeconds == 0 {
		cfg.Governance.DefaultVoteDurationSeconds = 172800
	}

	return &cfg, nil
}
