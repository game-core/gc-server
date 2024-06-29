// Package userExchange ユーザー交換
package userExchange

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
)

type userExchangeMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserExchangeMysqlDao(conn *database.MysqlHandler) userExchange.UserExchangeMysqlRepository {
	return &userExchangeMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userExchangeMysqlDao) Find(ctx context.Context, userId string, masterExchangeId int64) (*userExchange.UserExchange, error) {
	t := NewUserExchange()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userExchange.SetUserExchange(t.UserId, t.MasterExchangeId, t.ResetAt), nil
}

func (s *userExchangeMysqlDao) FindOrNil(ctx context.Context, userId string, masterExchangeId int64) (*userExchange.UserExchange, error) {
	t := NewUserExchange()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userExchange.SetUserExchange(t.UserId, t.MasterExchangeId, t.ResetAt), nil
}

func (s *userExchangeMysqlDao) FindList(ctx context.Context, userId string) (userExchange.UserExchanges, error) {
	ts := NewUserExchanges()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userExchange.NewUserExchanges()
	for _, t := range ts {
		ms = append(ms, userExchange.SetUserExchange(t.UserId, t.MasterExchangeId, t.ResetAt))
	}

	return ms, nil
}

func (s *userExchangeMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *userExchange.UserExchange) (*userExchange.UserExchange, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserExchange{
		UserId:           m.UserId,
		MasterExchangeId: m.MasterExchangeId,
		ResetAt:          m.ResetAt,
	}
	res := conn.Model(NewUserExchange()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userExchange.SetUserExchange(t.UserId, t.MasterExchangeId, t.ResetAt), nil
}

func (s *userExchangeMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms userExchange.UserExchanges) (userExchange.UserExchanges, error) {
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

	ts := NewUserExchanges()
	for _, m := range ms {
		t := &UserExchange{
			UserId:           m.UserId,
			MasterExchangeId: m.MasterExchangeId,
			ResetAt:          m.ResetAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserExchange()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userExchangeMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *userExchange.UserExchange) (*userExchange.UserExchange, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserExchange{
		UserId:           m.UserId,
		MasterExchangeId: m.MasterExchangeId,
		ResetAt:          m.ResetAt,
	}
	res := conn.Model(NewUserExchange()).WithContext(ctx).Select("user_id", "master_exchange_id", "reset_at").Where("user_id = ?", m.UserId).Where("master_exchange_id = ?", m.MasterExchangeId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userExchange.SetUserExchange(t.UserId, t.MasterExchangeId, t.ResetAt), nil
}

func (s *userExchangeMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userExchange.UserExchanges) (userExchange.UserExchanges, error) {
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

	ts := NewUserExchanges()
	for _, m := range ms {
		t := &UserExchange{
			UserId:           m.UserId,
			MasterExchangeId: m.MasterExchangeId,
			ResetAt:          m.ResetAt,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserExchange()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "master_exchange_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"reset_at"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userExchangeMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *userExchange.UserExchange) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserExchange()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_exchange_id = ?", m.MasterExchangeId).Delete(NewUserExchange())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userExchangeMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userExchange.UserExchanges) error {
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
		ks = append(ks, []interface{}{m.UserId, m.MasterExchangeId})
	}

	res := conn.Model(NewUserExchange()).WithContext(ctx).Where("(user_id, master_exchange_id) IN ?", ks).Delete(NewUserExchange())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
