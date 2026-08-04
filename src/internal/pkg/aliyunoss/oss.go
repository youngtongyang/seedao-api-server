package aliyunoss

import (
	"bytes"
	"fmt"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/seedao/seedao-api-server/internal/config"
)

// Uploader uploads object bytes and returns a public URL.
type Uploader interface {
	Configured() bool
	Upload(objectKey string, body []byte, contentType string) (publicURL string, err error)
}

// Client wraps Aliyun OSS.
type Client struct {
	bucket        *oss.Bucket
	publicBaseURL string
	prefix        string
}

// NewClient creates an OSS uploader. Empty credentials yield a Disabled client.
func NewClient(cfg config.AliyunOSSConfig) (Uploader, error) {
	if cfg.Endpoint == "" || cfg.Bucket == "" || cfg.AccessKeyID == "" || cfg.AccessKeySecret == "" {
		return &Disabled{}, nil
	}

	endpoint := stripScheme(cfg.Endpoint)
	region := strings.TrimSpace(cfg.Region)
	if region == "" {
		region = regionFromEndpoint(endpoint)
	}
	if region == "" {
		return nil, fmt.Errorf("aliyun oss: region is required (set aliyun_oss.region, e.g. cn-heyuan)")
	}

	// New buckets forbid Signature V1; use AuthV4 + Region.
	client, err := oss.New(endpoint, cfg.AccessKeyID, cfg.AccessKeySecret,
		oss.AuthVersion(oss.AuthV4),
		oss.Region(region),
	)
	if err != nil {
		return nil, fmt.Errorf("aliyun oss: %w", err)
	}
	bucket, err := client.Bucket(cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("aliyun oss bucket: %w", err)
	}

	prefix := strings.Trim(cfg.Prefix, "/")
	if prefix == "" {
		prefix = "import"
	}

	publicBase := strings.TrimRight(cfg.PublicBaseURL, "/")
	if publicBase == "" {
		publicBase = fmt.Sprintf("https://%s.%s", cfg.Bucket, endpoint)
	}

	return &Client{
		bucket:        bucket,
		publicBaseURL: publicBase,
		prefix:        prefix,
	}, nil
}

func stripScheme(endpoint string) string {
	endpoint = strings.TrimSpace(endpoint)
	endpoint = strings.TrimPrefix(endpoint, "https://")
	endpoint = strings.TrimPrefix(endpoint, "http://")
	return endpoint
}

// regionFromEndpoint maps oss-cn-heyuan.aliyuncs.com -> cn-heyuan.
func regionFromEndpoint(endpoint string) string {
	host := strings.ToLower(stripScheme(endpoint))
	if i := strings.Index(host, ".aliyuncs.com"); i > 0 {
		host = host[:i]
	}
	if strings.HasPrefix(host, "oss-") {
		return strings.TrimPrefix(host, "oss-")
	}
	return ""
}

func (c *Client) Configured() bool { return true }

func (c *Client) Upload(objectKey string, body []byte, contentType string) (string, error) {
	key := path.Join(c.prefix, strings.TrimLeft(objectKey, "/"))
	opts := []oss.Option{}
	if contentType != "" {
		opts = append(opts, oss.ContentType(contentType))
	}
	if err := c.bucket.PutObject(key, bytes.NewReader(body), opts...); err != nil {
		return "", err
	}
	return c.publicBaseURL + "/" + key, nil
}

// Disabled is a no-op uploader used when OSS is not configured.
type Disabled struct{}

func (d *Disabled) Configured() bool { return false }

func (d *Disabled) Upload(objectKey string, body []byte, contentType string) (string, error) {
	return "", fmt.Errorf("aliyun oss is not configured")
}
