// Package userAccount ユーザーアカウント
package userAccount

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
)

type userAccountMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserAccountMysqlDao(conn *database.MysqlHandler) userAccount.UserAccountMysqlRepository {
	return &userAccountMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userAccountMysqlDao) Find(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountMysqlDao) FindOrNil(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	t := NewUserAccount()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountMysqlDao) FindList(ctx context.Context, userId string) (userAccount.UserAccounts, error) {
	ts := NewUserAccounts()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userAccount.NewUserAccounts()
	for _, t := range ts {
		ms = append(ms, userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt))
	}

	return ms, nil
}

func (s *userAccountMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccount{
		UserId:   m.UserId,
		Name:     m.Name,
		Password: m.Password,
		LoginAt:  m.LoginAt,
		LogoutAt: m.LogoutAt,
	}
	res := conn.Model(NewUserAccount()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms userAccount.UserAccounts) (userAccount.UserAccounts, error) {
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

	ts := NewUserAccounts()
	for _, m := range ms {
		t := &UserAccount{
			UserId:   m.UserId,
			Name:     m.Name,
			Password: m.Password,
			LoginAt:  m.LoginAt,
			LogoutAt: m.LogoutAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) (*userAccount.UserAccount, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserAccount{
		UserId:   m.UserId,
		Name:     m.Name,
		Password: m.Password,
		LoginAt:  m.LoginAt,
		LogoutAt: m.LogoutAt,
	}
	res := conn.Model(NewUserAccount()).WithContext(ctx).Select("user_id", "name", "password", "login_at", "logout_at").Where("user_id = ?", m.UserId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userAccount.SetUserAccount(t.UserId, t.Name, t.Password, t.LoginAt, t.LogoutAt), nil
}

func (s *userAccountMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userAccount.UserAccounts) (userAccount.UserAccounts, error) {
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

	ts := NewUserAccounts()
	for _, m := range ms {
		t := &UserAccount{
			UserId:   m.UserId,
			Name:     m.Name,
			Password: m.Password,
			LoginAt:  m.LoginAt,
			LogoutAt: m.LogoutAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserAccount()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "password", "login_at", "logout_at"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userAccountMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *userAccount.UserAccount) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("user_id = ?", m.UserId).Delete(NewUserAccount())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userAccountMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userAccount.UserAccounts) error {
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

	res := conn.Model(NewUserAccount()).WithContext(ctx).Where("(user_id) IN ?", ks).Delete(NewUserAccount())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
