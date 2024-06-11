package profile

import (
	"context"

	profileProto "github.com/game-core/gc-server/api/game/presentation/proto/profile"
	"github.com/game-core/gc-server/api/game/presentation/proto/profile/userProfile"
	"github.com/game-core/gc-server/internal/errors"
	profileService "github.com/game-core/gc-server/pkg/domain/model/profile"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type ProfileUsecase interface {
	Get(ctx context.Context, req *profileProto.ProfileGetRequest) (*profileProto.ProfileGetResponse, error)
	Create(ctx context.Context, req *profileProto.ProfileCreateRequest) (*profileProto.ProfileCreateResponse, error)
	Update(ctx context.Context, req *profileProto.ProfileUpdateRequest) (*profileProto.ProfileUpdateResponse, error)
}

type profileUsecase struct {
	profileService     profileService.ProfileService
	transactionService transactionService.TransactionService
}

func NewProfileUsecase(
	profileService profileService.ProfileService,
	transactionService transactionService.TransactionService,
) ProfileUsecase {
	return &profileUsecase{
		profileService:     profileService,
		transactionService: transactionService,
	}
}

// Get プロフィールを作成する
func (s *profileUsecase) Get(ctx context.Context, req *profileProto.ProfileGetRequest) (*profileProto.ProfileGetResponse, error) {
	result, err := s.profileService.Get(ctx, profileService.SetProfileGetRequest(req.UserId))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Get", err)
	}

	return profileProto.SetProfileGetResponse(
		userProfile.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}

// Create プロフィールを作成する
func (s *profileUsecase) Create(ctx context.Context, req *profileProto.ProfileCreateRequest) (*profileProto.ProfileCreateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	result, err := s.profileService.Create(ctx, tx, profileService.SetProfileCreateRequest(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Create", err)
	}

	return profileProto.SetProfileCreateResponse(
		userProfile.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}

// Update プロフィールを更新する
func (s *profileUsecase) Update(ctx context.Context, req *profileProto.ProfileUpdateRequest) (*profileProto.ProfileUpdateResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	result, err := s.profileService.Update(ctx, tx, profileService.SetProfileUpdateRequest(req.UserId, req.Name, req.Content))
	if err != nil {
		return nil, errors.NewMethodError("s.profileService.Update", err)
	}

	return profileProto.SetProfileUpdateResponse(
		userProfile.SetUserProfile(
			result.UserProfile.UserId,
			result.UserProfile.Name,
			result.UserProfile.Content,
		),
	), nil
}
