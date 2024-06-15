// Package userLoginBonus ユーザーログインボーナス
package userLoginBonus

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/userLoginBonus"
)

type userLoginBonusDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserLoginBonusDao(conn *database.MysqlHandler) userLoginBonus.UserLoginBonusMysqlRepository {
	return &userLoginBonusDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userLoginBonusDao) Find(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindOrNil(ctx context.Context, userId string, masterLoginBonusId int64) (*userLoginBonus.UserLoginBonus, error) {
	t := NewUserLoginBonus()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_login_bonus_id = ?", masterLoginBonusId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) FindList(ctx context.Context, userId string) (userLoginBonus.UserLoginBonuses, error) {
	ts := NewUserLoginBonuses()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userLoginBonus.NewUserLoginBonuses()
	for _, t := range ts {
		ms = append(ms, userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt))
	}

	return ms, nil
}

func (s *userLoginBonusDao) Create(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserLoginBonus{
		UserId:             m.UserId,
		MasterLoginBonusId: m.MasterLoginBonusId,
		ReceivedAt:         m.ReceivedAt,
	}
	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms userLoginBonus.UserLoginBonuses) (userLoginBonus.UserLoginBonuses, error) {
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

	ts := NewUserLoginBonuses()
	for _, m := range ms {
		t := &UserLoginBonus{
			UserId:             m.UserId,
			MasterLoginBonusId: m.MasterLoginBonusId,
			ReceivedAt:         m.ReceivedAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userLoginBonusDao) Update(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserLoginBonus{
		UserId:             m.UserId,
		MasterLoginBonusId: m.MasterLoginBonusId,
		ReceivedAt:         m.ReceivedAt,
	}
	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Select("user_id", "master_login_bonus_id", "received_at").Where("user_id = ?", m.UserId).Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userLoginBonus.SetUserLoginBonus(t.UserId, t.MasterLoginBonusId, t.ReceivedAt), nil
}

func (s *userLoginBonusDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userLoginBonus.UserLoginBonuses) (userLoginBonus.UserLoginBonuses, error) {
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

	ts := NewUserLoginBonuses()
	for _, m := range ms {
		t := &UserLoginBonus{
			UserId:             m.UserId,
			MasterLoginBonusId: m.MasterLoginBonusId,
			ReceivedAt:         m.ReceivedAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserLoginBonus()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "master_login_bonus_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"received_at"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userLoginBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *userLoginBonus.UserLoginBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_login_bonus_id = ?", m.MasterLoginBonusId).Delete(NewUserLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userLoginBonusDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userLoginBonus.UserLoginBonuses) error {
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
		ks = append(ks, []interface{}{m.UserId, m.MasterLoginBonusId})
	}

	res := conn.Model(NewUserLoginBonus()).WithContext(ctx).Where("(user_id, master_login_bonus_id) IN ?", ks).Delete(NewUserLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
