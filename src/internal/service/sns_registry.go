package service

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
)

func (s *SNSService) getRegistryByWallet(walletAddr string) (*model.SNSChainRegistry, error) {
	var registry model.SNSChainRegistry
	err := s.db.Where("LOWER(wallet) = ? AND claimed = ?", strings.ToLower(walletAddr), false).First(&registry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &registry, nil
}

func (s *SNSService) getRegistryByWalletForUpdate(tx *gorm.DB, walletAddr string) (*model.SNSChainRegistry, error) {
	var registry model.SNSChainRegistry
	err := tx.Set("gorm:query_option", "FOR UPDATE").
		Where("LOWER(wallet) = ? AND claimed = ?", phone.NormalizeWallet(walletAddr), false).
		First(&registry).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &registry, nil
}
