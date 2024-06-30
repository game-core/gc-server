// Package masterExchange 交換
package masterExchange

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchange"
)

type masterExchangeMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterExchangeMysqlDao(conn *database.MysqlHandler) masterExchange.MasterExchangeMysqlRepository {
	return &masterExchangeMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterExchangeMysqlDao) Find(ctx context.Context, masterExchangeId int64) (*masterExchange.MasterExchange, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "Find", fmt.Sprintf("%d_", masterExchangeId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchange.MasterExchange); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchange()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "Find", fmt.Sprintf("%d_", masterExchangeId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeMysqlDao) FindOrNil(ctx context.Context, masterExchangeId int64) (*masterExchange.MasterExchange, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "FindOrNil", fmt.Sprintf("%d_", masterExchangeId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchange.MasterExchange); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchange()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_exchange_id = ?", masterExchangeId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "FindOrNil", fmt.Sprintf("%d_", masterExchangeId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeMysqlDao) FindByMasterEventId(ctx context.Context, masterEventId int64) (*masterExchange.MasterExchange, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "FindByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchange.MasterExchange); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchange()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "FindByMasterEventId", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeMysqlDao) FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*masterExchange.MasterExchange, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "FindOrNilByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterExchange.MasterExchange); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterExchange()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "FindOrNilByMasterEventId", fmt.Sprintf("%d_", masterEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterExchangeMysqlDao) FindList(ctx context.Context) (masterExchange.MasterExchanges, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchange.MasterExchanges); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchanges()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchange.NewMasterExchanges()
	for _, t := range ts {
		ms = append(ms, masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeMysqlDao) FindListByMasterEventId(ctx context.Context, masterEventId int64) (masterExchange.MasterExchanges, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_exchange", "FindListByMasterEventId", fmt.Sprintf("%d_", masterEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterExchange.MasterExchanges); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterExchanges()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_event_id = ?", masterEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterExchange.NewMasterExchanges()
	for _, t := range ts {
		ms = append(ms, masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_exchange", "FindListByMasterEventId", fmt.Sprintf("%d_", masterEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterExchangeMysqlDao) Create(ctx context.Context, tx *gorm.DB, m *masterExchange.MasterExchange) (*masterExchange.MasterExchange, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchange{
		MasterExchangeId: m.MasterExchangeId,
		MasterEventId:    m.MasterEventId,
		Name:             m.Name,
	}
	res := conn.Model(NewMasterExchange()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name), nil
}

func (s *masterExchangeMysqlDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterExchange.MasterExchanges) (masterExchange.MasterExchanges, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchanges()
	for _, m := range ms {
		t := &MasterExchange{
			MasterExchangeId: m.MasterExchangeId,
			MasterEventId:    m.MasterEventId,
			Name:             m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchange()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeMysqlDao) Update(ctx context.Context, tx *gorm.DB, m *masterExchange.MasterExchange) (*masterExchange.MasterExchange, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterExchange{
		MasterExchangeId: m.MasterExchangeId,
		MasterEventId:    m.MasterEventId,
		Name:             m.Name,
	}
	res := conn.Model(NewMasterExchange()).WithContext(ctx).Select("master_exchange_id", "master_event_id", "name").Where("master_exchange_id = ?", m.MasterExchangeId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterExchange.SetMasterExchange(t.MasterExchangeId, t.MasterEventId, t.Name), nil
}

func (s *masterExchangeMysqlDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterExchange.MasterExchanges) (masterExchange.MasterExchanges, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterExchanges()
	for _, m := range ms {
		t := &MasterExchange{
			MasterExchangeId: m.MasterExchangeId,
			MasterEventId:    m.MasterEventId,
			Name:             m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterExchange()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_exchange_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_event_id", "name"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterExchangeMysqlDao) Delete(ctx context.Context, tx *gorm.DB, m *masterExchange.MasterExchange) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterExchange()).WithContext(ctx).Where("master_exchange_id = ?", m.MasterExchangeId).Delete(NewMasterExchange())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterExchangeMysqlDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterExchange.MasterExchanges) error {
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
		ks = append(ks, []interface{}{m.MasterExchangeId})
	}

	res := conn.Model(NewMasterExchange()).WithContext(ctx).Where("(master_exchange_id) IN ?", ks).Delete(NewMasterExchange())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
