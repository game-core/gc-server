// Package userExchangeItem ユーザー交換アイテム
package userExchangeItem

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
)

type userExchangeItemMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserExchangeItemMysqlDao(conn *database.MysqlHandler) userExchangeItem.UserExchangeItemMysqlRepository {
	return &userExchangeItemMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (s *userExchangeItemMysqlDao) Find(ctx context.Context, userId string, masterExchangeItemId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindOrNil(ctx context.Context, userId string, masterExchangeItemId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindOrNilByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindOrNilByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (*userExchangeItem.UserExchangeItem, error) {
	t := NewUserExchangeItem()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) FindList(ctx context.Context, userId string) (userExchangeItem.UserExchangeItems, error) {
	ts := NewUserExchangeItems()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userExchangeItem.NewUserExchangeItems()
	for _, t := range ts {
		ms = append(ms, userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count))
	}

	return ms, nil
}

func (s *userExchangeItemMysqlDao) FindListByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (userExchangeItem.UserExchangeItems, error) {
	ts := NewUserExchangeItems()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_id = ?", masterExchangeId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userExchangeItem.NewUserExchangeItems()
	for _, t := range ts {
		ms = append(ms, userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count))
	}

	return ms, nil
}

func (s *userExchangeItemMysqlDao) FindListByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (userExchangeItem.UserExchangeItems, error) {
	ts := NewUserExchangeItems()
	res := s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(userId)].ReadMysqlConn.WithContext(ctx).Where("user_id = ?", userId).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := userExchangeItem.NewUserExchangeItems()
	for _, t := range ts {
		ms = append(ms, userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count))
	}

	return ms, nil
}

func (s *userExchangeItemMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *userExchangeItem.UserExchangeItem) (*userExchangeItem.UserExchangeItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserExchangeItem{
		UserId:               m.UserId,
		MasterExchangeId:     m.MasterExchangeId,
		MasterExchangeItemId: m.MasterExchangeItemId,
		Count:                m.Count,
	}
	res := conn.Model(NewUserExchangeItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms userExchangeItem.UserExchangeItems) (userExchangeItem.UserExchangeItems, error) {
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

	ts := NewUserExchangeItems()
	for _, m := range ms {
		t := &UserExchangeItem{
			UserId:               m.UserId,
			MasterExchangeId:     m.MasterExchangeId,
			MasterExchangeItemId: m.MasterExchangeItemId,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserExchangeItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userExchangeItemMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *userExchangeItem.UserExchangeItem) (*userExchangeItem.UserExchangeItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	t := &UserExchangeItem{
		UserId:               m.UserId,
		MasterExchangeId:     m.MasterExchangeId,
		MasterExchangeItemId: m.MasterExchangeItemId,
		Count:                m.Count,
	}
	res := conn.Model(NewUserExchangeItem()).WithContext(ctx).Select("user_id", "master_exchange_id", "master_exchange_item_id", "count").Where("user_id = ?", m.UserId).Where("master_exchange_item_id = ?", m.MasterExchangeItemId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return userExchangeItem.SetUserExchangeItem(t.UserId, t.MasterExchangeId, t.MasterExchangeItemId, t.Count), nil
}

func (s *userExchangeItemMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms userExchangeItem.UserExchangeItems) (userExchangeItem.UserExchangeItems, error) {
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

	ts := NewUserExchangeItems()
	for _, m := range ms {
		t := &UserExchangeItem{
			UserId:               m.UserId,
			MasterExchangeId:     m.MasterExchangeId,
			MasterExchangeItemId: m.MasterExchangeItemId,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewUserExchangeItem()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "master_exchange_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_exchange_id", "count"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *userExchangeItemMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *userExchangeItem.UserExchangeItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.ShardMysqlConn.Shards[keys.GetShardKeyByUserId(m.UserId)].WriteMysqlConn
	}

	res := conn.Model(NewUserExchangeItem()).WithContext(ctx).Where("user_id = ?", m.UserId).Where("master_exchange_item_id = ?", m.MasterExchangeItemId).Delete(NewUserExchangeItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *userExchangeItemMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms userExchangeItem.UserExchangeItems) error {
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
		ks = append(ks, []interface{}{m.UserId, m.MasterExchangeItemId})
	}

	res := conn.Model(NewUserExchangeItem()).WithContext(ctx).Where("(user_id, master_exchange_item_id) IN ?", ks).Delete(NewUserExchangeItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
