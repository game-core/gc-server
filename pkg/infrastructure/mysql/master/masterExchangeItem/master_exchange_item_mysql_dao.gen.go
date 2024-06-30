// Package masterExchangeItem 交換アイテム
package masterExchangeItem

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeItem"
)

type masterExchangeItemMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterExchangeItemMysqlDao(conn *database.MysqlHandler) masterExchangeItem.MasterExchangeItemMysqlRepository {
	return &masterExchangeItemMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterExchangeItemMysqlDao) Find(ctx context.Context, masterExchangeItemId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "Find", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "Find", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindOrNil(ctx context.Context, masterExchangeItemId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindOrNil", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindOrNil", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindByMasterExchangeId(ctx context.Context, masterExchangeId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindOrNilByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindOrNilByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindOrNilByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindOrNilByMasterExchangeId(ctx context.Context, masterExchangeId int64) (*masterExchangeItem.MasterExchangeItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindOrNilByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeItem.MasterExchangeItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindOrNilByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeItemMysqlDao) FindList(ctx context.Context) (masterExchangeItem.MasterExchangeItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchangeItem.MasterExchangeItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchangeItems()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchangeItem.NewMasterExchangeItems()
	for _, t := range ts {
		ms = append(ms, masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeItemMysqlDao) FindListByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (masterExchangeItem.MasterExchangeItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindListByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchangeItem.MasterExchangeItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchangeItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchangeItem.NewMasterExchangeItems()
	for _, t := range ts {
		ms = append(ms, masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindListByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeItemMysqlDao) FindListByMasterExchangeId(ctx context.Context, masterExchangeId int64) (masterExchangeItem.MasterExchangeItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_item", "FindListByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchangeItem.MasterExchangeItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchangeItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_id = ?", masterExchangeId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchangeItem.NewMasterExchangeItems()
	for _, t := range ts {
		ms = append(ms, masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange_item", "FindListByMasterExchangeId", fmt.Sprintf("%d_", masterExchangeId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeItemMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterExchangeItem.MasterExchangeItem) (*masterExchangeItem.MasterExchangeItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchangeItem{
		MasterExchangeItemId: m.MasterExchangeItemId,
		MasterExchangeId:     m.MasterExchangeId,
		MasterItemId:         m.MasterItemId,
		Name:                 m.Name,
		Count:                m.Count,
	}
	res := conn.Model(NewMasterExchangeItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterExchangeItemMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterExchangeItem.MasterExchangeItems) (masterExchangeItem.MasterExchangeItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchangeItems()
	for _, m := range ms {
		t := &MasterExchangeItem{
			MasterExchangeItemId: m.MasterExchangeItemId,
			MasterExchangeId:     m.MasterExchangeId,
			MasterItemId:         m.MasterItemId,
			Name:                 m.Name,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchangeItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeItemMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterExchangeItem.MasterExchangeItem) (*masterExchangeItem.MasterExchangeItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchangeItem{
		MasterExchangeItemId: m.MasterExchangeItemId,
		MasterExchangeId:     m.MasterExchangeId,
		MasterItemId:         m.MasterItemId,
		Name:                 m.Name,
		Count:                m.Count,
	}
	res := conn.Model(NewMasterExchangeItem()).WithContext(ctx).Select("master_exchange_item_id", "master_exchange_id", "master_item_id", "name", "count").Where("master_exchange_item_id = ?", m.MasterExchangeItemId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchangeItem.SetMasterExchangeItem(t.MasterExchangeItemId, t.MasterExchangeId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterExchangeItemMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterExchangeItem.MasterExchangeItems) (masterExchangeItem.MasterExchangeItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchangeItems()
	for _, m := range ms {
		t := &MasterExchangeItem{
			MasterExchangeItemId: m.MasterExchangeItemId,
			MasterExchangeId:     m.MasterExchangeId,
			MasterItemId:         m.MasterItemId,
			Name:                 m.Name,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchangeItem()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_exchange_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_exchange_id", "master_item_id", "name", "count"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeItemMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterExchangeItem.MasterExchangeItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterExchangeItem()).WithContext(ctx).Where("master_exchange_item_id = ?", m.MasterExchangeItemId).Delete(NewMasterExchangeItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterExchangeItemMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterExchangeItem.MasterExchangeItems) error {
	if len(ms) <= 0 {
		return nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	var ks [][]interface{}
	for _, m := range ms {
		ks = append(ks, []interface{}{m.MasterExchangeItemId})
	}

	res := conn.Model(NewMasterExchangeItem()).WithContext(ctx).Where("(master_exchange_item_id) IN ?", ks).Delete(NewMasterExchangeItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
