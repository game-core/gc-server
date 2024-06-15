// Package masterExchangeCost 交換コスト
package masterExchangeCost

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeCost"
)

type masterExchangeCostDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterExchangeCostDao(conn *database.MysqlHandler) masterExchangeCost.MasterExchangeCostMysqlRepository {
	return &masterExchangeCostDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterExchangeCostDao) Find(ctx context.Context, masterExchangeCostId int64) (*masterExchangeCost.MasterExchangeCost, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "Find", fmt.Sprintf("%d_", masterExchangeCostId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeCost.MasterExchangeCost); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeCost()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_cost_id = ?", masterExchangeCostId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "Find", fmt.Sprintf("%d_", masterExchangeCostId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeCostDao) FindOrNil(ctx context.Context, masterExchangeCostId int64) (*masterExchangeCost.MasterExchangeCost, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "FindOrNil", fmt.Sprintf("%d_", masterExchangeCostId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeCost.MasterExchangeCost); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeCost()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_cost_id = ?", masterExchangeCostId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "FindOrNil", fmt.Sprintf("%d_", masterExchangeCostId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeCostDao) FindByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*masterExchangeCost.MasterExchangeCost, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "FindByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeCost.MasterExchangeCost); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeCost()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "FindByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeCostDao) FindOrNilByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*masterExchangeCost.MasterExchangeCost, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "FindOrNilByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchangeCost.MasterExchangeCost); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchangeCost()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "FindOrNilByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeCostDao) FindList(ctx context.Context) (masterExchangeCost.MasterExchangeCosts, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchangeCost.MasterExchangeCosts); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchangeCosts()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchangeCost.NewMasterExchangeCosts()
	for _, t := range ts {
		ms = append(ms, masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeCostDao) FindListByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (masterExchangeCost.MasterExchangeCosts, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange_cost", "FindListByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchangeCost.MasterExchangeCosts); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchangeCosts()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_item_id = ?", masterExchangeItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchangeCost.NewMasterExchangeCosts()
	for _, t := range ts {
		ms = append(ms, masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange_cost", "FindListByMasterExchangeItemId", fmt.Sprintf("%d_", masterExchangeItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeCostDao) Create(ctx context.Context, tx *gorm.DB, m *masterExchangeCost.MasterExchangeCost) (*masterExchangeCost.MasterExchangeCost, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchangeCost{
		MasterExchangeCostId: m.MasterExchangeCostId,
		MasterExchangeItemId: m.MasterExchangeItemId,
		MasterItemId:         m.MasterItemId,
		Name:                 m.Name,
		Count:                m.Count,
	}
	res := conn.Model(NewMasterExchangeCost()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterExchangeCostDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterExchangeCost.MasterExchangeCosts) (masterExchangeCost.MasterExchangeCosts, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchangeCosts()
	for _, m := range ms {
		t := &MasterExchangeCost{
			MasterExchangeCostId: m.MasterExchangeCostId,
			MasterExchangeItemId: m.MasterExchangeItemId,
			MasterItemId:         m.MasterItemId,
			Name:                 m.Name,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchangeCost()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeCostDao) Update(ctx context.Context, tx *gorm.DB, m *masterExchangeCost.MasterExchangeCost) (*masterExchangeCost.MasterExchangeCost, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchangeCost{
		MasterExchangeCostId: m.MasterExchangeCostId,
		MasterExchangeItemId: m.MasterExchangeItemId,
		MasterItemId:         m.MasterItemId,
		Name:                 m.Name,
		Count:                m.Count,
	}
	res := conn.Model(NewMasterExchangeCost()).WithContext(ctx).Select("master_exchange_cost_id", "master_exchange_item_id", "master_item_id", "name", "count").Where("master_exchange_cost_id = ?", m.MasterExchangeCostId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchangeCost.SetMasterExchangeCost(t.MasterExchangeCostId, t.MasterExchangeItemId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterExchangeCostDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterExchangeCost.MasterExchangeCosts) (masterExchangeCost.MasterExchangeCosts, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchangeCosts()
	for _, m := range ms {
		t := &MasterExchangeCost{
			MasterExchangeCostId: m.MasterExchangeCostId,
			MasterExchangeItemId: m.MasterExchangeItemId,
			MasterItemId:         m.MasterItemId,
			Name:                 m.Name,
			Count:                m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchangeCost()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_exchange_cost_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_exchange_item_id", "master_item_id", "name", "count"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeCostDao) Delete(ctx context.Context, tx *gorm.DB, m *masterExchangeCost.MasterExchangeCost) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterExchangeCost()).WithContext(ctx).Where("master_exchange_cost_id = ?", m.MasterExchangeCostId).Delete(NewMasterExchangeCost())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterExchangeCostDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterExchangeCost.MasterExchangeCosts) error {
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
		ks = append(ks, []interface{}{m.MasterExchangeCostId})
	}

	res := conn.Model(NewMasterExchangeCost()).WithContext(ctx).Where("(master_exchange_cost_id) IN ?", ks).Delete(NewMasterExchangeCost())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
