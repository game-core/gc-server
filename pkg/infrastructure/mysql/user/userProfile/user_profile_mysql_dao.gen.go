// Package userProfile ユーザープロフィール
package userProfile

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/profile/userProfile"
)

type userProfileMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserProfileMysqlDao(conn *database.MysqlHandler) userProfile.UserProfileMysqlRepository {
	return &userProfileMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userProfileMysqlDao) Find(ctx context.Context, userId string) (*userProfile.UserProfile, error) {
	t := NewUserProfile()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileMysqlDao) FindOrNil(ctx context.Context, userId string) (*userProfile.UserProfile, error) {
	t := NewUserProfile()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileMysqlDao) FindList(ctx context.Context, userId string) (userProfile.UserProfiles, error) {
	ts := NewUserProfiles()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userProfile.NewUserProfiles()
	for _, t := range ts {
		ms = append(ms, userProfile.SetUserProfile(t.UserId, t.Name, t.Content))
	}

	return ms, nil
}

func (s *userProfileMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) (*userProfile.UserProfile, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserProfile{
		UserId:  m.UserId,
		Name:    m.Name,
		Content: m.Content,
	}
	res := conn.Model(NewUserProfile()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms userProfile.UserProfiles) (userProfile.UserProfiles, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	fms := ms[0]
	for _, m := range ms {
		if m.UserId != fms.UserId {
			return nil, errors.NewError("userId is invalid")
		}
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteMysqlConn
	}

	ts := NewUserProfiles()
	for _, m := range ms {
		t := &UserProfile{
			UserId:  m.UserId,
			Name:    m.Name,
			Content: m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserProfile()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userProfileMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) (*userProfile.UserProfile, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserProfile{
		UserId:  m.UserId,
		Name:    m.Name,
		Content: m.Content,
	}
	res := conn.Model(NewUserProfile()).WithContext(ctx).Select("user_id", "name", "content").Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userProfile.SetUserProfile(t.UserId, t.Name, t.Content), nil
}

func (s *userProfileMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userProfile.UserProfiles) (userProfile.UserProfiles, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	fms := ms[0]
	for _, m := range ms {
		if m.UserId != fms.UserId {
			return nil, errors.NewError("userId is invalid")
		}
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteMysqlConn
	}

	ts := NewUserProfiles()
	for _, m := range ms {
		t := &UserProfile{
			UserId:  m.UserId,
			Name:    m.Name,
			Content: m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserProfile()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "content"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userProfileMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *userProfile.UserProfile) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserProfile()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserProfile())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userProfileMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userProfile.UserProfiles) error {
	if len(ms) <= 0 {
		return nil
	}

	fms := ms[0]
	for _, m := range ms {
		if m.UserId != fms.UserId {
			return errors.NewError("userId is invalid")
		}
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(fms.UserId)].WriteMysqlConn
	}

	var ks [][]interface{}
	for _, m := range ms {
		ks = append(ks, []interface{}{m.UserId})
	}

	res := conn.Model(NewUserProfile()).WithContext(ctx).Where("(user_id) IN ?", ks).Delete(NewUserProfile())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
