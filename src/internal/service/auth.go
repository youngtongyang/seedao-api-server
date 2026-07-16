package service

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/jwtutil"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
	"github.com/seedao/seedao-api-server/internal/pkg/snsutil"
)

type AuthService struct {
	db     *gorm.DB
	sms    *SMSService
	sns    *SNSService
	wallet *WalletService
	jwt    *jwtutil.Manager
}

func NewAuthService(db *gorm.DB, sms *SMSService, sns *SNSService, wallet *WalletService, jwt *jwtutil.Manager) *AuthService {
	return &AuthService{db: db, sms: sms, sns: sns, wallet: wallet, jwt: jwt}
}

type PhoneCheckResult struct {
	Phone      string `json:"phone"`
	Registered bool   `json:"registered"`
}

func (s *AuthService) CheckPhone(phoneNum string) (*PhoneCheckResult, error) {
	if err := phone.Validate(phoneNum); err != nil {
		return nil, ErrInvalidParam
	}

	exists, err := phoneExists(s.db, phoneNum)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPhoneAlreadyExists
	}

	return &PhoneCheckResult{
		Phone:      phoneNum,
		Registered: false,
	}, nil
}

type AuthUser struct {
	ID                  int64  `json:"id"`
	SNSName             string `json:"sns_name"`
	RegisterSource      string `json:"register_source"`
	MembershipExpiresAt string `json:"membership_expires_at"`
}

type AuthResult struct {
	Token     string   `json:"token"`
	TokenExp  int64    `json:"token_exp"`
	User      AuthUser `json:"user"`
}

type RegisterNewSNSInput struct {
	Phone      string
	Code       string
	SNSName    string
}

type RegisterExistingSNSInput struct {
	Phone      string
	Code       string
	Wallet     string
	WalletType string
	Domain     string
	Message    string
	Signature  string
}

func (s *AuthService) Login(ctx context.Context, phoneNum, code string) (*AuthResult, error) {
	if err := phone.Validate(phoneNum); err != nil {
		return nil, ErrInvalidParam
	}
	if err := s.sms.ValidateCode(ctx, phoneNum, code); err != nil {
		return nil, err
	}

	user, err := getUserByPhone(s.db, phoneNum)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrPhoneNotRegistered
	}

	if err := s.sms.ConsumeCode(ctx, phoneNum); err != nil {
		return nil, err
	}

	return s.buildAuthResult(user)
}

func (s *AuthService) RegisterNewSNS(ctx context.Context, input RegisterNewSNSInput) (*AuthResult, error) {
	if err := phone.Validate(input.Phone); err != nil {
		return nil, ErrInvalidParam
	}
	if err := snsutil.ValidateFormat(input.SNSName); err != nil {
		return nil, ErrInvalidParam
	}

	exists, err := phoneExists(s.db, input.Phone)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPhoneAlreadyExists
	}

	name := snsutil.Normalize(input.SNSName)
	available, err := s.sns.IsAvailable(name)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, ErrSNSUnavailable
	}

	if err := s.sms.Verify(ctx, input.Phone, input.Code); err != nil {
		return nil, err
	}

	now := time.Now()
	var user model.User

	err = s.db.Transaction(func(tx *gorm.DB) error {
		user = model.User{
			RegisterSource: model.RegisterSourceNewSNS,
			Status:         model.UserStatusActive,
			// MembershipExpiresAt stays nil until payment succeeds.
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.UserPhone{
			UserID:     user.ID,
			Phone:      input.Phone,
			PhoneHash:  phone.Hash(input.Phone),
			VerifiedAt: now,
			CreatedAt:  now,
		}).Error; err != nil {
			return err
		}

		reservedAt := now
		return tx.Create(&model.SNSName{
			Name:       name,
			UserID:     user.ID,
			MintStatus: model.MintStatusReserved,
			ReservedAt: &reservedAt,
			CreatedAt:  now,
		}).Error
	})
	if err != nil {
		return nil, err
	}

	user.SNSName = &model.SNSName{Name: name}
	return s.buildAuthResult(&user)
}

func (s *AuthService) RegisterExistingSNS(ctx context.Context, input RegisterExistingSNSInput) (*AuthResult, error) {
	if err := phone.Validate(input.Phone); err != nil {
		return nil, ErrInvalidParam
	}
	if err := s.sms.Verify(ctx, input.Phone, input.Code); err != nil {
		return nil, err
	}

	exists, err := phoneExists(s.db, input.Phone)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrPhoneAlreadyExists
	}

	walletAddr := phone.NormalizeWallet(input.Wallet)

	if err := s.wallet.VerifyWalletSignature(ctx, VerifySNSInput{
		Wallet:    input.Wallet,
		Domain:    input.Domain,
		Message:   input.Message,
		Signature: input.Signature,
	}); err != nil {
		return nil, err
	}

	now := time.Now()
	var user model.User
	var snsName string

	err = s.db.Transaction(func(tx *gorm.DB) error {
		registry, err := s.sns.getRegistryByWalletForUpdate(tx, walletAddr)
		if err != nil {
			return err
		}
		if registry == nil {
			return ErrWalletNoSNS
		}
		snsName = registry.SNSName

		user = model.User{
			RegisterSource: model.RegisterSourceExistingSNS,
			Status:         model.UserStatusActive,
			// MembershipExpiresAt stays nil until payment succeeds.
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.UserPhone{
			UserID:     user.ID,
			Phone:      input.Phone,
			PhoneHash:  phone.Hash(input.Phone),
			VerifiedAt: now,
			CreatedAt:  now,
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(&model.SNSName{
			Name:       snsName,
			UserID:     user.ID,
			MintStatus: model.MintStatusOnChain,
			Wallet:     &walletAddr,
			CreatedAt:  now,
		}).Error; err != nil {
			return err
		}

		registry.Claimed = true
		registry.UpdatedAt = now
		return tx.Save(registry).Error
	})
	if err != nil {
		return nil, err
	}

	user.SNSName = &model.SNSName{Name: snsName, Wallet: &walletAddr}
	return s.buildAuthResult(&user)
}

func (s *AuthService) buildAuthResult(user *model.User) (*AuthResult, error) {
	token, exp, err := s.jwt.Sign(user.ID)
	if err != nil {
		return nil, err
	}

	snsName := ""
	if user.SNSName != nil {
		snsName = user.SNSName.Name
	}

	return &AuthResult{
		Token:    token,
		TokenExp: exp,
		User: AuthUser{
			ID:                  user.ID,
			SNSName:             snsName,
			RegisterSource:      user.RegisterSource,
			MembershipExpiresAt: formatMembershipExpiry(user.MembershipExpiresAt),
		},
	}, nil
}
