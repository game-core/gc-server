// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccountToken"
)

type userAccountTokenMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserAccountTokenMysqlDao(conn *database.MysqlHandler) userAccountToken.UserAccountTokenMysqlRepository {
	return &userAccountTokenMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userAccountTokenMysqlDao) Find(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenMysqlDao) FindOrNil(ctx context.Context, userId string) (*userAccountToken.UserAccountToken, error) {
	t := NewUserAccountToken()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenMysqlDao) FindList(ctx context.Context, userId string) (userAccountToken.UserAccountTokens, error) {
	ts := NewUserAccountTokens()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAccountToken.NewUserAccountTokens()
	for _, t := range ts {
		ms = append(ms, userAccountToken.SetUserAccountToken(t.UserId, t.Token))
	}

	return ms, nil
}

func (s *userAccountTokenMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}
	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAccountToken.UserAccountTokens) (userAccountToken.UserAccountTokens, error) {
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

	ts := NewUserAccountTokens()
	for _, m := range ms {
		t := &UserAccountToken{
			UserId: m.UserId,
			Token:  m.Token,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountTokenMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) (*userAccountToken.UserAccountToken, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccountToken{
		UserId: m.UserId,
		Token:  m.Token,
	}
	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Select("user_id", "token").Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccountToken.SetUserAccountToken(t.UserId, t.Token), nil
}

func (s *userAccountTokenMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userAccountToken.UserAccountTokens) (userAccountToken.UserAccountTokens, error) {
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

	ts := NewUserAccountTokens()
	for _, m := range ms {
		t := &UserAccountToken{
			UserId: m.UserId,
			Token:  m.Token,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccountToken()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountTokenMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *userAccountToken.UserAccountToken) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccountToken())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userAccountTokenMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userAccountToken.UserAccountTokens) error {
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

	res := conn.Model(NewUserAccountToken()).WithContext(ctx).Where("(user_id) IN ?", ks).Delete(NewUserAccountToken())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
