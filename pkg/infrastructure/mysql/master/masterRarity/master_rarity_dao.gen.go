// Package masterRarity レアリティ
package masterRarity

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/internal/cashes"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/rarity/masterRarity"
)

type masterRarityDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
	Cache          *cache.Cache
}

func NewMasterRarityDao(conn *database.MysqlHandler) masterRarity.MasterRarityMysqlRepository {
	return &masterRarityDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
		Cache:          cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterRarityDao) Find(ctx context.Context, masterRarityId int64) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "Find", fmt.Sprintf("%d_", masterRarityId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_rarity_id = ?", masterRarityId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "Find", fmt.Sprintf("%d_", masterRarityId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindOrNil(ctx context.Context, masterRarityId int64) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindOrNil", fmt.Sprintf("%d_", masterRarityId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_rarity_id = ?", masterRarityId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindOrNil", fmt.Sprintf("%d_", masterRarityId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindByMasterRarityEnum(ctx context.Context, masterRarityEnum masterRarity.MasterRarityEnum) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_rarity_enum = ?", masterRarityEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindOrNilByMasterRarityEnum(ctx context.Context, masterRarityEnum masterRarity.MasterRarityEnum) (*masterRarity.MasterRarity, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindOrNilByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterRarity.MasterRarity); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterRarity()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_rarity_enum = ?", masterRarityEnum).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum)
	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindOrNilByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterRarityDao) FindList(ctx context.Context) (masterRarity.MasterRarities, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterRarity.MasterRarities); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRarities()
	res := s.ReadMysqlConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRarity.NewMasterRarities()
	for _, t := range ts {
		ms = append(ms, masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRarityDao) FindListByMasterRarityEnum(ctx context.Context, masterRarityEnum masterRarity.MasterRarityEnum) (masterRarity.MasterRarities, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_rarity", "FindListByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)))
	if found {
		if cachedEntity, ok := cachedResult.(masterRarity.MasterRarities); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterRarities()
	res := s.ReadMysqlConn.WithContext(ctx).Where("master_rarity_enum = ?", masterRarityEnum).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterRarity.NewMasterRarities()
	for _, t := range ts {
		ms = append(ms, masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_rarity", "FindListByMasterRarityEnum", fmt.Sprintf("%d_", masterRarityEnum)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterRarityDao) Create(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) (*masterRarity.MasterRarity, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRarity{
		MasterRarityId:   m.MasterRarityId,
		Name:             m.Name,
		MasterRarityEnum: m.MasterRarityEnum,
	}
	res := conn.Model(NewMasterRarity()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum), nil
}

func (s *masterRarityDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterRarity.MasterRarities) (masterRarity.MasterRarities, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRarities()
	for _, m := range ms {
		t := &MasterRarity{
			MasterRarityId:   m.MasterRarityId,
			Name:             m.Name,
			MasterRarityEnum: m.MasterRarityEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterRarity()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRarityDao) Update(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) (*masterRarity.MasterRarity, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	t := &MasterRarity{
		MasterRarityId:   m.MasterRarityId,
		Name:             m.Name,
		MasterRarityEnum: m.MasterRarityEnum,
	}
	res := conn.Model(NewMasterRarity()).WithContext(ctx).Where("master_rarity_id = ?", m.MasterRarityId).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterRarity.SetMasterRarity(t.MasterRarityId, t.Name, t.MasterRarityEnum), nil
}

func (s *masterRarityDao) UpdateList(ctx context.Context, tx *gorm.DB, ms masterRarity.MasterRarities) (masterRarity.MasterRarities, error) {
	if len(ms) <= 0 {
		return ms, nil
	}

	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	ts := NewMasterRarities()
	for _, m := range ms {
		t := &MasterRarity{
			MasterRarityId:   m.MasterRarityId,
			Name:             m.Name,
			MasterRarityEnum: m.MasterRarityEnum,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterRarity()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "master_rarity_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "master_rarity_enum"}),
	}).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterRarityDao) Delete(ctx context.Context, tx *gorm.DB, m *masterRarity.MasterRarity) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteMysqlConn
	}

	res := conn.Model(NewMasterRarity()).WithContext(ctx).Where("master_rarity_id = ?", m.MasterRarityId).Delete(NewMasterRarity())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func (s *masterRarityDao) DeleteList(ctx context.Context, tx *gorm.DB, ms masterRarity.MasterRarities) error {
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
		ks = append(ks, []interface{}{m.MasterRarityId})
	}

	res := conn.Model(NewMasterRarity()).WithContext(ctx).Where("(master_rarity_id) IN ?", ks).Delete(NewMasterRarity())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
