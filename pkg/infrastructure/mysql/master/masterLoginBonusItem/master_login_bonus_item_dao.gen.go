// Package masterLoginBonusItem ログインボーナスアイテム
package masterLoginBonusItem

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusItem"
)

type masterLoginBonusItemDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterLoginBonusItemDao(conn *database.MysqlHandler) masterLoginBonusItem.MasterLoginBonusItemMysqlRepository {
	return &masterLoginBonusItemDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusItemDao) Find(ctx context.Context, masterLoginBonusItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "Find", fmt.Sprintf("%d_", masterLoginBonusItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_item_id = ?", masterLoginBonusItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "Find", fmt.Sprintf("%d_", masterLoginBonusItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindOrNil(ctx context.Context, masterLoginBonusItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_item_id = ?", masterLoginBonusItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNil", fmt.Sprintf("%d_", masterLoginBonusItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindByMasterItemId(ctx context.Context, masterItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterItemId", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindOrNilByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindOrNilByMasterItemId(ctx context.Context, masterItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterItemId", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonusItem.MasterLoginBonusItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonusItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindOrNilByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusItemDao) FindList(ctx context.Context) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusItem.MasterLoginBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusItem.NewMasterLoginBonusItems()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusItemDao) FindListByMasterLoginBonusScheduleId(ctx context.Context, masterLoginBonusScheduleId int64) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusItem.MasterLoginBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusItem.NewMasterLoginBonusItems()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterLoginBonusScheduleId", fmt.Sprintf("%d_", masterLoginBonusScheduleId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusItemDao) FindListByMasterItemId(ctx context.Context, masterItemId int64) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterItemId", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusItem.MasterLoginBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusItem.NewMasterLoginBonusItems()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterItemId", fmt.Sprintf("%d_", masterItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusItemDao) FindListByMasterLoginBonusScheduleIdAndMasterItemId(ctx context.Context, masterLoginBonusScheduleId int64, masterItemId int64) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonusItem.MasterLoginBonusItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonusItems()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_login_bonus_schedule_id = ?", masterLoginBonusScheduleId).Where("master_item_id = ?", masterItemId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonusItem.NewMasterLoginBonusItems()
	for _, t := range ts {
		ms = append(ms, masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus_item", "FindListByMasterLoginBonusScheduleIdAndMasterItemId", fmt.Sprintf("%d_%d_", masterLoginBonusScheduleId, masterItemId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusItemDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonusItem.MasterLoginBonusItem) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonusItem{
		MasterLoginBonusItemId:     m.MasterLoginBonusItemId,
		MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
		MasterItemId:               m.MasterItemId,
		Name:                       m.Name,
		Count:                      m.Count,
	}
	res := conn.Model(NewMasterLoginBonusItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterLoginBonusItemDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusItem.MasterLoginBonusItems) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonusItems()
	for _, m := range ms {
		t := &MasterLoginBonusItem{
			MasterLoginBonusItemId:     m.MasterLoginBonusItemId,
			MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
			MasterItemId:               m.MasterItemId,
			Name:                       m.Name,
			Count:                      m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonusItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusItemDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonusItem.MasterLoginBonusItem) (*masterLoginBonusItem.MasterLoginBonusItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterLoginBonusItem{
		MasterLoginBonusItemId:     m.MasterLoginBonusItemId,
		MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
		MasterItemId:               m.MasterItemId,
		Name:                       m.Name,
		Count:                      m.Count,
	}
	res := conn.Model(NewMasterLoginBonusItem()).WithContext(ctx).Select("master_login_bonus_item_id", "master_login_bonus_schedule_id", "master_item_id", "name", "count").Where("master_login_bonus_item_id = ?", m.MasterLoginBonusItemId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonusItem.SetMasterLoginBonusItem(t.MasterLoginBonusItemId, t.MasterLoginBonusScheduleId, t.MasterItemId, t.Name, t.Count), nil
}

func (s *masterLoginBonusItemDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusItem.MasterLoginBonusItems) (masterLoginBonusItem.MasterLoginBonusItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterLoginBonusItems()
	for _, m := range ms {
		t := &MasterLoginBonusItem{
			MasterLoginBonusItemId:     m.MasterLoginBonusItemId,
			MasterLoginBonusScheduleId: m.MasterLoginBonusScheduleId,
			MasterItemId:               m.MasterItemId,
			Name:                       m.Name,
			Count:                      m.Count,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonusItem()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_login_bonus_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"master_login_bonus_schedule_id", "master_item_id", "name", "count"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusItemDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonusItem.MasterLoginBonusItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterLoginBonusItem()).WithContext(ctx).Where("master_login_bonus_item_id = ?", m.MasterLoginBonusItemId).Delete(NewMasterLoginBonusItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterLoginBonusItemDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterLoginBonusItem.MasterLoginBonusItems) error {
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
		ks = append(ks, []interface{}{m.MasterLoginBonusItemId})
	}

	res := conn.Model(NewMasterLoginBonusItem()).WithContext(ctx).Where("(master_login_bonus_item_id) IN ?", ks).Delete(NewMasterLoginBonusItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
