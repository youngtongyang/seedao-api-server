package redisstore

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/seedao/seedao-api-server/internal/config"
)

type Store struct {
	client *redis.Client
}

func New(cfg config.RedisConfig) *Store {
	return &Store{
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
	}
}

func (s *Store) Ping(ctx context.Context) error {
	return s.client.Ping(ctx).Err()
}

func (s *Store) Close() error {
	return s.client.Close()
}

func (s *Store) SetSMSCode(ctx context.Context, phoneHash, code string, ttl time.Duration) error {
	return s.client.Set(ctx, "sms:code:"+phoneHash, code, ttl).Err()
}

func (s *Store) GetSMSCode(ctx context.Context, phoneHash string) (string, error) {
	return s.client.Get(ctx, "sms:code:"+phoneHash).Result()
}

func (s *Store) DeleteSMSCode(ctx context.Context, phoneHash string) error {
	return s.client.Del(ctx, "sms:code:"+phoneHash).Err()
}

func (s *Store) SetPhoneRateLimit(ctx context.Context, phoneHash string, ttl time.Duration) error {
	return s.client.Set(ctx, "sms:rate:phone:"+phoneHash, "1", ttl).Err()
}

func (s *Store) PhoneRateLimited(ctx context.Context, phoneHash string) (bool, error) {
	n, err := s.client.Exists(ctx, "sms:rate:phone:"+phoneHash).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func (s *Store) IncrIPRate(ctx context.Context, ip string, ttl time.Duration) (int64, error) {
	key := "sms:rate:ip:" + ip
	pipe := s.client.Pipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, ttl)
	if _, err := pipe.Exec(ctx); err != nil {
		return 0, err
	}
	return incr.Val(), nil
}

func (s *Store) GetSNSCheckCache(ctx context.Context, name string) (string, error) {
	return s.client.Get(ctx, "sns:check:"+name).Result()
}

func (s *Store) SetSNSCheckCache(ctx context.Context, name, value string, ttl time.Duration) error {
	return s.client.Set(ctx, "sns:check:"+name, value, ttl).Err()
}

func (s *Store) SetWalletNonce(ctx context.Context, address, nonce string, ttl time.Duration) error {
	return s.client.Set(ctx, "wallet:nonce:"+address, nonce, ttl).Err()
}

func (s *Store) GetWalletNonce(ctx context.Context, address string) (string, error) {
	return s.client.Get(ctx, "wallet:nonce:"+address).Result()
}

func (s *Store) DeleteWalletNonce(ctx context.Context, address string) error {
	return s.client.Del(ctx, "wallet:nonce:"+address).Err()
}

func WalletNonceKey(address string) string {
	return fmt.Sprintf("wallet:nonce:%s", address)
}

func (s *Store) TryPaymentNotifyLock(ctx context.Context, orderNo string, ttl time.Duration) (bool, error) {
	return s.client.SetNX(ctx, "payment:notify:lock:"+orderNo, "1", ttl).Result()
}

func (s *Store) ReleasePaymentNotifyLock(ctx context.Context, orderNo string) error {
	return s.client.Del(ctx, "payment:notify:lock:"+orderNo).Err()
}
