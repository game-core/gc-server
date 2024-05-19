// Package masterItem アイテム
package masterItem

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/masterItem"
)

type masterItemDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterItemDao(conn *database.MysqlHandler) masterItem.MasterItemMysqlRepository {
	return &masterItemDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterItemDao) Find(ctx context.Context, masterItemId int64) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "Find", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterItem.SetMasterItem(t.MasterItemId, t.Name, t.MasterResourceEnum, t.MasterRarityEnum, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "Find", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindOrNil(ctx context.Context, masterItemId int64) (*masterItem.MasterItem, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindOrNil", fmt.Sprintf("%d_", masterItemId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterItem.MasterItem); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterItem()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_item_id = ?", masterItemId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterItem.SetMasterItem(t.MasterItemId, t.Name, t.MasterResourceEnum, t.MasterRarityEnum, t.Content)
	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindOrNil", fmt.Sprintf("%d_", masterItemId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterItemDao) FindList(ctx context.Context) (masterItem.MasterItems, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_item", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterItem.MasterItems); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterItems()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterItem.NewMasterItems()
	for _, t := range ts {
		ms = append(ms, masterItem.SetMasterItem(t.MasterItemId, t.Name, t.MasterResourceEnum, t.MasterRarityEnum, t.Content))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_item", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterItemDao) Create(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) (*masterItem.MasterItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterItem{
		MasterItemId:       m.MasterItemId,
		Name:               m.Name,
		MasterResourceEnum: m.MasterResourceEnum,
		MasterRarityEnum:   m.MasterRarityEnum,
		Content:            m.Content,
	}
	res := conn.Model(NewMasterItem()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterItem.SetMasterItem(t.MasterItemId, t.Name, t.MasterResourceEnum, t.MasterRarityEnum, t.Content), nil
}

func (s *masterItemDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterItem.MasterItems) (masterItem.MasterItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterItems()
	for _, m := range ms {
		t := &MasterItem{
			MasterItemId:       m.MasterItemId,
			Name:               m.Name,
			MasterResourceEnum: m.MasterResourceEnum,
			MasterRarityEnum:   m.MasterRarityEnum,
			Content:            m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterItem()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterItemDao) Update(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) (*masterItem.MasterItem, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterItem{
		MasterItemId:       m.MasterItemId,
		Name:               m.Name,
		MasterResourceEnum: m.MasterResourceEnum,
		MasterRarityEnum:   m.MasterRarityEnum,
		Content:            m.Content,
	}
	res := conn.Model(NewMasterItem()).WithContext(ctx).Where("master_item_id = ?", m.MasterItemId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterItem.SetMasterItem(t.MasterItemId, t.Name, t.MasterResourceEnum, t.MasterRarityEnum, t.Content), nil
}

func (s *masterItemDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterItem.MasterItems) (masterItem.MasterItems, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterItems()
	for _, m := range ms {
		t := &MasterItem{
			MasterItemId:       m.MasterItemId,
			Name:               m.Name,
			MasterResourceEnum: m.MasterResourceEnum,
			MasterRarityEnum:   m.MasterRarityEnum,
			Content:            m.Content,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterItem()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_item_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_resource_enum", "master_rarity_enum", "content"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterItemDao) Delete(ctx context.Context, tx *gorm.DB, m *masterItem.MasterItem) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterItem()).WithContext(ctx).Where("master_item_id = ?", m.MasterItemId).Delete(NewMasterItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterItemDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterItem.MasterItems) error {
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
		ks = append(ks, []interface{}{m.MasterItemId})
	}

	res := conn.Model(NewMasterItem()).WithContext(ctx).Where("(master_item_id) IN ?", ks).Delete(NewMasterItem())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
